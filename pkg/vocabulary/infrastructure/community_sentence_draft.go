package infrastructure

import (
	"database/sql"

	apperrors "github.com/namhq1989/vocab-booster-english-hub/internal/utils/error"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/namhq1989/vocab-booster-english-hub/internal/database"
	"github.com/namhq1989/vocab-booster-english-hub/internal/database/gen/vocab-booster/public/model"
	"github.com/namhq1989/vocab-booster-english-hub/internal/database/gen/vocab-booster/public/table"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/domain"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/infrastructure/mapping"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type CommunitySentenceDraftRepository struct {
	db *database.Database
}

func NewCommunitySentenceDraftRepository(db *database.Database) CommunitySentenceDraftRepository {
	r := CommunitySentenceDraftRepository{
		db: db,
	}
	return r
}

func (r CommunitySentenceDraftRepository) getDB() *sql.DB {
	return r.db.GetDB()
}

func (CommunitySentenceDraftRepository) getTable() *table.CommunitySentenceDraftsTable {
	return table.CommunitySentenceDrafts
}

func (r CommunitySentenceDraftRepository) FindCommunitySentenceDraftByID(ctx *appcontext.AppContext, sentenceID string) (*domain.CommunitySentenceDraft, error) {
	if !database.IsValidID(sentenceID) {
		return nil, apperrors.Vocabulary.InvalidSentence
	}

	stmt := postgres.SELECT(
		r.getTable().AllColumns,
	).
		FROM(r.getTable()).
		WHERE(r.getTable().ID.EQ(postgres.String(sentenceID)))

	var doc model.CommunitySentenceDrafts
	if err := stmt.QueryContext(ctx.Context(), r.getDB(), &doc); err != nil {
		if r.db.IsNoRowsError(err) {
			return nil, nil
		}
		return nil, err
	}

	var (
		mapper    = mapping.CommunitySentenceDraftMapper{}
		result, _ = mapper.FromModelToDomain(doc)
	)
	return result, nil
}

func (r CommunitySentenceDraftRepository) CreateCommunitySentenceDraft(ctx *appcontext.AppContext, sentence domain.CommunitySentenceDraft) error {
	mapper := mapping.CommunitySentenceDraftMapper{}
	doc, err := mapper.FromDomainToModel(sentence)
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

func (r CommunitySentenceDraftRepository) UpdateCommunitySentenceDraft(ctx *appcontext.AppContext, sentence domain.CommunitySentenceDraft) error {
	mapper := mapping.CommunitySentenceDraftMapper{}
	doc, err := mapper.FromDomainToModel(sentence)
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

func (r CommunitySentenceDraftRepository) DeleteCommunitySentenceDraft(ctx *appcontext.AppContext, vocabularyID, userID string) error {
	if !database.IsValidID(userID) {
		return apperrors.User.InvalidUserID
	}

	if !database.IsValidID(vocabularyID) {
		return apperrors.Vocabulary.InvalidVocabularyID
	}

	stmt := r.getTable().
		DELETE().
		WHERE(
			r.getTable().VocabularyID.EQ(postgres.String(vocabularyID)).
				AND(r.getTable().UserID.EQ(postgres.String(userID))).
				AND(r.getTable().IsCorrect.EQ(postgres.Bool(false))),
		)

	_, err := stmt.ExecContext(ctx.Context(), r.getDB())
	return err
}

func (r CommunitySentenceDraftRepository) FindUserCommunitySentenceDrafts(ctx *appcontext.AppContext, filter domain.CommunitySentenceDraftFilter) ([]domain.CommunitySentenceDraft, error) {
	whereCond := r.getTable().UserID.EQ(postgres.String(filter.UserID))
	if filter.VocabularyID != "" {
		whereCond = whereCond.AND(r.getTable().VocabularyID.EQ(postgres.String(filter.VocabularyID)))
	}
	whereCond = whereCond.AND(r.getTable().CreatedAt.LT(postgres.TimestampzT(filter.Timestamp)))

	stmt := postgres.SELECT(
		r.getTable().AllColumns,
	).
		FROM(r.getTable()).
		WHERE(whereCond).
		LIMIT(filter.Limit).
		ORDER_BY(r.getTable().CreatedAt.DESC())

	var (
		docs   = make([]model.CommunitySentenceDrafts, 0)
		result = make([]domain.CommunitySentenceDraft, 0)
	)
	if err := stmt.QueryContext(ctx.Context(), r.getDB(), &docs); err != nil {
		if r.db.IsNoRowsError(err) {
			return result, nil
		}
		return result, err
	}

	var (
		mapper = mapping.CommunitySentenceDraftMapper{}
	)
	for _, doc := range docs {
		ue, _ := mapper.FromModelToDomain(doc)
		result = append(result, *ue)
	}
	return result, nil
}
