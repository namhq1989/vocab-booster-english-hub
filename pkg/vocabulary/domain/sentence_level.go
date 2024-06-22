package domain

type SentenceLevel string

const (
	SentenceLevelUnknown      SentenceLevel = ""
	SentenceLevelBeginner     SentenceLevel = "beginner"
	SentenceLevelIntermediate SentenceLevel = "intermediate"
	SentenceLevelAdvanced     SentenceLevel = "advanced"
)

func (s SentenceLevel) IsValid() bool {
	return s != SentenceLevelUnknown
}

func (s SentenceLevel) String() string {
	return string(s)
}

func ToSentenceLevel(value string) SentenceLevel {
	switch value {
	case SentenceLevelBeginner.String():
		return SentenceLevelBeginner
	case SentenceLevelIntermediate.String():
		return SentenceLevelIntermediate
	case SentenceLevelAdvanced.String():
		return SentenceLevelAdvanced
	default:
		return SentenceLevelUnknown
	}
}
