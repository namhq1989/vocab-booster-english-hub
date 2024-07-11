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

type VocabularyRepository struct {
	db *database.Database
}

func NewVocabularyRepository(db *database.Database) VocabularyRepository {
	r := VocabularyRepository{
		db: db,
	}
	return r
}

func (r VocabularyRepository) getDB() *sql.DB {
	return r.db.GetDB()
}

func (VocabularyRepository) getTable() *table.VocabulariesTable {
	return table.Vocabularies
}

func (r VocabularyRepository) FindVocabularyByID(ctx *appcontext.AppContext, vocabularyID string) (*domain.Vocabulary, error) {
	stmt := postgres.SELECT(
		r.getTable().AllColumns,
	).
		FROM(r.getTable()).
		WHERE(
			r.getTable().ID.EQ(postgres.String(vocabularyID)),
		)

	var doc model.Vocabularies
	if err := stmt.QueryContext(ctx.Context(), r.getDB(), &doc); err != nil {
		if r.db.IsNoRowsError(err) {
			return nil, nil
		}
		return nil, err
	}

	var (
		mapper    = mapping.VocabularyMapper{}
		result, _ = mapper.FromModelToDomain(doc)
	)
	return result, nil
}

func (r VocabularyRepository) FindVocabularyByTerm(ctx *appcontext.AppContext, term string) (*domain.Vocabulary, error) {
	stmt := postgres.SELECT(
		r.getTable().AllColumns,
	).
		FROM(r.getTable()).
		WHERE(
			r.getTable().Term.EQ(postgres.String(term)),
		)

	var doc model.Vocabularies
	if err := stmt.QueryContext(ctx.Context(), r.getDB(), &doc); err != nil {
		if r.db.IsNoRowsError(err) {
			return nil, nil
		}
		return nil, err
	}

	var (
		mapper    = mapping.VocabularyMapper{}
		result, _ = mapper.FromModelToDomain(doc)
	)
	return result, nil
}

func (r VocabularyRepository) CreateVocabulary(ctx *appcontext.AppContext, vocabulary domain.Vocabulary) error {
	mapper := mapping.VocabularyMapper{}
	doc, err := mapper.FromDomainToModel(vocabulary)
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

func (r VocabularyRepository) UpdateVocabulary(ctx *appcontext.AppContext, vocabulary domain.Vocabulary) error {
	mapper := mapping.VocabularyMapper{}
	doc, err := mapper.FromDomainToModel(vocabulary)
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

func (r VocabularyRepository) RandomPickVocabularyForExercise(ctx *appcontext.AppContext, numOfVocabulary int64) ([]domain.Vocabulary, error) {
	stmt := postgres.SELECT(
		r.getTable().AllColumns,
	).
		FROM(r.getTable()).
		ORDER_BY(postgres.Raw("RANDOM()")).
		LIMIT(numOfVocabulary)

	var docs []model.Vocabularies
	if err := stmt.QueryContext(ctx.Context(), r.getDB(), &docs); err != nil {
		if r.db.IsNoRowsError(err) {
			return nil, nil
		}
		return nil, err
	}

	var (
		result = make([]domain.Vocabulary, 0)
		mapper = mapping.VocabularyMapper{}
	)
	for _, doc := range docs {
		vocab, _ := mapper.FromModelToDomain(doc)
		result = append(result, *vocab)
	}
	return result, nil
}
