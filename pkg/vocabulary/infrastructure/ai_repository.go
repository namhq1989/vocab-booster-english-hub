package infrastructure

import (
	"github.com/namhq1989/vocab-booster-english-hub/core/appcontext"
	"github.com/namhq1989/vocab-booster-english-hub/internal/ai"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/domain"
)

type AIRepository struct {
	ai ai.Operations
}

func NewAIRepository(ai ai.Operations) AIRepository {
	return AIRepository{
		ai: ai,
	}
}

func (r AIRepository) GetVocabularyData(ctx *appcontext.AppContext, term string) (*domain.VocabularyData, error) {
	result, err := r.ai.VocabularyData(ctx, ai.VocabularyDataPayload{
		ToLanguage: domain.LanguageVietnamese.String(),
		Vocabulary: term,
	})
	if err != nil {
		return nil, err
	}

	return &domain.VocabularyData{
		IPA:      result.IPA,
		Synonyms: result.Synonyms,
		Antonyms: result.Antonyms,
	}, nil
}

func (r AIRepository) GetVocabularyExamples(ctx *appcontext.AppContext, vocabularyID, term string) ([]domain.VocabularyExample, error) {
	result, err := r.ai.VocabularyExamples(ctx, ai.VocabularyExamplesPayload{
		Vocabulary: term,
	})
	if err != nil {
		return nil, err
	}

	examples := make([]domain.VocabularyExample, 0)
	for _, example := range result.Examples {
		var e *domain.VocabularyExample
		e, err = domain.NewVocabularyExample(vocabularyID, example.English, example.Vietnamese, example.POS, example.Definition, example.Word)
		if err != nil {
			ctx.Logger().Error("failed to create vocabulary example", err, appcontext.Fields{
				"english": example.English, "vietnamese": example.Vietnamese, "pos": example.POS, "definition": example.Definition,
			})
		} else {
			examples = append(examples, *e)
		}
	}
	return examples, nil
}
