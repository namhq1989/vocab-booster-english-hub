package infrastructure

import (
	"database/sql"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/namhq1989/vocab-booster-english-hub/internal/database"
	"github.com/namhq1989/vocab-booster-english-hub/internal/database/gen/vocab-booster/public/model"
	"github.com/namhq1989/vocab-booster-english-hub/internal/database/gen/vocab-booster/public/table"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/domain"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/infrastructure/mapping"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type WordOfTheDayRepository struct {
	db *database.Database
}

func NewWordOfTheDayRepository(db *database.Database) WordOfTheDayRepository {
	r := WordOfTheDayRepository{
		db: db,
	}
	return r
}

func (r WordOfTheDayRepository) getDB() *sql.DB {
	return r.db.GetDB()
}

func (WordOfTheDayRepository) getTable() *table.WordOfTheDayTable {
	return table.WordOfTheDay
}

func (r WordOfTheDayRepository) FindWordOfTheDay(ctx *appcontext.AppContext) (*domain.WordOfTheDay, error) {
	stmt := postgres.SELECT(
		r.getTable().AllColumns,
	).
		FROM(r.getTable()).
		LIMIT(1).
		ORDER_BY(r.getTable().Date.DESC())

	var doc model.WordOfTheDay
	if err := stmt.QueryContext(ctx.Context(), r.getDB(), &doc); err != nil {
		if r.db.IsNoRowsError(err) {
			return nil, nil
		}
		return nil, err
	}

	var (
		mapper    = mapping.WordOfTheDayMapper{}
		result, _ = mapper.FromModelToDomain(doc)
	)
	return result, nil
}

func (r WordOfTheDayRepository) CreateWordOfTheDay(ctx *appcontext.AppContext, word domain.WordOfTheDay) error {
	mapper := mapping.WordOfTheDayMapper{}
	doc, err := mapper.FromDomainToModel(word)
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
