package infrastructure

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/namhq1989/vocab-booster-english-hub/core/appcontext"
	"github.com/namhq1989/vocab-booster-english-hub/internal/database"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/domain"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/infrastructure/dbmodel"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type VocabularyScrapeItemRepository struct {
	db             *database.Database
	collectionName string
}

func NewVocabularyScrapeItemRepository(db *database.Database) VocabularyScrapeItemRepository {
	r := VocabularyScrapeItemRepository{
		db:             db,
		collectionName: database.Collections.VocabularyScrapeItem,
	}
	r.ensureIndexes()
	return r
}

func (r VocabularyScrapeItemRepository) ensureIndexes() {
	var (
		ctx          = context.Background()
		opts         = options.CreateIndexes().SetMaxTime(time.Minute * 30)
		isTermUnique = true
		indexes      = []mongo.IndexModel{
			{
				Keys: bson.D{{Key: "term", Value: 1}, {Key: "createdAt", Value: -1}},
			},
			{
				Keys: bson.D{{Key: "term", Value: 1}},
				Options: &options.IndexOptions{
					Unique: &isTermUnique,
				},
			},
		}
	)

	if _, err := r.collection().Indexes().CreateMany(ctx, indexes, opts); err != nil {
		fmt.Printf("index collection %s err: %v \n", r.collectionName, err)
	}
}

func (r VocabularyScrapeItemRepository) collection() *mongo.Collection {
	return r.db.GetCollection(r.collectionName)
}

func (r VocabularyScrapeItemRepository) FindVocabularyScrapeItemByTerm(ctx *appcontext.AppContext, term string) (*domain.VocabularyScrapeItem, error) {
	var doc dbmodel.VocabularyScrapeItem
	if err := r.collection().FindOne(ctx.Context(), bson.M{
		"term": term,
	}).Decode(&doc); err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return nil, err
	} else if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, nil
	}

	// respond
	result := doc.ToDomain()
	return &result, nil
}

func (r VocabularyScrapeItemRepository) CreateVocabularyScrapeItems(ctx *appcontext.AppContext, items []domain.VocabularyScrapeItem) error {
	writeModels := make([]mongo.WriteModel, 0, len(items))
	for _, item := range items {
		doc, err := dbmodel.VocabularyScrapeItem{}.FromDomain(item)
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

func (r VocabularyScrapeItemRepository) DeleteVocabularyScrapeItemByTerm(ctx *appcontext.AppContext, term string) error {
	_, err := r.collection().DeleteOne(ctx.Context(), bson.M{"term": term}, nil)
	return err
}
