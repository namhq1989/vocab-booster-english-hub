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
	Vocabulary string `json:"vocabulary"`
}

type VocabularyDataResult struct {
	PosTags  []string            `json:"posTags"`
	IPA      string              `json:"ipa"`
	Synonyms []string            `json:"synonyms"`
	Antonyms []string            `json:"antonyms"`
	Examples []VocabularyExample `json:"examples"`
}

type VocabularyExample struct {
	Example    string `json:"example"`
	Word       string `json:"word"`
	Pos        string `json:"pos"`
	Definition string `json:"definition"`
}

const vocabularyDataPrompt = `
	{{timestamp}}
	Generate a JSON structure for "{{vocabulary}}" with:
	- An array of possible parts of speech (POS) in spaCy's format and lowercase.
	- IPA transcription with periods to indicate syllable breaks
	- Random 3-5 strong matches synonyms
	- Random 3-5 strong matches antonyms
    - For each POS, provide 2 examples with random difficulty levels:
      + beginner (A simple sentence with basic vocabulary and structure)
      + intermediate (A moderately complex sentence with some advanced vocabulary and structure)
      + advanced (A complex sentence with advanced vocabulary and intricate structure).
    Ensure the "definition" field contains only the English translation of the word form, not the entire sentence.
    The "word" field should contain the exact form of "{{vocabulary}}" as used in the example, including any multi-word phrases.
	The output should be only in JSON format without any markdown tags. Here is the required structure:

	{
      "posTags": "<list POS in spaCy's format>",
      "ipa": "<ipa>",
	  "synonyms": <list synonyms>,
      "antonyms": <list antonyms>,
	  "examples": [
        {
		  "example": "<English example>",
          "word": "<form of '{{vocabulary}}' in example>"
		  "pos": "<part of speech>",
		  "definition": "<English definition of the word form>"
        }
      ]
	}
`

func (ai AI) VocabularyData(ctx *appcontext.AppContext, payload VocabularyDataPayload) (*VocabularyDataResult, error) {
	prompt := strings.ReplaceAll(vocabularyDataPrompt, "{{vocabulary}}", payload.Vocabulary)
	prompt = strings.ReplaceAll(prompt, "{{timestamp}}", fmt.Sprintf("%d", time.Now().Unix()))

	// random int number
	randSource := rand.New(rand.NewSource(time.Now().UnixNano()))
	seed := randSource.Intn(10000)

	resp, err := ai.openai.CreateChatCompletion(ctx.Context(), openai.ChatCompletionRequest{
		// Model:       openai.GPT4o,
		Model:       openai.GPT3Dot5Turbo1106,
		Messages:    []openai.ChatCompletionMessage{{Role: openai.ChatMessageRoleUser, Content: prompt}},
		MaxTokens:   700,
		Temperature: 0.8,
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
