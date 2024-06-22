package infrastructure

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/namhq1989/vocab-booster-english-hub/core/appcontext"
	apperrors "github.com/namhq1989/vocab-booster-english-hub/core/error"
	"github.com/namhq1989/vocab-booster-english-hub/internal/database"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/exercise/domain"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/exercise/infrastructure/dbmodel"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ExerciseRepository struct {
	db             *database.Database
	collectionName string
}

func NewExerciseRepository(db *database.Database) ExerciseRepository {
	r := ExerciseRepository{
		db:             db,
		collectionName: database.Collections.Exercise,
	}
	r.ensureIndexes()
	return r
}

func (r ExerciseRepository) ensureIndexes() {
	var (
		ctx     = context.Background()
		opts    = options.CreateIndexes().SetMaxTime(time.Minute * 30)
		indexes = []mongo.IndexModel{
			{
				Keys: bson.D{{Key: "vocabulary", Value: 1}},
			},
			{
				Keys: bson.D{{Key: "createdAt", Value: -1}},
			},
			{
				Keys: bson.D{{Key: "vocabularyExampleId", Value: 1}},
			},
		}
	)

	if _, err := r.collection().Indexes().CreateMany(ctx, indexes, opts); err != nil {
		fmt.Printf("index collection %s err: %v \n", r.collectionName, err)
	}
}

func (r ExerciseRepository) collection() *mongo.Collection {
	return r.db.GetCollection(r.collectionName)
}

func (r ExerciseRepository) FindExerciseByID(ctx *appcontext.AppContext, exerciseID string) (*domain.Exercise, error) {
	eid, err := database.ObjectIDFromString(exerciseID)
	if err != nil {
		return nil, apperrors.Exercise.InvalidExerciseID
	}

	var doc dbmodel.Exercise
	if err = r.collection().FindOne(ctx.Context(), bson.M{
		"_id": eid,
	}).Decode(&doc); err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return nil, err
	} else if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, nil
	}

	result := doc.ToDomain()
	return &result, nil
}

func (r ExerciseRepository) FindExerciseByVocabularyExampleID(ctx *appcontext.AppContext, exampleID string) (*domain.Exercise, error) {
	eid, err := database.ObjectIDFromString(exampleID)
	if err != nil {
		return nil, apperrors.Exercise.InvalidExerciseID
	}

	var doc dbmodel.Exercise
	if err = r.collection().FindOne(ctx.Context(), bson.M{
		"vocabularyExampleId": eid,
	}).Decode(&doc); err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return nil, err
	} else if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, nil
	}

	result := doc.ToDomain()
	return &result, nil
}

func (r ExerciseRepository) CreateExercise(ctx *appcontext.AppContext, exercise domain.Exercise) error {
	doc, err := dbmodel.Exercise{}.FromDomain(exercise)
	if err != nil {
		return err
	}

	_, err = r.collection().InsertOne(ctx.Context(), &doc)
	return err
}

func (r ExerciseRepository) UpdateExercise(ctx *appcontext.AppContext, exercise domain.Exercise) error {
	doc, err := dbmodel.Exercise{}.FromDomain(exercise)
	if err != nil {
		return err
	}

	_, err = r.collection().UpdateOne(ctx.Context(), bson.M{"_id": doc.ID}, bson.M{"$set": doc})
	return err
}
