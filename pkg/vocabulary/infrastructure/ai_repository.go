package infrastructure

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/namhq1989/vocab-booster-english-hub/internal/ai"
	"github.com/namhq1989/vocab-booster-english-hub/internal/nlp"
	"github.com/namhq1989/vocab-booster-english-hub/internal/utils/manipulation"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/domain"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
	"github.com/namhq1989/vocab-booster-utilities/language"
)

type AIRepository struct {
	ai  ai.Operations
	nlp nlp.Operations
}

func NewAIRepository(ai ai.Operations, nlp nlp.Operations) AIRepository {
	return AIRepository{
		ai:  ai,
		nlp: nlp,
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
		if !strings.Contains(example.Example, example.Word) {
			ctx.Logger().Error("failed to find word in example", nil, appcontext.Fields{"example": example, "word": example.Word})
			continue
		}

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

func (r AIRepository) GrammarEvaluation(ctx *appcontext.AppContext, sentence string) ([]domain.SentenceGrammarError, error) {
	result, err := r.ai.GrammarEvaluation(ctx, ai.GrammarEvaluationPayload{
		Sentence: sentence,
	})
	if err != nil {
		return nil, err
	}

	grammarErrors := make([]domain.SentenceGrammarError, 0)
	for _, ge := range result.Errors {
		translated, _ := r.translateGrammarErrorMessage(ctx, ge.Message)
		if translated != nil {
			sge, _ := domain.NewSentenceGrammarError(*translated, ge.Segment, ge.Replacement)
			if sge != nil {
				grammarErrors = append(grammarErrors, *sge)
			}
		}
	}

	return grammarErrors, nil
}

func (r AIRepository) translateGrammarErrorMessage(ctx *appcontext.AppContext, message string) (*language.Multilingual, error) {
	re := regexp.MustCompile(`'[^']+'`)
	keywords := re.FindAllString(message, -1)

	placeholderText := message
	for i, keyword := range keywords {
		placeholder := fmt.Sprintf("P_%d", i)
		placeholderText = strings.Replace(placeholderText, keyword, placeholder, 1)
	}

	translatedResult, err := r.nlp.Translate(ctx, placeholderText)
	if err != nil {
		return nil, err
	}

	for i, keyword := range keywords {
		placeholder := fmt.Sprintf("P_%d", i)
		translatedResult.Vietnamese = strings.Replace(translatedResult.Vietnamese, placeholder, keyword, 1)
	}

	return translatedResult, nil
}
