package mapping

import (
	"github.com/namhq1989/vocab-booster-english-hub/internal/database/gen/vocab-booster/public/model"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/domain"
)

type VocabularyScrapingItemMapper struct{}

func (VocabularyScrapingItemMapper) FromModelToDomain(item model.VocabularyScrapingItems) (*domain.VocabularyScrapingItem, error) {
	return &domain.VocabularyScrapingItem{
		ID:        item.ID,
		Term:      item.Term,
		CreatedAt: item.CreatedAt,
	}, nil
}

func (VocabularyScrapingItemMapper) FromDomainToModel(item domain.VocabularyScrapingItem) (*model.VocabularyScrapingItems, error) {
	return &model.VocabularyScrapingItems{
		ID:        item.ID,
		Term:      item.Term,
		CreatedAt: item.CreatedAt,
	}, nil
}
