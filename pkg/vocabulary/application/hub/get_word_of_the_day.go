package hub

import (
	"github.com/namhq1989/vocab-booster-english-hub/internal/genproto/vocabularypb"
	apperrors "github.com/namhq1989/vocab-booster-english-hub/internal/utils/error"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/domain"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/dto"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
	"github.com/namhq1989/vocab-booster-utilities/language"
)

type GetWordOfTheDayHandler struct {
	wordOfTheDayRepository domain.WordOfTheDayRepository
	cachingRepository      domain.CachingRepository
}

func NewGetWordOfTheDayHandler(wordOfTheDayRepository domain.WordOfTheDayRepository, cachingRepository domain.CachingRepository) GetWordOfTheDayHandler {
	return GetWordOfTheDayHandler{
		wordOfTheDayRepository: wordOfTheDayRepository,
		cachingRepository:      cachingRepository,
	}
}

func (h GetWordOfTheDayHandler) GetWordOfTheDay(ctx *appcontext.AppContext, req *vocabularypb.GetWordOfTheDayRequest) (*vocabularypb.GetWordOfTheDayResponse, error) {
	ctx.Logger().Info("[hub] new get word of the day request", appcontext.Fields{"lang": req.GetLang()})

	ctx.Logger().Text("get country from lang")
	lang := language.ToLanguage(req.GetLang())
	country := lang.GetCountry()

	ctx.Logger().Text("find word of the day in caching layer")
	wotd, err := h.cachingRepository.GetWordOfTheDay(ctx, country)
	if err != nil {
		ctx.Logger().Error("failed to find word of the day in caching layer", err, appcontext.Fields{})
		return nil, err
	}
	if wotd != nil {
		ctx.Logger().Text("word of the day found in caching layer")
	} else {
		ctx.Logger().Text("word of the day not found in caching layer, find in db")
		wotd, err = h.wordOfTheDayRepository.FindWordOfTheDay(ctx, country)
		if err != nil {
			ctx.Logger().Error("failed to find word of the day in db", err, appcontext.Fields{})
			return nil, err
		}
		if wotd == nil {
			ctx.Logger().ErrorText("word of the day not found")
			return nil, apperrors.Vocabulary.VocabularyNotFound
		}

		ctx.Logger().Text("set word of the day in caching layer")
		err = h.cachingRepository.SetWordOfTheDay(ctx, *wotd, country)
		if err != nil {
			ctx.Logger().Error("failed to set word of the day in caching layer", err, appcontext.Fields{})
		}
	}

	ctx.Logger().Text("convert to response data")
	result := &vocabularypb.GetWordOfTheDayResponse{
		Vocabulary:  dto.ConvertVocabularyBriefFromDomainToGrpc(wotd.Vocabulary),
		Information: dto.ConvertMultilingualToGrpcData(wotd.WordOfTheDay.Information),
	}

	ctx.Logger().Text("done get word of the day request")
	return result, nil
}
