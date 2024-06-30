package infrastructure

import (
	"database/sql"

	"github.com/go-jet/jet/v2/postgres"
	apperrors "github.com/namhq1989/vocab-booster-english-hub/core/error"
	"github.com/namhq1989/vocab-booster-english-hub/internal/database/gen/vocab-booster/public/model"

	"github.com/namhq1989/vocab-booster-english-hub/core/appcontext"
	"github.com/namhq1989/vocab-booster-english-hub/internal/database"
	"github.com/namhq1989/vocab-booster-english-hub/internal/database/gen/vocab-booster/public/table"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/domain"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/infrastructure/mapping"
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
