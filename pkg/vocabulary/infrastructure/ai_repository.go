package infrastructure

import (
	"github.com/namhq1989/vocab-booster-english-hub/core/appcontext"
	"github.com/namhq1989/vocab-booster-english-hub/internal/ai"
	"github.com/namhq1989/vocab-booster-english-hub/internal/utils/manipulation"
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

func (r AIRepository) GetVocabularyData(ctx *appcontext.AppContext, vocabulary string) (*domain.AIVocabularyData, error) {
	result, err := r.ai.VocabularyData(ctx, ai.VocabularyDataPayload{
		Vocabulary: vocabulary,
	})
	if err != nil {
		return nil, err
	}

	posTags := make([]string, 0)
	for _, posTag := range result.PosTags {
		posTags = append(posTags, domain.MappingAIPos(posTag))
	}

	examples := make([]domain.AIVocabularyExample, 0)
	for _, example := range result.Examples {
		examples = append(examples, domain.AIVocabularyExample{
			Example:    manipulation.RemoveSuffix(example.Example, "."),
			Word:       example.Word,
			Pos:        domain.MappingAIPos(example.Pos),
			Definition: example.Definition,
		})
	}

	return &domain.AIVocabularyData{
		PosTags:  posTags,
		IPA:      result.IPA,
		Synonyms: result.Synonyms,
		Antonyms: result.Antonyms,
		Examples: examples,
	}, nil
}
