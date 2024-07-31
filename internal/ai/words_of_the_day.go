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

type WordOfTheDayPayload struct {
	Country string
	Date    string
}

type WordOfTheDayResult struct {
	Word        string `json:"word"`
	Information string `json:"information"`
}

const wordOfTheDayPrompt = `
	{{timestamp}}
	Provide 1 vocabulary word that are related to significant historical events that happened {{country}} on {{date}} (DD/MM) in any year,
	including a brief context or explanation of the related event
	Focus on more recent events (within the last century).
	The output should be only in JSON format without any markdown tags. Here is the required structure:
	{
	  "word": "",
	  "information": ""
    }
`

func (ai AI) WordOfTheDay(ctx *appcontext.AppContext, payload WordOfTheDayPayload) (*WordOfTheDayResult, error) {
	prompt := strings.ReplaceAll(wordOfTheDayPrompt, "{{date}}", payload.Date)

	if payload.Country != "" {
		prompt = strings.ReplaceAll(prompt, "{{country}}", "in "+payload.Country)
	} else {
		prompt = strings.ReplaceAll(prompt, "{{country}}", "")
	}
	prompt = strings.ReplaceAll(prompt, "{{timestamp}}", fmt.Sprintf("%d", time.Now().Unix()))

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

	var result = &WordOfTheDayResult{}
	if err = json.Unmarshal([]byte(resp.Choices[0].Message.Content), &result); err != nil {
		ctx.Logger().Print("[ai] WordOfTheDay data", resp.Choices[0].Message.Content)
		return nil, err
	}

	return result, nil
}
