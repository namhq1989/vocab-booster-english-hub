package domain

import (
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type AIRepository interface {
	VocabularyExamples(ctx *appcontext.AppContext, vocabulary string, partsOfSpeech []string) ([]AIVocabularyExample, error)
	GrammarEvaluation(ctx *appcontext.AppContext, sentence string) ([]SentenceGrammarError, error)
}

type AIVocabularyExample struct {
	Example string
	Word    string
}

var posMapping = map[string]string{
	"adjective":    "adj",
	"adj":          "adj",
	"noun":         "noun",
	"verb":         "verb",
	"adverb":       "adv",
	"adv":          "adv",
	"pronoun":      "pron",
	"preposition":  "adp",
	"adp":          "adp",
	"conjunction":  "cconj",
	"cconj":        "cconj",
	"determiner":   "det",
	"det":          "det",
	"exclamation":  "intj",
	"interjection": "intj",
	"intj":         "intj",
	"numeral":      "num",
	"num":          "num",
	"particle":     "part",
	"part":         "part",
	"proper noun":  "propn",
	"propn":        "propn",
	"punctuation":  "punct",
	"punct":        "punct",
	"symbol":       "sym",
	"sym":          "sym",
	"x":            "x",
}

func MappingPos(tag string) string {
	if mappedTag, exists := posMapping[tag]; exists {
		return mappedTag
	} else {
		return tag
	}
}
