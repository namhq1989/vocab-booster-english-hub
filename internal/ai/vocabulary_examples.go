package ai

//
// import (
// 	"encoding/json"
// 	"fmt"
// 	"math/rand"
// 	"strings"
// 	"time"
//
// 	"github.com/namhq1989/vocab-booster-english-hub/core/appcontext"
// 	"github.com/sashabaranov/go-openai"
// )
//
// type VocabularyExamplesPayload struct {
// 	Vocabulary    string
// 	PartsOfSpeech []string
// }
//
// type VocabularyExamplesResult struct {
// 	Examples []VocabularyExample `json:"examples"`
// }
//
// type VocabularyExample struct {
// 	POS        string `json:"pos"`
// 	Definition string `json:"definition"`
// 	Example    string `json:"example"`
// }
//
// const vocabularyExamplePrompt = `
// 	{{timestamp}}
// 	Provide examples for the {{forms}} forms of the word {{vocabulary}}, each including two examples and their Vietnamese definitions.
// 	The output should be only in JSON format without any markdown tags. Here is the required structure:
//
// 	"examples": [
//       {
//         "pos": "<part of speech>",
//         "example": "<English example>",
// 		"definition": "<Vietnamese definition>"
//       }
// 	]
// `
//
// func (ai *AI) VocabularyExamples(ctx *appcontext.AppContext, payload VocabularyExamplesPayload) (*VocabularyExamplesResult, error) {
// 	prompt := strings.ReplaceAll(vocabularyExamplePrompt, "{{vocabulary}}", payload.Vocabulary)
// 	prompt = strings.ReplaceAll(prompt, "{{timestamp}}", fmt.Sprintf("%d", time.Now().Unix()))
//
// 	// parts of speech
// 	partsOfSpeech := strings.Join(payload.PartsOfSpeech, " and ")
// 	prompt = strings.ReplaceAll(prompt, "{{forms}}", partsOfSpeech)
//
// 	// random int number
// 	randSource := rand.New(rand.NewSource(time.Now().UnixNano()))
// 	seed := randSource.Intn(10000)
//
// 	resp, err := ai.openai.CreateChatCompletion(ctx.Context(), openai.ChatCompletionRequest{
// 		Model:       openai.GPT3Dot5Turbo1106,
// 		Messages:    []openai.ChatCompletionMessage{{Role: openai.ChatMessageRoleUser, Content: prompt}},
// 		MaxTokens:   700,
// 		Temperature: 0.7,
// 		Seed:        &seed,
// 	})
//
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	var result VocabularyExamplesResult
// 	if err = json.Unmarshal([]byte(resp.Choices[0].Message.Content), &result); err != nil {
// 		ctx.Logger().Print("[ai] VocabularyExamples data", resp.Choices[0].Message.Content)
// 		return nil, err
// 	}
//
// 	return &result, nil
// }
