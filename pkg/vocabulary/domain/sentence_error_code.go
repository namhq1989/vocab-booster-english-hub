package domain

type SentenceErrorCode string

const (
	SentenceErrorCodeEmpty             SentenceErrorCode = ""
	SentenceErrorCodeInvalidGrammar    SentenceErrorCode = "invalid_grammar"
	SentenceErrorCodeInvalidVocabulary SentenceErrorCode = "invalid_vocabulary"
	SentenceErrorCodeInvalidTense      SentenceErrorCode = "invalid_tense"
	SentenceErrorCodeIsNotEnglish      SentenceErrorCode = "is_not_english"
)

func (s SentenceErrorCode) IsEmpty() bool {
	return s == SentenceErrorCodeEmpty
}

func (s SentenceErrorCode) String() string {
	return string(s)
}

func ToSentenceErrorCode(value string) SentenceErrorCode {
	switch value {
	case SentenceErrorCodeInvalidGrammar.String():
		return SentenceErrorCodeInvalidGrammar
	case SentenceErrorCodeInvalidVocabulary.String():
		return SentenceErrorCodeInvalidVocabulary
	case SentenceErrorCodeInvalidTense.String():
		return SentenceErrorCodeInvalidTense
	case SentenceErrorCodeIsNotEnglish.String():
		return SentenceErrorCodeIsNotEnglish
	default:
		return SentenceErrorCodeEmpty
	}
}
