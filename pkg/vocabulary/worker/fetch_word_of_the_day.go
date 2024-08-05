package worker

import (
	apperrors "github.com/namhq1989/vocab-booster-english-hub/internal/utils/error"
	"github.com/namhq1989/vocab-booster-english-hub/internal/utils/manipulation"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/domain"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
	"github.com/namhq1989/vocab-booster-utilities/language"
)

type FetchWordOfTheDayHandler struct {
	wordOfTheDayRepository domain.WordOfTheDayRepository
	aiRepository           domain.AIRepository
	cachingRepository      domain.CachingRepository
	service                domain.Service
}

func NewFetchWordOfTheDayHandler(
	wordOfTheDayRepository domain.WordOfTheDayRepository,
	aiRepository domain.AIRepository,
	cachingRepository domain.CachingRepository,
	service domain.Service,
) FetchWordOfTheDayHandler {
	return FetchWordOfTheDayHandler{
		wordOfTheDayRepository: wordOfTheDayRepository,
		aiRepository:           aiRepository,
		cachingRepository:      cachingRepository,
		service:                service,
	}
}

var supportedCountries = []string{language.English.GetCountry(), language.Vietnamese.GetCountry()}

func (w FetchWordOfTheDayHandler) FetchWordOfTheDay(ctx *appcontext.AppContext, _ domain.QueueFetchWordOfTheDayPayload) error {
	date := manipulation.NowUTC().Format("02/01")

	for _, country := range supportedCountries {
		ctx.Logger().Info("fetching word of the day", appcontext.Fields{"country": country, "date": date})
		wotd, err := w.aiRepository.WordOfTheDay(ctx, country, date)
		if err != nil {
			ctx.Logger().Error("failed to create word of the day", err, appcontext.Fields{"country": country, "date": date})
			return err
		}

		ctx.Logger().Info("search vocabulary with word of the day", appcontext.Fields{"term": wotd.Word})
		vocabulary, _, err := w.service.SearchVocabulary(ctx, "system", wotd.Word)
		if err != nil {
			ctx.Logger().Error("failed to search vocabulary", err, appcontext.Fields{"term": wotd.Word})
			return err
		}
		if vocabulary == nil {
			ctx.Logger().ErrorText("vocabulary not found")
			return apperrors.Vocabulary.VocabularyNotFound
		}

		ctx.Logger().Text("create word of the day model")
		word, err := domain.NewWordOfTheDay(vocabulary.ID, country, wotd.Information)
		if err != nil {
			ctx.Logger().Error("failed to create word of the day", err, appcontext.Fields{"country": country, "date": date})
			return err
		}

		ctx.Logger().Text("persist word of the day in db")
		if err = w.wordOfTheDayRepository.CreateWordOfTheDay(ctx, *word); err != nil {
			ctx.Logger().Error("failed to persist word of the day", err, appcontext.Fields{"country": country, "date": date})
			return err
		}

		ctx.Logger().Text("delete word of the day in caching layer")
		if err = w.cachingRepository.DeleteWordOfTheDay(ctx, country); err != nil {
			ctx.Logger().Error("failed to delete word of the day in caching layer", err, appcontext.Fields{"country": country})
		}
	}

	return nil
}
