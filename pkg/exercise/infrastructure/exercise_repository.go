package infrastructure

import (
	"database/sql"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/namhq1989/vocab-booster-english-hub/core/appcontext"
	"github.com/namhq1989/vocab-booster-english-hub/internal/database"
	"github.com/namhq1989/vocab-booster-english-hub/internal/database/gen/vocab-booster/public/model"
	"github.com/namhq1989/vocab-booster-english-hub/internal/database/gen/vocab-booster/public/table"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/exercise/domain"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/exercise/infrastructure/mapping"
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

	whereCond := postgres.AND(postgres.BoolExp(postgres.COALESCE(ues.IsMastered, postgres.Bool(false).EQ(postgres.Bool(false)))))
	if filter.Level.String() != "" {
		whereCond = whereCond.AND(e.Level.EQ(postgres.String(filter.Level.String())))
	}

	stmt := postgres.SELECT(
		e.ID, e.Level, e.Audio, e.Vocabulary, e.Content, e.Translated, e.CorrectAnswer, e.Options,
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
		ctx.Logger().Print("err", err)
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
