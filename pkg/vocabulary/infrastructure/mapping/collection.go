package mapping

import (
	"github.com/namhq1989/vocab-booster-english-hub/internal/database/gen/vocab-booster/public/model"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/domain"
)

type CollectionMapper struct{}

func (CollectionMapper) FromModelToDomain(collection model.Collections) (*domain.Collection, error) {
	return &domain.Collection{
		ID:              collection.ID,
		UserID:          collection.UserID,
		Name:            collection.Name,
		Description:     collection.Description,
		NumOfVocabulary: int(collection.NumOfVocabulary),
		CreatedAt:       collection.CreatedAt,
		UpdatedAt:       collection.UpdatedAt,
	}, nil
}

func (CollectionMapper) FromDomainToModel(collection domain.Collection) (*model.Collections, error) {
	return &model.Collections{
		ID:              collection.ID,
		UserID:          collection.UserID,
		Name:            collection.Name,
		Description:     collection.Description,
		NumOfVocabulary: int32(collection.NumOfVocabulary),
		CreatedAt:       collection.CreatedAt,
		UpdatedAt:       collection.UpdatedAt,
	}, nil
}
