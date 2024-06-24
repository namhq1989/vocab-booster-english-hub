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

func (r UserExerciseStatusRepository) getTable() *table.UserExerciseStatusesTable {
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
		WHERE(r.getTable().ExerciseID.EQ(postgres.String(exerciseID))).
		WHERE(r.getTable().UserID.EQ(postgres.String(userID)))

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
