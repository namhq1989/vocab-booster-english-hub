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

type CommunitySentenceRepository struct {
	db *database.Database
}

func NewCommunitySentenceRepository(db *database.Database) CommunitySentenceRepository {
	r := CommunitySentenceRepository{
		db: db,
	}
	return r
}

func (r CommunitySentenceRepository) getDB() *sql.DB {
	return r.db.GetDB()
}

func (CommunitySentenceRepository) getTable() *table.CommunitySentencesTable {
	return table.CommunitySentences
}

func (r CommunitySentenceRepository) CreateCommunitySentence(ctx *appcontext.AppContext, sentence domain.CommunitySentence) error {
	mapper := mapping.CommunitySentenceMapper{}
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

func (r CommunitySentenceRepository) UpdateCommunitySentence(ctx *appcontext.AppContext, sentence domain.CommunitySentence) error {
	mapper := mapping.CommunitySentenceMapper{}
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

func (r CommunitySentenceRepository) FindCommunitySentenceByID(ctx *appcontext.AppContext, sentenceID string) (*domain.CommunitySentence, error) {
	if !database.IsValidID(sentenceID) {
		return nil, apperrors.Vocabulary.InvalidSentence
	}

	stmt := postgres.SELECT(
		r.getTable().AllColumns,
	).
		FROM(r.getTable()).
		WHERE(r.getTable().ID.EQ(postgres.String(sentenceID)))

	var doc model.CommunitySentences
	if err := stmt.QueryContext(ctx.Context(), r.getDB(), &doc); err != nil {
		if r.db.IsNoRowsError(err) {
			return nil, nil
		}
		return nil, err
	}

	var (
		mapper    = mapping.CommunitySentenceMapper{}
		result, _ = mapper.FromModelToDomain(doc)
	)
	return result, nil
}

func (r CommunitySentenceRepository) FindCommunitySentences(ctx *appcontext.AppContext, filter domain.VocabularyCommunitySentenceFilter) ([]domain.ExtendedCommunitySentence, error) {
	var (
		cs  = r.getTable().AS("cs")
		csl = table.CommunitySentenceLikes.AS("csl")
	)

	isLikedExpr := postgres.CASE().
		WHEN(csl.UserID.IS_NOT_NULL()).
		THEN(postgres.Bool(true)).
		ELSE(postgres.Bool(false)).
		AS("csl.is_liked")

	stmt := postgres.SELECT(
		cs.ID, cs.Content, cs.StatsLike, cs.Level, cs.CreatedAt,
		isLikedExpr,
	).
		FROM(
			cs.LEFT_JOIN(csl, csl.SentenceID.EQ(cs.ID).
				AND(csl.UserID.EQ(postgres.String(filter.UserID)))),
		).
		WHERE(
			cs.VocabularyID.EQ(postgres.String(filter.VocabularyID)).
				AND(cs.CreatedAt.LT(postgres.TimestampzT(filter.Timestamp))),
		).
		LIMIT(filter.Limit).
		ORDER_BY(cs.CreatedAt.DESC())

	var (
		docs   = make([]mapping.ExtendedCommunitySentence, 0)
		result = make([]domain.ExtendedCommunitySentence, 0)
	)
	if err := stmt.QueryContext(ctx.Context(), r.getDB(), &docs); err != nil {
		if r.db.IsNoRowsError(err) {
			return result, nil
		}
		return result, err
	}

	var (
		mapper = mapping.ExtendedCommunitySentenceMapper{}
	)
	for _, doc := range docs {
		ue, _ := mapper.FromModelToDomain(doc)
		result = append(result, *ue)
	}
	return result, nil
}
