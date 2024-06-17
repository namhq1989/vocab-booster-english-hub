package infrastructure

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/namhq1989/vocab-booster-english-hub/core/appcontext"
	apperrors "github.com/namhq1989/vocab-booster-english-hub/core/error"
	"github.com/namhq1989/vocab-booster-english-hub/internal/database"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/domain"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/infrastructure/dbmodel"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type VerbConjugationRepository struct {
	db             *database.Database
	collectionName string
}

func NewVerbConjugationRepository(db *database.Database) VerbConjugationRepository {
	r := VerbConjugationRepository{
		db:             db,
		collectionName: database.Collections.VerbConjugation,
	}
	r.ensureIndexes()
	return r
}

func (r VerbConjugationRepository) ensureIndexes() {
	var (
		ctx               = context.Background()
		opts              = options.CreateIndexes().SetMaxTime(time.Minute * 30)
		isVerbValueUnique = true
		indexes           = []mongo.IndexModel{
			{
				Keys: bson.D{{Key: "value", Value: 1}, {Key: "form", Value: 1}},
				Options: &options.IndexOptions{
					Unique: &isVerbValueUnique,
				},
			},
		}
	)

	if _, err := r.collection().Indexes().CreateMany(ctx, indexes, opts); err != nil {
		fmt.Printf("index collection %s err: %v \n", r.collectionName, err)
	}
}

func (r VerbConjugationRepository) collection() *mongo.Collection {
	return r.db.GetCollection(r.collectionName)
}

func (r VerbConjugationRepository) FindVerbConjugationByValue(ctx *appcontext.AppContext, value string) (*domain.VerbConjugation, error) {
	var doc dbmodel.VerbConjugation
	if err := r.collection().FindOne(ctx.Context(), bson.M{
		"value": value,
	}).Decode(&doc); err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return nil, err
	} else if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, nil
	}

	// respond
	result := doc.ToDomain()
	return &result, nil
}

func (r VerbConjugationRepository) CreateVerbConjugations(ctx *appcontext.AppContext, verbs []domain.VerbConjugation) error {
	writeModels := make([]mongo.WriteModel, 0, len(verbs))
	for _, verb := range verbs {
		doc, err := dbmodel.VerbConjugation{}.FromDomain(verb)
		if err != nil {
			return err
		}

		writeModels = append(writeModels, mongo.NewInsertOneModel().SetDocument(*doc))
	}

	bulkOptions := options.BulkWrite().SetOrdered(false)
	_, err := r.collection().BulkWrite(ctx.Context(), writeModels, bulkOptions)
	if err != nil && mongo.IsDuplicateKeyError(err) {
		return nil
	}
	return err
}

func (r VerbConjugationRepository) FindVerbConjugationByVocabularyID(ctx *appcontext.AppContext, vocabularyID string) ([]domain.VerbConjugation, error) {
	vid, err := database.ObjectIDFromString(vocabularyID)
	if err != nil {
		return nil, apperrors.Vocabulary.InvalidVocabularyID
	}

	var (
		condition = bson.M{"vocabularyId": vid}
		result    = make([]domain.VerbConjugation, 0)
	)

	cursor, err := r.collection().Find(ctx.Context(), condition, nil)
	if err != nil {
		return result, err
	}
	defer func() { _ = cursor.Close(ctx.Context()) }()

	var docs []dbmodel.VerbConjugation
	if err = cursor.All(ctx.Context(), &docs); err != nil {
		return result, err
	}

	for _, doc := range docs {
		result = append(result, doc.ToDomain())
	}
	return result, nil
}
