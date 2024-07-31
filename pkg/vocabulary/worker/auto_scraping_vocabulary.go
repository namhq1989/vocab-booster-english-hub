package worker

import (
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/domain"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type AutoScrapingVocabularyHandler struct {
	vocabularyScrapingItemRepository domain.VocabularyScrapingItemRepository
	service                          domain.Service
}

func NewAutoScrapingVocabularyHandler(
	vocabularyScrapingItemRepository domain.VocabularyScrapingItemRepository,
	service domain.Service,
) AutoScrapingVocabularyHandler {
	return AutoScrapingVocabularyHandler{
		vocabularyScrapingItemRepository: vocabularyScrapingItemRepository,
		service:                          service,
	}
}

func (w AutoScrapingVocabularyHandler) AutoScrapingVocabulary(ctx *appcontext.AppContext, _ domain.QueueAutoScrapingVocabularyPayload) error {
	ctx.Logger().Text("picking random scraping item in db")
	scrapingItem, err := w.vocabularyScrapingItemRepository.RandomPickVocabularyScrapingItem(ctx)
	if err != nil {
		ctx.Logger().Error("failed to pick random vocabulary scrape item", err, appcontext.Fields{})
		return err
	}
	if scrapingItem == nil {
		ctx.Logger().Text("no item for scraping, respond")
		return nil
	}

	ctx.Logger().Info("item found, create new vocabulary model", appcontext.Fields{"term": scrapingItem.Term})
	_, err = domain.NewVocabulary("system", scrapingItem.Term)
	if err != nil {
		ctx.Logger().Error("failed to create vocabulary model", err, appcontext.Fields{"term": scrapingItem.Term})
		return err
	}

	_, _, err = w.service.SearchVocabulary(ctx, "system", scrapingItem.Term)
	if err != nil {
		ctx.Logger().Error("failed to search vocabulary", err, appcontext.Fields{"term": scrapingItem.Term})
		return err
	}

	ctx.Logger().Text("delete scraping item")
	if err = w.vocabularyScrapingItemRepository.DeleteVocabularyScrapingItemByTerm(ctx, scrapingItem.Term); err != nil {
		ctx.Logger().Error("failed to delete scraping item", err, appcontext.Fields{})
	}

	return nil
}
