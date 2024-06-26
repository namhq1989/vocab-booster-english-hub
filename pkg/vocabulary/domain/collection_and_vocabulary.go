package domain

import (
	"time"

	"github.com/namhq1989/vocab-booster-english-hub/core/appcontext"
	apperrors "github.com/namhq1989/vocab-booster-english-hub/core/error"
	"github.com/namhq1989/vocab-booster-english-hub/internal/database"
)

type CollectionAndVocabularyRepository interface {
	CreateCollectionAndVocabulary(ctx *appcontext.AppContext, cav CollectionAndVocabulary) error
	DeleteCollectionAndVocabulary(ctx *appcontext.AppContext, cav CollectionAndVocabulary) error
	FindCollectionAndVocabulary(ctx *appcontext.AppContext, collectionID, vocabularyID string) (*CollectionAndVocabulary, error)
}

type CollectionAndVocabulary struct {
	ID           string
	CollectionID string
	VocabularyID string
	Value        string
	CreatedAt    time.Time
}

func NewCollectionAndVocabulary(collectionID, vocabularyID, value string) (*CollectionAndVocabulary, error) {
	if !database.IsValidID(collectionID) {
		return nil, apperrors.Collection.InvalidCollectionID
	}

	if !database.IsValidID(vocabularyID) {
		return nil, apperrors.Vocabulary.InvalidVocabularyID
	}

	if value == "" {
		return nil, apperrors.Vocabulary.InvalidTerm
	}

	return &CollectionAndVocabulary{
		ID:           database.NewStringID(),
		CollectionID: collectionID,
		VocabularyID: vocabularyID,
		Value:        value,
		CreatedAt:    time.Now(),
	}, nil
}
