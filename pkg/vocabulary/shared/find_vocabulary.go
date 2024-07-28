package shared

import (
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/domain"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

func (s Service) FindVocabulary(ctx *appcontext.AppContext, term string) (*domain.Vocabulary, error) {
	ctx.Logger().Text("find vocabulary in cache")
	vocabulary, err := s.cachingRepository.GetVocabularyByTerm(ctx, term)
	if err == nil && vocabulary != nil {
		ctx.Logger().Text("found vocabulary in cache, return")
		return vocabulary, nil
	}

	ctx.Logger().Text("find vocabulary in db with term")
	vocabulary, err = s.vocabularyRepository.FindVocabularyByTerm(ctx, term)
	if err != nil {
		ctx.Logger().Error("failed to find vocabulary with term", err, appcontext.Fields{})
		return nil, err
	}
	if vocabulary == nil {
		ctx.Logger().ErrorText("vocabulary not found")
		return nil, nil
	}

	ctx.Logger().Text("cache vocabulary")
	if err = s.cachingRepository.SetVocabularyByTerm(ctx, vocabulary.Term, vocabulary); err != nil {
		ctx.Logger().Error("failed to cache vocabulary", err, appcontext.Fields{})
	}

	return vocabulary, nil
}
