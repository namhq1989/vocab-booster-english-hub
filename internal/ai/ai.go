package ai

import (
	"fmt"

	"github.com/namhq1989/vocab-booster-english-hub/core/appcontext"

	"github.com/sashabaranov/go-openai"
)

type Operations interface {
	VocabularyData(ctx *appcontext.AppContext, payload VocabularyDataPayload) (*VocabularyDataResult, error)
}

type AI struct {
	openai *openai.Client
}

func NewAIClient(apiKey string) *AI {
	fmt.Printf("⚡️ [ai]: OpenAI connected \n")

	return &AI{openai: openai.NewClient(apiKey)}
}
