package infrastructure

import (
	"database/sql"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/namhq1989/vocab-booster-english-hub/internal/database"
	"github.com/namhq1989/vocab-booster-english-hub/internal/database/gen/vocab-booster/public/model"
	"github.com/namhq1989/vocab-booster-english-hub/internal/database/gen/vocab-booster/public/table"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/domain"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/infrastructure/mapping"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type VocabularyExampleRepository struct {
	db *database.Database
}

func NewVocabularyExampleRepository(db *database.Database) VocabularyExampleRepository {
	r := VocabularyExampleRepository{
		db: db,
	}
	return r
}

func (r VocabularyExampleRepository) getDB() *sql.DB {
	return r.db.GetDB()
}

func (VocabularyExampleRepository) getTable() *table.VocabularyExamplesTable {
	return table.VocabularyExamples
}

func (r VocabularyExampleRepository) FindVocabularyExamplesByVocabularyID(ctx *appcontext.AppContext, vocabularyID string) ([]domain.VocabularyExample, error) {
	stmt := postgres.SELECT(
		r.getTable().AllColumns,
	).
		FROM(r.getTable()).
		WHERE(
			r.getTable().VocabularyID.EQ(postgres.String(vocabularyID)),
		).
		ORDER_BY(r.getTable().CreatedAt.DESC())

	var (
		docs   []model.VocabularyExamples
		result = make([]domain.VocabularyExample, 0)
	)
	if err := stmt.QueryContext(ctx.Context(), r.getDB(), &docs); err != nil {
		if r.db.IsNoRowsError(err) {
			return result, nil
		}
		return result, err
	}

	mapper := mapping.VocabularyExampleMapper{}
	for _, doc := range docs {
		verb, _ := mapper.FromModelToDomain(doc)
		result = append(result, *verb)
	}
	return result, nil
}

func (r VocabularyExampleRepository) CreateVocabularyExamples(ctx *appcontext.AppContext, examples []domain.VocabularyExample) error {
	var (
		docs   = make([]model.VocabularyExamples, 0)
		mapper = mapping.VocabularyExampleMapper{}
	)
	for _, example := range examples {
		doc, err := mapper.FromDomainToModel(example)
		if err == nil {
			docs = append(docs, *doc)
		} else {
			ctx.Logger().Error("failed to mapping vocabulary example", err, appcontext.Fields{"example": example})
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

func (r VocabularyExampleRepository) UpdateVocabularyExample(ctx *appcontext.AppContext, example domain.VocabularyExample) error {
	mapper := mapping.VocabularyExampleMapper{}
	doc, err := mapper.FromDomainToModel(example)
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
