package infrastructure

import (
	"context"
	"fmt"
	"time"

	apperrors "github.com/namhq1989/vocab-booster-english-hub/core/error"

	"github.com/namhq1989/vocab-booster-english-hub/core/appcontext"
	"github.com/namhq1989/vocab-booster-english-hub/internal/database"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/domain"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/infrastructure/dbmodel"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type VocabularyExampleRepository struct {
	db             *database.Database
	collectionName string
}

func NewVocabularyExampleRepository(db *database.Database) VocabularyExampleRepository {
	r := VocabularyExampleRepository{
		db:             db,
		collectionName: database.Collections.VocabularyExample,
	}
	r.ensureIndexes()
	return r
}

func (r VocabularyExampleRepository) ensureIndexes() {
	var (
		ctx     = context.Background()
		opts    = options.CreateIndexes().SetMaxTime(time.Minute * 30)
		indexes = []mongo.IndexModel{
			{
				Keys: bson.D{{Key: "vocabularyId", Value: 1}, {Key: "createdAt", Value: -1}},
			},
		}
	)

	if _, err := r.collection().Indexes().CreateMany(ctx, indexes, opts); err != nil {
		fmt.Printf("index collection %s err: %v \n", r.collectionName, err)
	}
}

func (r VocabularyExampleRepository) collection() *mongo.Collection {
	return r.db.GetCollection(r.collectionName)
}

func (r VocabularyExampleRepository) FindVocabularyExamplesByVocabularyID(ctx *appcontext.AppContext, vocabularyID string) ([]domain.VocabularyExample, error) {
	vid, err := database.ObjectIDFromString(vocabularyID)
	if err != nil {
		return nil, apperrors.Vocabulary.InvalidVocabularyID
	}

	var (
		condition = bson.M{"vocabularyId": vid}
		result    = make([]domain.VocabularyExample, 0)
	)

	cursor, err := r.collection().Find(ctx.Context(), condition, &options.FindOptions{
		Sort: bson.M{"createdAt": -1},
	})
	if err != nil {
		return result, err
	}
	defer func() { _ = cursor.Close(ctx.Context()) }()

	var docs []dbmodel.VocabularyExample
	if err = cursor.All(ctx.Context(), &docs); err != nil {
		return result, err
	}

	for _, doc := range docs {
		result = append(result, doc.ToDomain())
	}
	return result, nil
}

func (r VocabularyExampleRepository) CreateVocabularyExamples(ctx *appcontext.AppContext, examples []domain.VocabularyExample) error {
	writeModels := make([]mongo.WriteModel, 0, len(examples))
	for _, example := range examples {
		doc, err := dbmodel.VocabularyExample{}.FromDomain(example)
		if err != nil {
			return err
		}

		writeModels = append(writeModels, mongo.NewInsertOneModel().SetDocument(*doc))
	}

	bulkOptions := options.BulkWrite().SetOrdered(false)
	_, err := r.collection().BulkWrite(ctx.Context(), writeModels, bulkOptions)
	return err
}

func (r VocabularyExampleRepository) UpdateVocabularyExample(ctx *appcontext.AppContext, example domain.VocabularyExample) error {
	doc, err := dbmodel.VocabularyExample{}.FromDomain(example)
	if err != nil {
		return err
	}

	_, err = r.collection().UpdateByID(ctx.Context(), doc.ID, bson.M{"$set": doc})
	return err
}
