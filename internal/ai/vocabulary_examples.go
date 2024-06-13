package ai

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/namhq1989/vocab-booster-english-hub/core/appcontext"
	"github.com/sashabaranov/go-openai"
)

type VocabularyExamplesPayload struct {
	Vocabulary string `json:"vocabulary"`
}

type VocabularyExamplesResult struct {
	Examples []VocabularyExample `json:"examples"`
}

type VocabularyExample struct {
	English    string `json:"en"`
	Vietnamese string `json:"vi"`
	POS        string `json:"pos"`
	Definition string `json:"definition"`
	Word       string `json:"word"`
}

const vocabularyExamplePrompt = `
	{{timestamp}}
	Listing all possible parts of speech of "{{vocabulary}}".
	Provide 1 example for each usage along with their translations into Vietnamese.
	Each English example must have at least 50 characters, include the current word, the definition (in Vietnamese) and the part of speech of current word in the sentence.
	The output should be only in JSON format without any markdown tags. Here is the required structure:

	"examples": [
      {
        "en": "<English example>",
        "vi": "<Vietnamese translation>",
		"word": "<word>",
		"definition": "<Vietnamese definition>",
        "pos": "<part of speech>"
      }
	]
`

func (ai *AI) VocabularyExamples(ctx *appcontext.AppContext, payload VocabularyExamplesPayload) (*VocabularyExamplesResult, error) {
	prompt := strings.ReplaceAll(vocabularyExamplePrompt, "{{vocabulary}}", payload.Vocabulary)
	prompt = strings.ReplaceAll(prompt, "{{timestamp}}", fmt.Sprintf("%d", time.Now().Unix()))

	// random int number
	randSource := rand.New(rand.NewSource(time.Now().UnixNano()))
	seed := randSource.Intn(10000)

	resp, err := ai.openai.CreateChatCompletion(ctx.Context(), openai.ChatCompletionRequest{
		Model:       openai.GPT3Dot5Turbo1106,
		Messages:    []openai.ChatCompletionMessage{{Role: openai.ChatMessageRoleUser, Content: prompt}},
		MaxTokens:   700,
		Temperature: 0.7,
		Seed:        &seed,
	})

	if err != nil {
		return nil, err
	}

	var result VocabularyExamplesResult
	if err = json.Unmarshal([]byte(resp.Choices[0].Message.Content), &result); err != nil {
		ctx.Logger().Print("[ai] VocabularyExamples data", resp.Choices[0].Message.Content)
		return nil, err
	}

	return &result, nil
}
