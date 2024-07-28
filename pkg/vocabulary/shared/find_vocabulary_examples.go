package shared

import (
	apperrors "github.com/namhq1989/vocab-booster-english-hub/internal/utils/error"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/domain"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

func (s Service) FindVocabularyExamples(ctx *appcontext.AppContext, vocabularyID string) ([]domain.VocabularyExample, error) {
	ctx.Logger().Text("find vocabulary examples in cache")
	examples, err := s.cachingRepository.GetVocabularyExamplesByVocabularyID(ctx, vocabularyID)
	if err == nil && len(examples) > 0 {
		ctx.Logger().Text("found vocabulary examples in cache, return")
		return examples, nil
	}

	ctx.Logger().Text("find vocabulary examples in db")
	examples, err = s.vocabularyExampleRepository.FindVocabularyExamplesByVocabularyID(ctx, vocabularyID)
	if err != nil {
		ctx.Logger().Error("failed to find vocabulary examples", err, appcontext.Fields{})
		return nil, err
	}
	if len(examples) == 0 {
		return make([]domain.VocabularyExample, 0), apperrors.Vocabulary.VocabularyNotFound
	}

	ctx.Logger().Text("cache vocabulary examples")
	if err = s.cachingRepository.SetVocabularyExamplesByVocabularyID(ctx, vocabularyID, examples); err != nil {
		ctx.Logger().Error("failed to cache vocabulary examples", err, appcontext.Fields{})
	}

	return examples, nil
}
