package mapping

import "github.com/namhq1989/vocab-booster-utilities/language"

type PosTag struct {
	Word  string `json:"word"`
	Value string `json:"value"`
	Level int    `json:"level"`
}

type Sentiment struct {
	Polarity     float64 `json:"polarity"`
	Subjectivity float64 `json:"subjectivity"`
}

type Dependency struct {
	Word   string `json:"word"`
	DepRel string `json:"dep_rel"`
	Head   string `json:"head"`
}

type Verb struct {
	Base                string `json:"base"`
	Past                string `json:"past"`
	PastParticiple      string `json:"past_participle"`
	Gerund              string `json:"gerund"`
	ThirdPersonSingular string `json:"third_person_singular"`
}

type VocabularyMainWord struct {
	Word       string                       `json:"word"`
	Base       string                       `json:"base"`
	Pos        string                       `json:"pos"`
	Definition string                       `json:"definition"`
	Translated language.TranslatedLanguages `json:"translated"`
}
