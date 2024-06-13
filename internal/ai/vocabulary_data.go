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

type VocabularyDataPayload struct {
	ToLanguage string `json:"toLanguage"`
	Vocabulary string `json:"vocabulary"`
}

type VocabularyDataResult struct {
	IPA      string   `json:"ipa"`
	Synonyms []string `json:"synonyms"`
	Antonyms []string `json:"antonyms"`
}

const vocabularyDataPrompt = `
	{{timestamp}}
	Generate a JSON structure for "{{vocabulary}}" with:
	- IPA transcription
	- Up to 3 strong matches synonyms
	- Up to 3 strong matches antonyms
	The output should be only in JSON format without any markdown tags. Here is the required structure:

	{
      "ipa": "<ipa>",
	  "synonyms": <list synonyms>,
      "antonyms": <list antonyms>
	}
`

func (ai *AI) VocabularyData(ctx *appcontext.AppContext, payload VocabularyDataPayload) (*VocabularyDataResult, error) {
	prompt := strings.ReplaceAll(vocabularyDataPrompt, "{{vocabulary}}", payload.Vocabulary)
	prompt = strings.ReplaceAll(prompt, "{{toLanguage}}", payload.ToLanguage)
	prompt = strings.ReplaceAll(prompt, "{{timestamp}}", fmt.Sprintf("%d", time.Now().Unix()))

	// random int number
	randSource := rand.New(rand.NewSource(time.Now().UnixNano()))
	seed := randSource.Intn(10000)

	resp, err := ai.openai.CreateChatCompletion(ctx.Context(), openai.ChatCompletionRequest{
		Model:       openai.GPT3Dot5Turbo1106,
		Messages:    []openai.ChatCompletionMessage{{Role: openai.ChatMessageRoleUser, Content: prompt}},
		MaxTokens:   400,
		Temperature: 0.5,
		Seed:        &seed,
	})

	if err != nil {
		return nil, err
	}

	var result VocabularyDataResult
	if err = json.Unmarshal([]byte(resp.Choices[0].Message.Content), &result); err != nil {
		ctx.Logger().Print("[ai] VocabularyData data", resp.Choices[0].Message.Content)
		return nil, err
	}

	return &result, nil
}
