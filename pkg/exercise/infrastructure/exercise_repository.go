package infrastructure

import (
	"context"
	"fmt"
	"time"

	"github.com/namhq1989/vocab-booster-english-hub/pkg/exercise/domain"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/exercise/infrastructure/dbmodel"

	"github.com/namhq1989/vocab-booster-english-hub/core/appcontext"
	"github.com/namhq1989/vocab-booster-english-hub/internal/database"
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
		}
	)

	if _, err := r.collection().Indexes().CreateMany(ctx, indexes, opts); err != nil {
		fmt.Printf("index collection %s err: %v \n", r.collectionName, err)
	}
}

func (r ExerciseRepository) collection() *mongo.Collection {
	return r.db.GetCollection(r.collectionName)
}

func (r ExerciseRepository) CreateExercise(ctx *appcontext.AppContext, exercise domain.Exercise) error {
	doc, err := dbmodel.Exercise{}.FromDomain(exercise)
	if err != nil {
		return err
	}

	_, err = r.collection().InsertOne(ctx.Context(), &doc)
	return err
}
