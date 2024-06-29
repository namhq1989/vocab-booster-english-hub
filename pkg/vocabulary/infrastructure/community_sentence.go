package infrastructure

import (
	"database/sql"

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
