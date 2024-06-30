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

type GrammarEvaluationPayload struct {
	Sentence string `json:"sentence"`
}

type GrammarEvaluationResult struct {
	Errors []GrammarEvaluationError `json:"errors"`
}

type GrammarEvaluationError struct {
	Message     string `json:"message"`
	Segment     string `json:"segment"`
	Replacement string `json:"replacement"`
}

const grammarEvaluationPrompt = `
	{{timestamp}}
	Check the following sentence for simple grammar errors, focusing only on grammatical errors and ignoring context or stylistic issues.
	Ignore awkward phrasing, missing periods at the end of the sentence, errors related to parallel structure, and correct possessive forms.
	Ensure that suggestions are really important and do not alter the intended meaning of the sentence.
	"{{sentence}}"
	The output should be only in JSON format without any markdown tags. Here is the required structure:
	{
      errors: [{
        "message": "detailed description including why the correction is necessary",
		"segment": "error phrase plus 1 word before and 1 word after",
		"replacement": "suggested replacement"
  	  }]
    }
`

func (ai AI) GrammarEvaluation(ctx *appcontext.AppContext, payload GrammarEvaluationPayload) (*GrammarEvaluationResult, error) {
	prompt := strings.ReplaceAll(grammarEvaluationPrompt, "{{sentence}}", payload.Sentence)
	prompt = strings.ReplaceAll(prompt, "{{timestamp}}", fmt.Sprintf("%d", time.Now().Unix()))

	// random int number
	randSource := rand.New(rand.NewSource(time.Now().UnixNano()))
	seed := randSource.Intn(10000)

	resp, err := ai.openai.CreateChatCompletion(ctx.Context(), openai.ChatCompletionRequest{
		Model: openai.GPT4o,
		// Model:       openai.GPT3Dot5Turbo1106,
		Messages:    []openai.ChatCompletionMessage{{Role: openai.ChatMessageRoleUser, Content: prompt}},
		MaxTokens:   500,
		Temperature: 0.6,
		Seed:        &seed,
	})

	if err != nil {
		return nil, err
	}

	var result GrammarEvaluationResult
	if err = json.Unmarshal([]byte(resp.Choices[0].Message.Content), &result); err != nil {
		ctx.Logger().Print("[ai] GrammarEvaluation data", resp.Choices[0].Message.Content)
		return nil, err
	}

	return &result, nil
}
