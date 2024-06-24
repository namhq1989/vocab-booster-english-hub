package infrastructure

import (
	"database/sql"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/namhq1989/vocab-booster-english-hub/core/appcontext"
	"github.com/namhq1989/vocab-booster-english-hub/internal/database"
	"github.com/namhq1989/vocab-booster-english-hub/internal/database/gen/vocab-booster/public/model"
	"github.com/namhq1989/vocab-booster-english-hub/internal/database/gen/vocab-booster/public/table"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/domain"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/infrastructure/mapping"
)

type VocabularyScrapingItemRepository struct {
	db *database.Database
}

func NewVocabularyScrapingItemRepository(db *database.Database) VocabularyScrapingItemRepository {
	r := VocabularyScrapingItemRepository{
		db: db,
	}
	return r
}

func (r VocabularyScrapingItemRepository) getDB() *sql.DB {
	return r.db.GetDB()
}

func (r VocabularyScrapingItemRepository) getTable() *table.VocabularyScrapingItemsTable {
	return table.VocabularyScrapingItems
}

func (r VocabularyScrapingItemRepository) FindVocabularyScrapingItemByTerm(ctx *appcontext.AppContext, term string) (*domain.VocabularyScrapingItem, error) {
	stmt := postgres.SELECT(
		r.getTable().AllColumns,
	).
		FROM(r.getTable()).
		WHERE(
			r.getTable().Term.EQ(postgres.String(term)),
		)

	var doc model.VocabularyScrapingItems
	if err := stmt.QueryContext(ctx.Context(), r.getDB(), &doc); err != nil {
		if r.db.IsNoRowsError(err) {
			return nil, nil
		}
		return nil, err
	}

	var (
		mapper    = mapping.VocabularyScrapingItemMapper{}
		result, _ = mapper.FromModelToDomain(doc)
	)
	return result, nil
}

func (r VocabularyScrapingItemRepository) CreateVocabularyScrapingItems(ctx *appcontext.AppContext, items []domain.VocabularyScrapingItem) error {
	var (
		docs   = make([]model.VocabularyScrapingItems, 0)
		mapper = mapping.VocabularyScrapingItemMapper{}
	)
	for _, item := range items {
		doc, err := mapper.FromDomainToModel(item)
		if err == nil {
			docs = append(docs, *doc)
		} else {
			ctx.Logger().Error("failed to mapping vocabulary scraping item", err, appcontext.Fields{"item": item})
		}
	}

	stmt := r.getTable().INSERT(
		r.getTable().AllColumns,
	).
		MODELS(docs).
		ON_CONFLICT().DO_NOTHING()

	_, err := stmt.ExecContext(ctx.Context(), r.getDB())
	return err
}

func (r VocabularyScrapingItemRepository) DeleteVocabularyScrapingItemByTerm(ctx *appcontext.AppContext, term string) error {
	stmt := r.getTable().
		DELETE().
		WHERE(r.getTable().Term.EQ(postgres.String(term)))

	_, err := stmt.ExecContext(ctx.Context(), r.getDB())
	return err
}

func (r VocabularyScrapingItemRepository) RandomPickVocabularyScrapingItem(ctx *appcontext.AppContext) (*domain.VocabularyScrapingItem, error) {
	stmt := postgres.SELECT(
		r.getTable().AllColumns,
	).
		FROM(r.getTable()).
		ORDER_BY(postgres.Raw("RANDOM()")).
		LIMIT(1)

	var docs []model.VocabularyScrapingItems
	if err := stmt.QueryContext(ctx.Context(), r.getDB(), &docs); err != nil {
		if r.db.IsNoRowsError(err) {
			return nil, nil
		}
		return nil, err
	}

	if len(docs) == 0 {
		return nil, nil
	}

	var (
		mapper    = mapping.VocabularyScrapingItemMapper{}
		result, _ = mapper.FromModelToDomain(docs[0])
	)
	return result, nil
}
