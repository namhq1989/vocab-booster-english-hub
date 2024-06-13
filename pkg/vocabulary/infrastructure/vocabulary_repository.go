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

type VocabularyRepository struct {
	db             *database.Database
	collectionName string
}

func NewVocabularyRepository(db *database.Database) VocabularyRepository {
	r := VocabularyRepository{
		db:             db,
		collectionName: database.Collections.Vocabulary,
	}
	r.ensureIndexes()
	return r
}

func (r VocabularyRepository) ensureIndexes() {
	var (
		ctx     = context.Background()
		opts    = options.CreateIndexes().SetMaxTime(time.Minute * 30)
		indexes = []mongo.IndexModel{
			{
				Keys: bson.D{{Key: "term", Value: 1}, {Key: "createdAt", Value: -1}},
			},
		}
	)

	if _, err := r.collection().Indexes().CreateMany(ctx, indexes, opts); err != nil {
		fmt.Printf("index collection %s err: %v \n", r.collectionName, err)
	}
}

func (r VocabularyRepository) collection() *mongo.Collection {
	return r.db.GetCollection(r.collectionName)
}

func (r VocabularyRepository) FindVocabularyByID(ctx *appcontext.AppContext, vocabularyID string) (*domain.Vocabulary, error) {
	id, err := database.ObjectIDFromString(vocabularyID)
	if err != nil {
		return nil, apperrors.Common.InvalidID
	}

	// find
	var doc dbmodel.Vocabulary
	if err = r.collection().FindOne(ctx.Context(), bson.M{
		"_id": id,
	}).Decode(&doc); err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return nil, err
	} else if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, nil
	}

	// respond
	result := doc.ToDomain()
	return &result, nil
}

func (r VocabularyRepository) FindVocabularyByTerm(ctx *appcontext.AppContext, term string) (*domain.Vocabulary, error) {
	var doc dbmodel.Vocabulary
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

func (r VocabularyRepository) CreateVocabulary(ctx *appcontext.AppContext, vocabulary domain.Vocabulary) error {
	doc, err := dbmodel.Vocabulary{}.FromDomain(vocabulary)
	if err != nil {
		return err
	}

	_, err = r.collection().InsertOne(ctx.Context(), &doc)
	return err
}

func (r VocabularyRepository) UpdateVocabulary(ctx *appcontext.AppContext, vocabulary domain.Vocabulary) error {
	doc, err := dbmodel.Vocabulary{}.FromDomain(vocabulary)
	if err != nil {
		return err
	}

	_, err = r.collection().UpdateByID(ctx.Context(), doc.ID, bson.M{"$set": doc})
	return err
}
