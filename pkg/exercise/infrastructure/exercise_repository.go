package infrastructure

import (
	"database/sql"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/namhq1989/vocab-booster-english-hub/core/appcontext"
	"github.com/namhq1989/vocab-booster-english-hub/internal/database"
	"github.com/namhq1989/vocab-booster-english-hub/internal/database/gen/vocab-booster/public/model"
	"github.com/namhq1989/vocab-booster-english-hub/internal/database/gen/vocab-booster/public/table"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/exercise/domain"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/exercise/infrastructure/mapping"
)

type ExerciseRepository struct {
	db *database.Database
}

func NewExerciseRepository(db *database.Database) ExerciseRepository {
	r := ExerciseRepository{
		db: db,
	}
	return r
}

func (r ExerciseRepository) getDB() *sql.DB {
	return r.db.GetDB()
}

func (ExerciseRepository) getTable() *table.ExercisesTable {
	return table.Exercises
}

func (r ExerciseRepository) FindExerciseByID(ctx *appcontext.AppContext, exerciseID string) (*domain.Exercise, error) {
	stmt := postgres.SELECT(
		r.getTable().AllColumns,
	).
		FROM(r.getTable()).
		WHERE(r.getTable().ID.EQ(postgres.String(exerciseID)))

	var doc model.Exercises
	if err := stmt.QueryContext(ctx.Context(), r.getDB(), &doc); err != nil {
		if r.db.IsNoRowsError(err) {
			return nil, nil
		}
		return nil, err
	}

	var (
		mapper    = mapping.ExerciseMapper{}
		result, _ = mapper.FromModelToDomain(doc)
	)
	return result, nil
}

func (r ExerciseRepository) FindExerciseByVocabularyExampleID(ctx *appcontext.AppContext, exampleID string) (*domain.Exercise, error) {
	stmt := postgres.SELECT(
		r.getTable().AllColumns,
	).
		FROM(r.getTable()).
		WHERE(r.getTable().VocabularyExampleID.EQ(postgres.String(exampleID)))

	var doc model.Exercises
	if err := stmt.QueryContext(ctx.Context(), r.getDB(), &doc); err != nil {
		if r.db.IsNoRowsError(err) {
			return nil, nil
		}
		return nil, err
	}

	var (
		mapper    = mapping.ExerciseMapper{}
		result, _ = mapper.FromModelToDomain(doc)
	)
	return result, nil
}

func (r ExerciseRepository) CreateExercise(ctx *appcontext.AppContext, exercise domain.Exercise) error {
	mapper := mapping.ExerciseMapper{}
	doc, err := mapper.FromDomainToModel(exercise)
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

func (r ExerciseRepository) UpdateExercise(ctx *appcontext.AppContext, exercise domain.Exercise) error {
	mapper := mapping.ExerciseMapper{}
	doc, err := mapper.FromDomainToModel(exercise)
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
