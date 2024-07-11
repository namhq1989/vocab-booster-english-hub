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

type CommunitySentenceLikeRepository struct {
	db *database.Database
}

func NewCommunitySentenceLikeRepository(db *database.Database) CommunitySentenceLikeRepository {
	r := CommunitySentenceLikeRepository{
		db: db,
	}
	return r
}

func (r CommunitySentenceLikeRepository) getDB() *sql.DB {
	return r.db.GetDB()
}

func (CommunitySentenceLikeRepository) getTable() *table.CommunitySentenceLikesTable {
	return table.CommunitySentenceLikes
}

func (r CommunitySentenceLikeRepository) FindCommunitySentenceLike(ctx *appcontext.AppContext, userID, sentenceID string) (*domain.CommunitySentenceLike, error) {
	if !database.IsValidID(userID) {
		return nil, apperrors.User.InvalidUserID
	}

	if !database.IsValidID(sentenceID) {
		return nil, apperrors.Vocabulary.InvalidSentence
	}

	stmt := postgres.SELECT(
		r.getTable().AllColumns,
	).
		FROM(r.getTable()).
		WHERE(
			r.getTable().UserID.EQ(postgres.String(userID)).
				AND(r.getTable().SentenceID.EQ(postgres.String(sentenceID))),
		)

	var doc model.CommunitySentenceLikes
	if err := stmt.QueryContext(ctx.Context(), r.getDB(), &doc); err != nil {
		if r.db.IsNoRowsError(err) {
			return nil, nil
		}
		return nil, err
	}

	var (
		mapper    = mapping.CommunitySentenceLikeMapper{}
		result, _ = mapper.FromModelToDomain(doc)
	)
	return result, nil
}

func (r CommunitySentenceLikeRepository) CreateCommunitySentenceLike(ctx *appcontext.AppContext, like domain.CommunitySentenceLike) error {
	mapper := mapping.CommunitySentenceLikeMapper{}
	doc, err := mapper.FromDomainToModel(like)
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

func (r CommunitySentenceLikeRepository) DeleteCommunitySentenceLike(ctx *appcontext.AppContext, like domain.CommunitySentenceLike) error {
	stmt := r.getTable().
		DELETE().
		WHERE(r.getTable().ID.EQ(postgres.String(like.ID)))

	_, err := stmt.ExecContext(ctx.Context(), r.getDB())
	return err
}
