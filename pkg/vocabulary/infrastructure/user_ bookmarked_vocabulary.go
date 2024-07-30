package infrastructure

import (
	"database/sql"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/namhq1989/vocab-booster-english-hub/internal/database"
	"github.com/namhq1989/vocab-booster-english-hub/internal/database/gen/vocab-booster/public/model"
	"github.com/namhq1989/vocab-booster-english-hub/internal/database/gen/vocab-booster/public/table"
	apperrors "github.com/namhq1989/vocab-booster-english-hub/internal/utils/error"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/domain"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/infrastructure/mapping"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type UserBookmarkedVocabularyRepository struct {
	db *database.Database
}

func NewUserBookmarkedVocabularyRepository(db *database.Database) UserBookmarkedVocabularyRepository {
	r := UserBookmarkedVocabularyRepository{
		db: db,
	}
	return r
}

func (r UserBookmarkedVocabularyRepository) getDB() *sql.DB {
	return r.db.GetDB()
}

func (UserBookmarkedVocabularyRepository) getTable() *table.UserBookmarkedVocabularyTable {
	return table.UserBookmarkedVocabulary
}

func (r UserBookmarkedVocabularyRepository) FindBookmarkedVocabulary(ctx *appcontext.AppContext, userID, vocabularyID string) (*domain.UserBookmarkedVocabulary, error) {
	if !database.IsValidID(userID) {
		return nil, apperrors.User.InvalidUserID
	}

	if !database.IsValidID(vocabularyID) {
		return nil, apperrors.Vocabulary.InvalidVocabularyID
	}

	stmt := postgres.SELECT(
		r.getTable().AllColumns,
	).
		FROM(r.getTable()).
		WHERE(
			r.getTable().UserID.EQ(postgres.String(userID)).
				AND(r.getTable().VocabularyID.EQ(postgres.String(vocabularyID))),
		)

	var doc model.UserBookmarkedVocabulary
	if err := stmt.QueryContext(ctx.Context(), r.getDB(), &doc); err != nil {
		if r.db.IsNoRowsError(err) {
			return nil, nil
		}
		return nil, err
	}

	var (
		mapper    = mapping.UserBookmarkedVocabularyMapper{}
		result, _ = mapper.FromModelToDomain(doc)
	)
	return result, nil
}

func (r UserBookmarkedVocabularyRepository) FindBookmarkedVocabulariesByUserID(ctx *appcontext.AppContext, userID string, filter domain.UserBookmarkedVocabularyFilter) ([]domain.ExtendedUserBookmarkedVocabulary, error) {
	if !database.IsValidID(userID) {
		return nil, apperrors.User.InvalidUserID
	}

	var (
		ubv = r.getTable().AS("ubv")
		v   = table.Vocabularies.AS("v")
	)

	stmt := postgres.SELECT(
		v.ID, v.Term, v.Ipa, v.PartsOfSpeech, v.Audio,
		ubv.BookmarkedAt,
	).
		FROM(
			ubv.LEFT_JOIN(v, v.ID.EQ(ubv.VocabularyID)),
		).
		WHERE(
			ubv.UserID.EQ(postgres.String(userID)).
				AND(ubv.BookmarkedAt.LT_EQ(postgres.TimestampzT(filter.Timestamp))),
		).
		LIMIT(filter.Limit).
		ORDER_BY(ubv.BookmarkedAt.DESC())

	var (
		docs   = make([]mapping.ExtendedUserBookmarkedVocabulary, 0)
		result = make([]domain.ExtendedUserBookmarkedVocabulary, 0)
	)
	if err := stmt.QueryContext(ctx.Context(), r.getDB(), &docs); err != nil {
		if r.db.IsNoRowsError(err) {
			return result, nil
		}
		return result, err
	}

	var (
		mapper = mapping.ExtendedUserBookmarkedVocabularyMapper{}
	)
	for _, doc := range docs {
		ue, _ := mapper.FromModelToDomain(doc)
		result = append(result, *ue)
	}
	return result, nil
}

func (r UserBookmarkedVocabularyRepository) CreateUserBookmarkedVocabulary(ctx *appcontext.AppContext, ubv domain.UserBookmarkedVocabulary) error {
	mapper := mapping.UserBookmarkedVocabularyMapper{}
	doc, err := mapper.FromDomainToModel(ubv)
	if err != nil {
		return err
	}

	stmt := r.getTable().INSERT(
		r.getTable().AllColumns,
	).
		MODEL(doc).
		ON_CONFLICT().DO_UPDATE(
		postgres.SET(r.getTable().BookmarkedAt.SET(postgres.TimestampzT(doc.BookmarkedAt))),
	)

	_, err = stmt.ExecContext(ctx.Context(), r.getDB())
	return err
}

func (r UserBookmarkedVocabularyRepository) DeleteUserBookmarkedVocabulary(ctx *appcontext.AppContext, ubv domain.UserBookmarkedVocabulary) error {
	stmt := r.getTable().
		DELETE().
		WHERE(r.getTable().ID.EQ(postgres.String(ubv.ID)))

	_, err := stmt.ExecContext(ctx.Context(), r.getDB())
	return err
}
