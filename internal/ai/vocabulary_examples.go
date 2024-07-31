package ai

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/namhq1989/vocab-booster-utilities/appcontext"
	"github.com/sashabaranov/go-openai"
)

type VocabularyDataPayload struct {
	Vocabulary    string   `json:"vocabulary"`
	PartsOfSpeech []string `json:"partsOfSpeech"`
}

type VocabularyExample struct {
	Example string `json:"example"`
	Word    string `json:"word"`
}

const vocabularyExamplesPrompt = `
	{{timestamp}}
	The term "{{vocabulary}}" has the POS of {{pos}}. Provide 2 examples for each POS with different forms of the term and random difficulty levels.
	Ensure the "word" field contains the exact form of "{{vocabulary}}" as used in the example, including any multi-word phrases.
	The output should be only in JSON format without any markdown tags. Here is the required structure:

	[
        {
		  "example": <example sentence>,
          "word": form of '{{vocabulary}}' in example>
         }
	]
`

func (ai AI) VocabularyExamples(ctx *appcontext.AppContext, payload VocabularyDataPayload) ([]VocabularyExample, error) {
	prompt := strings.ReplaceAll(vocabularyExamplesPrompt, "{{vocabulary}}", payload.Vocabulary)
	prompt = strings.ReplaceAll(prompt, "{{timestamp}}", fmt.Sprintf("%d", time.Now().Unix()))

	// parts of speech
	var (
		pos      = ""
		totalPos = len(payload.PartsOfSpeech)
	)
	for i, v := range payload.PartsOfSpeech {
		if i == totalPos-1 {
			pos = v
			continue
		}
		pos = pos + ", " + v
	}
	prompt = strings.ReplaceAll(prompt, "{{pos}}", pos)

	// random int number
	randSource := rand.New(rand.NewSource(time.Now().UnixNano()))
	seed := randSource.Intn(10000)

	resp, err := ai.openai.CreateChatCompletion(ctx.Context(), openai.ChatCompletionRequest{
		Model:       openai.GPT4oMini,
		Messages:    []openai.ChatCompletionMessage{{Role: openai.ChatMessageRoleUser, Content: prompt}},
		MaxTokens:   1000,
		Temperature: 0.7,
		Seed:        &seed,
	})

	if err != nil {
		return nil, err
	}

	var result = make([]VocabularyExample, 0)
	if err = json.Unmarshal([]byte(resp.Choices[0].Message.Content), &result); err != nil {
		ctx.Logger().Print("[ai] VocabularyExamples data", resp.Choices[0].Message.Content)
		return nil, err
	}

	return result, nil
}
