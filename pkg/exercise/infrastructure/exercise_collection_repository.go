package infrastructure

import (
	"database/sql"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/namhq1989/vocab-booster-english-hub/internal/database"
	"github.com/namhq1989/vocab-booster-english-hub/internal/database/gen/vocab-booster/public/model"
	"github.com/namhq1989/vocab-booster-english-hub/internal/database/gen/vocab-booster/public/table"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/exercise/domain"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/exercise/infrastructure/mapping"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type ExerciseCollectionRepository struct {
	db *database.Database
}

func NewExerciseCollectionRepository(db *database.Database) ExerciseCollectionRepository {
	r := ExerciseCollectionRepository{
		db: db,
	}
	return r
}

func (r ExerciseCollectionRepository) getDB() *sql.DB {
	return r.db.GetDB()
}

func (ExerciseCollectionRepository) getTable() *table.ExerciseCollectionsTable {
	return table.ExerciseCollections
}

func (r ExerciseCollectionRepository) CreateExerciseCollection(ctx *appcontext.AppContext, collection domain.ExerciseCollection) error {
	mapper := mapping.ExerciseCollectionMapper{}
	doc, err := mapper.FromDomainToModel(collection)
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

func (r ExerciseCollectionRepository) UpdateExerciseCollection(ctx *appcontext.AppContext, collection domain.ExerciseCollection) error {
	mapper := mapping.ExerciseCollectionMapper{}
	doc, err := mapper.FromDomainToModel(collection)
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

func (r ExerciseCollectionRepository) CountExerciseCollections(ctx *appcontext.AppContext) (int64, error) {
	stmt := postgres.SELECT(
		postgres.COUNT(r.getTable().ID).AS("count_result.total"),
	).
		FROM(r.getTable())

	var result = database.CountResult{}
	err := stmt.QueryContext(ctx.Context(), r.getDB(), &result)
	return result.Total, err
}

func (r ExerciseCollectionRepository) FindExerciseCollections(ctx *appcontext.AppContext) ([]domain.ExerciseCollection, error) {
	stmt := postgres.SELECT(r.getTable().AllColumns).
		FROM(r.getTable()).
		ORDER_BY(r.getTable().Order.ASC())

	var (
		docs   = make([]model.ExerciseCollections, 0)
		result = make([]domain.ExerciseCollection, 0)
	)
	if err := stmt.QueryContext(ctx.Context(), r.getDB(), &docs); err != nil {
		ctx.Logger().Print("err", err)
		if r.db.IsNoRowsError(err) {
			return result, nil
		}
		return result, err
	}

	var (
		mapper = mapping.ExerciseCollectionMapper{}
	)
	for _, doc := range docs {
		collection, _ := mapper.FromModelToDomain(doc)
		result = append(result, *collection)
	}
	return result, nil
}
