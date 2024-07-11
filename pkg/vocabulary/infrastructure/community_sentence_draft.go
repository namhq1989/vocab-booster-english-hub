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

func (r CommunitySentenceDraftRepository) DeleteCommunitySentenceDrafts(ctx *appcontext.AppContext, vocabularyID, userID string) error {
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
