package mapping

import (
	"github.com/namhq1989/vocab-booster-english-hub/internal/database/gen/vocab-booster/public/model"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/domain"
)

type CollectionAndVocabularyMapper struct{}

func (CollectionAndVocabularyMapper) FromModelToDomain(cav model.CollectionAndVocabularies) (*domain.CollectionAndVocabulary, error) {
	return &domain.CollectionAndVocabulary{
		ID:           cav.ID,
		CollectionID: cav.CollectionID,
		VocabularyID: cav.VocabularyID,
		Value:        cav.Value,
		CreatedAt:    cav.CreatedAt,
	}, nil
}

func (CollectionAndVocabularyMapper) FromDomainToModel(cav domain.CollectionAndVocabulary) (*model.CollectionAndVocabularies, error) {
	return &model.CollectionAndVocabularies{
		ID:           cav.ID,
		CollectionID: cav.CollectionID,
		VocabularyID: cav.VocabularyID,
		Value:        cav.Value,
		CreatedAt:    cav.CreatedAt,
	}, nil
}
