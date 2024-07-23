package infrastructure

import (
	"database/sql"
	"strings"
	"time"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/namhq1989/vocab-booster-english-hub/internal/database"
	"github.com/namhq1989/vocab-booster-english-hub/internal/database/gen/vocab-booster/public/model"
	"github.com/namhq1989/vocab-booster-english-hub/internal/database/gen/vocab-booster/public/table"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/exercise/domain"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/exercise/infrastructure/mapping"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type ExerciseRepository struct {
	db *database.Database
}

func NewExerciseRepository(db *database.Database) ExerciseRepository {
	r := ExerciseRepository{
		db: db,
	}
	return r
}

func (r ExerciseRepository) getDB() *sql.DB {
	return r.db.GetDB()
}

func (ExerciseRepository) getTable() *table.ExercisesTable {
	return table.Exercises
}

func (r ExerciseRepository) FindExerciseByID(ctx *appcontext.AppContext, exerciseID string) (*domain.Exercise, error) {
	stmt := postgres.SELECT(
		r.getTable().AllColumns,
	).
		FROM(r.getTable()).
		WHERE(r.getTable().ID.EQ(postgres.String(exerciseID)))

	var doc model.Exercises
	if err := stmt.QueryContext(ctx.Context(), r.getDB(), &doc); err != nil {
		if r.db.IsNoRowsError(err) {
			return nil, nil
		}
		return nil, err
	}

	var (
		mapper    = mapping.ExerciseMapper{}
		result, _ = mapper.FromModelToDomain(doc)
	)
	return result, nil
}

func (r ExerciseRepository) FindExerciseByVocabularyExampleID(ctx *appcontext.AppContext, exampleID string) (*domain.Exercise, error) {
	stmt := postgres.SELECT(
		r.getTable().AllColumns,
	).
		FROM(r.getTable()).
		WHERE(r.getTable().VocabularyExampleID.EQ(postgres.String(exampleID)))

	var doc model.Exercises
	if err := stmt.QueryContext(ctx.Context(), r.getDB(), &doc); err != nil {
		if r.db.IsNoRowsError(err) {
			return nil, nil
		}
		return nil, err
	}

	var (
		mapper    = mapping.ExerciseMapper{}
		result, _ = mapper.FromModelToDomain(doc)
	)
	return result, nil
}

func (r ExerciseRepository) CreateExercise(ctx *appcontext.AppContext, exercise domain.Exercise) error {
	mapper := mapping.ExerciseMapper{}
	doc, err := mapper.FromDomainToModel(exercise)
	if err != nil {
		return err
	}

	stmt := r.getTable().INSERT(
		r.getTable().AllColumns,
	).
		MODEL(doc)

	_, err = stmt.ExecContext(ctx.Context(), r.getDB())
	return err
}

func (r ExerciseRepository) UpdateExercise(ctx *appcontext.AppContext, exercise domain.Exercise) error {
	mapper := mapping.ExerciseMapper{}
	doc, err := mapper.FromDomainToModel(exercise)
	if err != nil {
		return err
	}

	stmt := r.getTable().UPDATE(
		r.getTable().AllColumns,
	).
		MODEL(doc).
		WHERE(
			r.getTable().ID.EQ(postgres.String(doc.ID)),
		)

	_, err = stmt.ExecContext(ctx.Context(), r.getDB())
	return err
}

func (r ExerciseRepository) PickRandomExercisesForUser(ctx *appcontext.AppContext, filter domain.UserExerciseFilter) ([]domain.UserExercise, error) {
	var (
		e   = r.getTable().AS("e")
		ues = table.UserExerciseStatuses.AS("ues")
	)

	whereCond := ues.ExerciseID.IS_NULL().OR(ues.IsMastered.EQ(postgres.Bool(false)))
	if filter.CollectionCriteria != "" {
		parts := strings.Split(filter.CollectionCriteria, "=")
		if len(parts) == 2 {
			if parts[0] == "level" && parts[1] != "" {
				whereCond = whereCond.AND(e.Level.EQ(postgres.String(parts[1])))
			}
		}
	}

	stmt := postgres.SELECT(
		e.ID, e.Level, e.Audio, e.Vocabulary, e.Content, e.CorrectAnswer, e.Options,
		ues.CorrectStreak, ues.IsFavorite, ues.IsMastered, ues.NextReviewAt,
	).
		FROM(
			e.LEFT_JOIN(ues, e.ID.EQ(ues.ExerciseID).AND(ues.UserID.EQ(postgres.String(filter.UserID)))),
		).
		WHERE(whereCond).
		ORDER_BY(
			postgres.Raw("RANDOM()"),
		).
		LIMIT(filter.NumOfExercises)

	var (
		docs   = make([]mapping.UserExercise, 0)
		result = make([]domain.UserExercise, 0)
	)
	if err := stmt.QueryContext(ctx.Context(), r.getDB(), &docs); err != nil {
		if r.db.IsNoRowsError(err) {
			return result, nil
		}
		return result, err
	}

	var (
		mapper = mapping.UserExerciseMapper{}
	)
	for _, doc := range docs {
		ue, _ := mapper.FromModelToDomain(doc, filter.Lang)
		result = append(result, *ue)
	}
	return result, nil
}

func (r ExerciseRepository) CountExercisesByCriteria(ctx *appcontext.AppContext, criteria string, ts time.Time) (int64, error) {
	whereCond := r.getTable().CreatedAt.GT(postgres.TimestampzT(ts))

	if criteria != "" {
		parts := strings.Split(criteria, "=")
		if len(parts) == 2 {
			if parts[0] == "level" {
				whereCond = whereCond.AND(r.getTable().Level.EQ(postgres.String(parts[1])))
			}
		}
	}

	stmt := postgres.SELECT(
		postgres.COUNT(r.getTable().ID).AS("count_result.total"),
	).
		FROM(r.getTable()).
		WHERE(whereCond)

	var result = database.CountResult{}
	err := stmt.QueryContext(ctx.Context(), r.getDB(), &result)
	return result.Total, err
}
