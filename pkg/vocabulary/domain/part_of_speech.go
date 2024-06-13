package domain

import "strings"

type PartOfSpeech string

const (
	// Parts of speech from ChatGPT

	PartOfSpeechUnknown      PartOfSpeech = ""
	PartOfSpeechNoun         PartOfSpeech = "noun"
	PartOfSpeechPronoun      PartOfSpeech = "pronoun"
	PartOfSpeechVerb         PartOfSpeech = "verb"
	PartOfSpeechAdjective    PartOfSpeech = "adjective"
	PartOfSpeechAdverb       PartOfSpeech = "adverb"
	PartOfSpeechConjunction  PartOfSpeech = "conjunction"
	PartOfSpeechInterjection PartOfSpeech = "interjection"
	PartOfSpeechPreposition  PartOfSpeech = "preposition"
	PartOfSpeechDeterminer   PartOfSpeech = "determiner"
	PartOfSpeechArticle      PartOfSpeech = "article"
	PartOfSpeechNumeral      PartOfSpeech = "numeral"
)

func (s PartOfSpeech) String() string {
	switch s {
	case PartOfSpeechNoun, PartOfSpeechPronoun, PartOfSpeechVerb, PartOfSpeechAdjective, PartOfSpeechAdverb,
		PartOfSpeechConjunction, PartOfSpeechInterjection, PartOfSpeechPreposition,
		PartOfSpeechDeterminer, PartOfSpeechArticle, PartOfSpeechNumeral:
		return string(s)
	default:
		return ""
	}
}

func (s PartOfSpeech) IsValid() bool {
	return s != PartOfSpeechUnknown
}

func ToPartOfSpeech(value string) PartOfSpeech {
	switch strings.ToLower(value) {
	case PartOfSpeechNoun.String():
		return PartOfSpeechNoun
	case PartOfSpeechPronoun.String():
		return PartOfSpeechPronoun
	case PartOfSpeechVerb.String():
		return PartOfSpeechVerb
	case PartOfSpeechAdjective.String():
		return PartOfSpeechAdjective
	case PartOfSpeechAdverb.String():
		return PartOfSpeechAdverb
	case PartOfSpeechConjunction.String():
		return PartOfSpeechConjunction
	case PartOfSpeechInterjection.String():
		return PartOfSpeechInterjection
	case PartOfSpeechPreposition.String():
		return PartOfSpeechPreposition
	case PartOfSpeechDeterminer.String():
		return PartOfSpeechDeterminer
	case PartOfSpeechArticle.String():
		return PartOfSpeechArticle
	case PartOfSpeechNumeral.String():
		return PartOfSpeechNumeral
	default:
		return PartOfSpeechUnknown
	}
}
