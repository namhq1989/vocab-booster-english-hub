package infrastructure

import (
	"database/sql"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/namhq1989/vocab-booster-english-hub/internal/database"
	"github.com/namhq1989/vocab-booster-english-hub/internal/database/gen/vocab-booster/public/table"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/exercise/domain"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/exercise/infrastructure/mapping"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type UserExerciseCollectionStatusRepository struct {
	db *database.Database
}

func NewUserExerciseCollectionStatusRepository(db *database.Database) UserExerciseCollectionStatusRepository {
	r := UserExerciseCollectionStatusRepository{
		db: db,
	}
	return r
}

func (r UserExerciseCollectionStatusRepository) getDB() *sql.DB {
	return r.db.GetDB()
}

func (UserExerciseCollectionStatusRepository) getTable() *table.UserExerciseCollectionStatusTable {
	return table.UserExerciseCollectionStatus
}

func (r UserExerciseCollectionStatusRepository) IncreaseUserExerciseCollectionStatusStats(ctx *appcontext.AppContext, uecs domain.UserExerciseCollectionStatus, numOfExercises int64) error {
	mapper := mapping.UserExerciseCollectionStatusMapper{}
	doc, err := mapper.FromDomainToModel(uecs)
	if err != nil {
		return err
	}

	stmt := r.getTable().INSERT(r.getTable().AllColumns).
		MODEL(doc).
		ON_CONFLICT(r.getTable().UserID, r.getTable().CollectionID).DO_UPDATE(
		postgres.SET(
			r.getTable().InteractedExercises.SET(r.getTable().InteractedExercises.ADD(postgres.Int64(numOfExercises))),
		),
	)

	_, err = stmt.ExecContext(ctx.Context(), r.getDB())
	return err
}
