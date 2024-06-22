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

type UserExerciseStatusRepository struct {
	db             *database.Database
	collectionName string
}

func NewUserExerciseStatusRepository(db *database.Database) UserExerciseStatusRepository {
	r := UserExerciseStatusRepository{
		db:             db,
		collectionName: database.Collections.UserExerciseStatus,
	}
	r.ensureIndexes()
	return r
}

func (r UserExerciseStatusRepository) ensureIndexes() {
	var (
		ctx                  = context.Background()
		opts                 = options.CreateIndexes().SetMaxTime(time.Minute * 30)
		isUserExerciseUnique = true
		indexes              = []mongo.IndexModel{
			{
				Keys: bson.D{{Key: "userId", Value: 1}, {Key: "lastCompletedAt", Value: -1}},
			},
			{
				Keys: bson.D{{Key: "userId", Value: 1}, {Key: "exerciseId", Value: 1}},
				Options: &options.IndexOptions{
					Unique: &isUserExerciseUnique,
				},
			},
		}
	)

	if _, err := r.collection().Indexes().CreateMany(ctx, indexes, opts); err != nil {
		fmt.Printf("index collection %s err: %v \n", r.collectionName, err)
	}
}

func (r UserExerciseStatusRepository) collection() *mongo.Collection {
	return r.db.GetCollection(r.collectionName)
}

func (r UserExerciseStatusRepository) CreateUserExerciseStatus(ctx *appcontext.AppContext, status domain.UserExerciseStatus) error {
	doc, err := dbmodel.UserExerciseStatus{}.FromDomain(status)
	if err != nil {
		return err
	}

	_, err = r.collection().InsertOne(ctx.Context(), &doc)
	return err
}

func (r UserExerciseStatusRepository) UpdateUserExerciseStatus(ctx *appcontext.AppContext, status domain.UserExerciseStatus) error {
	doc, err := dbmodel.UserExerciseStatus{}.FromDomain(status)
	if err != nil {
		return err
	}

	_, err = r.collection().UpdateOne(ctx.Context(), bson.M{"_id": doc.ID}, bson.M{"$set": doc})
	return err
}

func (r UserExerciseStatusRepository) FindUserExerciseStatus(ctx *appcontext.AppContext, exerciseID, userID string) (*domain.UserExerciseStatus, error) {
	eid, err := database.ObjectIDFromString(exerciseID)
	if err != nil {
		return nil, apperrors.Exercise.InvalidExerciseID
	}

	uid, err := database.ObjectIDFromString(userID)
	if err != nil {
		return nil, apperrors.User.InvalidUserID
	}

	var doc dbmodel.UserExerciseStatus
	if err = r.collection().FindOne(ctx.Context(), bson.M{
		"userId":     uid,
		"exerciseId": eid,
	}).Decode(&doc); err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return nil, err
	} else if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, nil
	}

	result := doc.ToDomain()
	return &result, nil
}
