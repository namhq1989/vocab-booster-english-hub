package infrastructure

import (
	"database/sql"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/namhq1989/vocab-booster-english-hub/core/appcontext"
	apperrors "github.com/namhq1989/vocab-booster-english-hub/core/error"
	"github.com/namhq1989/vocab-booster-english-hub/internal/database"
	"github.com/namhq1989/vocab-booster-english-hub/internal/database/gen/vocab-booster/public/model"
	"github.com/namhq1989/vocab-booster-english-hub/internal/database/gen/vocab-booster/public/table"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/domain"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/infrastructure/mapping"
)

type CollectionAndVocabularyRepository struct {
	db *database.Database
}

func NewCollectionAndVocabularyRepository(db *database.Database) CollectionAndVocabularyRepository {
	r := CollectionAndVocabularyRepository{
		db: db,
	}
	return r
}

func (r CollectionAndVocabularyRepository) getDB() *sql.DB {
	return r.db.GetDB()
}

func (CollectionAndVocabularyRepository) getTable() *table.CollectionAndVocabulariesTable {
	return table.CollectionAndVocabularies
}

func (r CollectionAndVocabularyRepository) CreateCollectionAndVocabulary(ctx *appcontext.AppContext, cav domain.CollectionAndVocabulary) error {
	mapper := mapping.CollectionAndVocabularyMapper{}
	doc, err := mapper.FromDomainToModel(cav)
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

func (r CollectionAndVocabularyRepository) DeleteCollectionAndVocabulary(ctx *appcontext.AppContext, cav domain.CollectionAndVocabulary) error {
	stmt := r.getTable().DELETE().
		WHERE(
			r.getTable().ID.EQ(postgres.String(cav.ID)),
		)

	_, err := stmt.ExecContext(ctx.Context(), r.getDB())
	return err
}

func (r CollectionAndVocabularyRepository) FindCollectionAndVocabulary(ctx *appcontext.AppContext, collectionID, vocabularyID string) (*domain.CollectionAndVocabulary, error) {
	if !database.IsValidID(collectionID) || !database.IsValidID(vocabularyID) {
		return nil, apperrors.Common.InvalidID
	}

	stmt := postgres.SELECT(
		r.getTable().AllColumns,
	).
		FROM(r.getTable()).
		WHERE(
			r.getTable().CollectionID.EQ(postgres.String(collectionID)).
				AND(r.getTable().VocabularyID.EQ(postgres.String(vocabularyID))),
		)

	var doc model.CollectionAndVocabularies
	if err := stmt.QueryContext(ctx.Context(), r.getDB(), &doc); err != nil {
		if r.db.IsNoRowsError(err) {
			return nil, nil
		}
		return nil, err
	}

	var (
		mapper    = mapping.CollectionAndVocabularyMapper{}
		result, _ = mapper.FromModelToDomain(doc)
	)
	return result, nil
}
