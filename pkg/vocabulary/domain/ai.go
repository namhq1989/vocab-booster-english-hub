package domain

import (
	"github.com/namhq1989/vocab-booster-english-hub/core/appcontext"
)

type AIRepository interface {
	GetVocabularyData(ctx *appcontext.AppContext, vocabulary string) (*AIVocabularyData, error)
	GrammarEvaluation(ctx *appcontext.AppContext, sentence, lang string) ([]SentenceGrammarError, error)
}

type AIVocabularyData struct {
	PosTags  []string
	IPA      string
	Synonyms []string
	Antonyms []string
	Examples []AIVocabularyExample
}

type AIVocabularyExample struct {
	Example    string
	Word       string
	Pos        string
	Definition string
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

func MappingAIPos(tag string) string {
	if mappedTag, exists := posMapping[tag]; exists {
		return mappedTag
	} else {
		return tag
	}
}
