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

type UserVocabularyCollectionRepository struct {
	db *database.Database
}

func NewUserVocabularyCollectionRepository(db *database.Database) UserVocabularyCollectionRepository {
	r := UserVocabularyCollectionRepository{
		db: db,
	}
	return r
}

func (r UserVocabularyCollectionRepository) getDB() *sql.DB {
	return r.db.GetDB()
}

func (UserVocabularyCollectionRepository) getTable() *table.UserVocabularyCollectionsTable {
	return table.UserVocabularyCollections
}

func (r UserVocabularyCollectionRepository) CreateUserVocabularyCollection(ctx *appcontext.AppContext, collection domain.UserVocabularyCollection) error {
	mapper := mapping.UserVocabularyCollectionMapper{}
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

func (r UserVocabularyCollectionRepository) UpdateUserVocabularyCollection(ctx *appcontext.AppContext, collection domain.UserVocabularyCollection) error {
	mapper := mapping.UserVocabularyCollectionMapper{}
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

func (r UserVocabularyCollectionRepository) FindUserVocabularyCollectionsByUserID(ctx *appcontext.AppContext, userID string) ([]domain.UserVocabularyCollection, error) {
	if !database.IsValidID(userID) {
		return nil, apperrors.Common.InvalidID
	}

	stmt := postgres.SELECT(
		r.getTable().AllColumns,
	).
		FROM(r.getTable()).
		WHERE(
			r.getTable().UserID.EQ(postgres.String(userID)),
		).
		ORDER_BY(
			r.getTable().CreatedAt.DESC(),
		)

	var docs []model.UserVocabularyCollections
	if err := stmt.QueryContext(ctx.Context(), r.getDB(), &docs); err != nil {
		if r.db.IsNoRowsError(err) {
			return nil, nil
		}
		return nil, err
	}

	var (
		result = make([]domain.UserVocabularyCollection, len(docs))
		mapper = mapping.UserVocabularyCollectionMapper{}
	)
	for _, doc := range docs {
		collection, _ := mapper.FromModelToDomain(doc)
		result = append(result, *collection)
	}
	return result, nil
}

func (r UserVocabularyCollectionRepository) FindUserVocabularyCollectionByID(ctx *appcontext.AppContext, id string) (*domain.UserVocabularyCollection, error) {
	if !database.IsValidID(id) {
		return nil, apperrors.Common.InvalidID
	}

	stmt := postgres.SELECT(
		r.getTable().AllColumns,
	).
		FROM(r.getTable()).
		WHERE(r.getTable().ID.EQ(postgres.String(id)))

	var doc model.UserVocabularyCollections
	if err := stmt.QueryContext(ctx.Context(), r.getDB(), &doc); err != nil {
		if r.db.IsNoRowsError(err) {
			return nil, nil
		}
		return nil, err
	}

	var (
		mapper    = mapping.UserVocabularyCollectionMapper{}
		result, _ = mapper.FromModelToDomain(doc)
	)
	return result, nil
}

func (r UserVocabularyCollectionRepository) CountTotalUserVocabularyCollectionsByUserID(ctx *appcontext.AppContext, userID string) (int64, error) {
	if !database.IsValidID(userID) {
		return 0, apperrors.Common.InvalidID
	}

	stmt := postgres.SELECT(
		postgres.COUNT(r.getTable().ID).AS("count_result.total"),
	).
		FROM(r.getTable()).
		WHERE(
			r.getTable().UserID.EQ(postgres.String(userID)),
		)

	var result = database.CountResult{}
	err := stmt.QueryContext(ctx.Context(), r.getDB(), &result)
	return result.Total, err
}
