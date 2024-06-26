package mapping

import (
	"github.com/namhq1989/vocab-booster-english-hub/internal/database/gen/vocab-booster/public/model"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/domain"
)

type UserVocabularyCollectionMapper struct{}

func (UserVocabularyCollectionMapper) FromModelToDomain(collection model.UserVocabularyCollections) (*domain.UserVocabularyCollection, error) {
	return &domain.UserVocabularyCollection{
		ID:              collection.ID,
		UserID:          collection.UserID,
		Name:            collection.Name,
		NumOfVocabulary: int(collection.NumOfVocabulary),
		CreatedAt:       collection.CreatedAt,
	}, nil
}

func (UserVocabularyCollectionMapper) FromDomainToModel(collection domain.UserVocabularyCollection) (*model.UserVocabularyCollections, error) {
	return &model.UserVocabularyCollections{
		ID:              collection.ID,
		UserID:          collection.UserID,
		Name:            collection.Name,
		NumOfVocabulary: int32(collection.NumOfVocabulary),
		CreatedAt:       collection.CreatedAt,
	}, nil
}
