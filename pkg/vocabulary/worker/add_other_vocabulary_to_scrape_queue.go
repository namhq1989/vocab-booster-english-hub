package worker

import (
	"slices"

	"github.com/namhq1989/vocab-booster-english-hub/core/appcontext"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/domain"
)

type AddOtherVocabularyToScrapeQueueHandler struct {
	vocabularyRepository           domain.VocabularyRepository
	vocabularyScrapeItemRepository domain.VocabularyScrapeItemRepository
}

func NewAddOtherVocabularyToScrapeQueueHandler(
	vocabularyRepository domain.VocabularyRepository,
	vocabularyScrapeItemRepository domain.VocabularyScrapeItemRepository,
) AddOtherVocabularyToScrapeQueueHandler {
	return AddOtherVocabularyToScrapeQueueHandler{
		vocabularyRepository:           vocabularyRepository,
		vocabularyScrapeItemRepository: vocabularyScrapeItemRepository,
	}
}

func (w AddOtherVocabularyToScrapeQueueHandler) AddOtherVocabularyToScrapeQueue(ctx *appcontext.AppContext, payload domain.QueueAddOtherVocabularyToScrapeQueuePayload) error {
	ctx.Logger().Info("add other vocabulary to scrape queue", appcontext.Fields{"exampleID": payload.Example.ID})

	var (
		vocabulary  = make([]string, 0)
		scrapeItems = make([]domain.VocabularyScrapeItem, 0)
	)

	ctx.Logger().Text("collect vocabulary for scraping, focus on some basic POS tags only")
	for _, p := range payload.Example.PosTags {
		if slices.Contains(domain.ScrapePosTagList, p.Value) {
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

	ctx.Logger().Text("insert scrape items to db")
	if err := w.vocabularyScrapeItemRepository.CreateVocabularyScrapeItems(ctx, scrapeItems); err != nil {
		ctx.Logger().Error("failed to insert scrape items to db", err, appcontext.Fields{"scrapeItems": scrapeItems})
		return err
	}

	return nil
}

func (w AddOtherVocabularyToScrapeQueueHandler) checkAndInsertItem(ctx *appcontext.AppContext, scrapeItems *[]domain.VocabularyScrapeItem, term string) {
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
	item, err := w.vocabularyScrapeItemRepository.FindVocabularyScrapeItemByTerm(ctx, term)
	if err != nil {
		ctx.Logger().Error("failed to find scrape item", err, appcontext.Fields{"term": term})
		return
	} else if item != nil {
		ctx.Logger().Info("vocabulary scrape item is already in collection", appcontext.Fields{"term": term})
	}

	item, err = domain.NewVocabularyScrapeItem(term)
	if err != nil {
		ctx.Logger().Error("failed to create vocabulary scrape item", err, appcontext.Fields{"term": term})
		return
	}

	ctx.Logger().Text("add items to bulk documents")
	*scrapeItems = append(*scrapeItems, *item)
}
