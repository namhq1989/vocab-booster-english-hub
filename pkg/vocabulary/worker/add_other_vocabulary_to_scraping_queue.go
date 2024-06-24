package worker

import (
	"slices"

	"github.com/namhq1989/vocab-booster-english-hub/core/appcontext"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/domain"
)

type AddOtherVocabularyToScrapingQueueHandler struct {
	vocabularyRepository             domain.VocabularyRepository
	VocabularyScrapingItemRepository domain.VocabularyScrapingItemRepository
}

func NewAddOtherVocabularyToScrapingQueueHandler(
	vocabularyRepository domain.VocabularyRepository,
	VocabularyScrapingItemRepository domain.VocabularyScrapingItemRepository,
) AddOtherVocabularyToScrapingQueueHandler {
	return AddOtherVocabularyToScrapingQueueHandler{
		vocabularyRepository:             vocabularyRepository,
		VocabularyScrapingItemRepository: VocabularyScrapingItemRepository,
	}
}

func (w AddOtherVocabularyToScrapingQueueHandler) AddOtherVocabularyToScrapingQueue(ctx *appcontext.AppContext, payload domain.QueueAddOtherVocabularyToScrapingQueuePayload) error {
	ctx.Logger().Info("add other vocabulary to scraping queue", appcontext.Fields{"exampleID": payload.Example.ID})

	var (
		vocabulary  = make([]string, 0)
		scrapeItems = make([]domain.VocabularyScrapingItem, 0)
	)

	ctx.Logger().Text("collect vocabulary for scraping, focus on some basic POS tags only")
	for _, p := range payload.Example.PosTags {
		if p.Level >= 3 && slices.Contains(domain.ScrapingPosTagList, p.Value) {
			vocabulary = append(vocabulary, p.Word)
		}
	}

	ctx.Logger().Text("add verbs with base forms to vocabulary collection")
	for _, v := range payload.Example.Verbs {
		vocabulary = append(vocabulary, v.Base)
	}

	ctx.Logger().Text("collect vocabulary can be scraped")
	for _, v := range vocabulary {
		w.checkAndInsertItem(ctx, &scrapeItems, v)
	}

	if len(scrapeItems) == 0 {
		ctx.Logger().Text("no vocabulary to scrape")
		return nil
	}

	ctx.Logger().Text("insert scrape items in db")
	if err := w.VocabularyScrapingItemRepository.CreateVocabularyScrapingItems(ctx, scrapeItems); err != nil {
		ctx.Logger().Error("failed to insert scrape items in db", err, appcontext.Fields{"scrapeItems": scrapeItems})
		return err
	}

	return nil
}

func (w AddOtherVocabularyToScrapingQueueHandler) checkAndInsertItem(ctx *appcontext.AppContext, scrapeItems *[]domain.VocabularyScrapingItem, term string) {
	ctx.Logger().Info("check and insert item", appcontext.Fields{"term": term})

	ctx.Logger().Text("check in vocabulary collection first")
	vocabulary, err := w.vocabularyRepository.FindVocabularyByTerm(ctx, term)
	if err != nil {
		ctx.Logger().Error("failed to find vocabulary", err, appcontext.Fields{"term": term})
		return
	} else if vocabulary != nil {
		ctx.Logger().Info("vocabulary is already in collection", appcontext.Fields{"term": term})
		return
	}

	ctx.Logger().Text("check in scrape item collection")
	item, err := w.VocabularyScrapingItemRepository.FindVocabularyScrapingItemByTerm(ctx, term)
	if err != nil {
		ctx.Logger().Error("failed to find scrape item", err, appcontext.Fields{"term": term})
		return
	} else if item != nil {
		ctx.Logger().Info("vocabulary scrape item is already in collection", appcontext.Fields{"term": term})
	}

	item, err = domain.NewVocabularyScrapingItem(term)
	if err != nil {
		ctx.Logger().Error("failed to create vocabulary scrape item", err, appcontext.Fields{"term": term})
		return
	}

	ctx.Logger().Text("add items to bulk documents")
	*scrapeItems = append(*scrapeItems, *item)
}
