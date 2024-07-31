package ai

import (
	"fmt"

	"github.com/namhq1989/vocab-booster-utilities/appcontext"
	"github.com/sashabaranov/go-openai"
)

type Operations interface {
	VocabularyExamples(ctx *appcontext.AppContext, payload VocabularyDataPayload) ([]VocabularyExample, error)
	GrammarEvaluation(ctx *appcontext.AppContext, payload GrammarEvaluationPayload) (*GrammarEvaluationResult, error)
	WordOfTheDay(ctx *appcontext.AppContext, payload WordOfTheDayPayload) (*WordOfTheDayResult, error)
}

type AI struct {
	openai *openai.Client
}

func NewAIClient(apiKey string) *AI {
	fmt.Printf("⚡️ [ai]: OpenAI connected \n")

	return &AI{openai: openai.NewClient(apiKey)}
}
