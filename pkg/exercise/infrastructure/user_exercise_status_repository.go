package infrastructure

import (
	"database/sql"
	"time"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/namhq1989/vocab-booster-english-hub/core/appcontext"
	"github.com/namhq1989/vocab-booster-english-hub/internal/database"
	"github.com/namhq1989/vocab-booster-english-hub/internal/database/gen/vocab-booster/public/model"
	"github.com/namhq1989/vocab-booster-english-hub/internal/database/gen/vocab-booster/public/table"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/exercise/domain"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/exercise/infrastructure/mapping"
)

type UserExerciseStatusRepository struct {
	db *database.Database
}

func NewUserExerciseStatusRepository(db *database.Database) UserExerciseStatusRepository {
	r := UserExerciseStatusRepository{
		db: db,
	}
	return r
}

func (r UserExerciseStatusRepository) getDB() *sql.DB {
	return r.db.GetDB()
}

func (UserExerciseStatusRepository) getTable() *table.UserExerciseStatusesTable {
	return table.UserExerciseStatuses
}

func (r UserExerciseStatusRepository) CreateUserExerciseStatus(ctx *appcontext.AppContext, status domain.UserExerciseStatus) error {
	mapper := mapping.UserExerciseStatusMapper{}
	doc, err := mapper.FromDomainToModel(status)
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

func (r UserExerciseStatusRepository) UpdateUserExerciseStatus(ctx *appcontext.AppContext, status domain.UserExerciseStatus) error {
	mapper := mapping.UserExerciseStatusMapper{}
	doc, err := mapper.FromDomainToModel(status)
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

func (r UserExerciseStatusRepository) FindUserExerciseStatus(ctx *appcontext.AppContext, exerciseID, userID string) (*domain.UserExerciseStatus, error) {
	stmt := postgres.SELECT(
		r.getTable().AllColumns,
	).
		FROM(r.getTable()).
		WHERE(
			r.getTable().ExerciseID.EQ(postgres.String(exerciseID)).
				AND(r.getTable().UserID.EQ(postgres.String(userID))),
		)

	var doc model.UserExerciseStatuses
	if err := stmt.QueryContext(ctx.Context(), r.getDB(), &doc); err != nil {
		if r.db.IsNoRowsError(err) {
			return nil, nil
		}
		return nil, err
	}

	var (
		mapper    = mapping.UserExerciseStatusMapper{}
		result, _ = mapper.FromModelToDomain(doc)
	)
	return result, nil
}

func (r UserExerciseStatusRepository) CountUserReadyToReviewExercises(ctx *appcontext.AppContext, userID string) (int64, error) {
	var now = time.Now()

	stmt := postgres.SELECT(
		postgres.COUNT(r.getTable().ID).AS("count_result.total"),
	).
		FROM(r.getTable()).
		WHERE(
			r.getTable().UserID.EQ(postgres.String(userID)).
				AND(r.getTable().NextReviewAt.LT(postgres.TimestampzT(now))),
		)

	var result = database.CountResult{}
	err := stmt.QueryContext(ctx.Context(), r.getDB(), &result)
	return result.Total, err
}

func (r UserExerciseStatusRepository) GetUserReadyToReviewExercises(ctx *appcontext.AppContext, filter domain.UserExerciseFilter) ([]domain.UserExercise, error) {
	var (
		ues = r.getTable().AS("ues")
		e   = table.Exercises.AS("e")
		now = time.Now()
	)

	whereCond := postgres.AND(ues.UserID.EQ(postgres.String(filter.UserID)))
	whereCond = whereCond.AND(ues.NextReviewAt.LT(postgres.TimestampzT(now)))
	if filter.Level.String() != "" {
		whereCond = whereCond.AND(e.Level.EQ(postgres.String(filter.Level.String())))
	}

	stmt := postgres.SELECT(
		e.ID, e.Level, e.Audio, e.Vocabulary, e.Content, e.Translated, e.CorrectAnswer, e.Options,
		ues.CorrectStreak, ues.IsFavorite, ues.IsMastered, ues.NextReviewAt,
	).
		FROM(
			ues.LEFT_JOIN(e, ues.ExerciseID.EQ(e.ID)),
		).
		WHERE(whereCond).
		ORDER_BY(
			ues.NextReviewAt,
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
