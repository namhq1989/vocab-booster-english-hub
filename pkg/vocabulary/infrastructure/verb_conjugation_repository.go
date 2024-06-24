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

type VerbConjugationRepository struct {
	db *database.Database
}

func NewVerbConjugationRepository(db *database.Database) VerbConjugationRepository {
	r := VerbConjugationRepository{
		db: db,
	}
	return r
}

func (r VerbConjugationRepository) getDB() *sql.DB {
	return r.db.GetDB()
}

func (r VerbConjugationRepository) getTable() *table.VerbConjugationsTable {
	return table.VerbConjugations
}

func (r VerbConjugationRepository) FindVerbConjugationByValue(ctx *appcontext.AppContext, value string) (*domain.VerbConjugation, error) {
	stmt := postgres.SELECT(
		r.getTable().AllColumns,
	).
		FROM(r.getTable()).
		WHERE(r.getTable().Value.EQ(postgres.String(value)))

	var doc model.VerbConjugations
	if err := stmt.QueryContext(ctx.Context(), r.getDB(), &doc); err != nil {
		if r.db.IsNoRowsError(err) {
			return nil, nil
		}
		return nil, err
	}

	var (
		mapper    = mapping.VerbConjugationMapper{}
		result, _ = mapper.FromModelToDomain(doc)
	)
	return result, nil
}

func (r VerbConjugationRepository) CreateVerbConjugations(ctx *appcontext.AppContext, verbs []domain.VerbConjugation) error {
	var (
		docs   = make([]model.VerbConjugations, 0)
		mapper = mapping.VerbConjugationMapper{}
	)
	for _, verb := range verbs {
		doc, err := mapper.FromDomainToModel(verb)
		if err == nil {
			docs = append(docs, *doc)
		} else {
			ctx.Logger().Error("failed to mapping verb conjugation", err, appcontext.Fields{"verb": verb})
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

func (r VerbConjugationRepository) FindVerbConjugationByVocabularyID(ctx *appcontext.AppContext, vocabularyID string) ([]domain.VerbConjugation, error) {
	stmt := postgres.SELECT(
		r.getTable().AllColumns,
	).
		FROM(r.getTable()).
		WHERE(
			r.getTable().VocabularyID.EQ(postgres.String(vocabularyID)),
		)

	var docs []model.VerbConjugations
	if err := stmt.QueryContext(ctx.Context(), r.getDB(), &docs); err != nil {
		if r.db.IsNoRowsError(err) {
			return nil, nil
		}
		return nil, err
	}

	var (
		result = make([]domain.VerbConjugation, len(docs))
		mapper = mapping.VerbConjugationMapper{}
	)
	for _, doc := range docs {
		verb, _ := mapper.FromModelToDomain(doc)
		result = append(result, *verb)
	}
	return result, nil
}
