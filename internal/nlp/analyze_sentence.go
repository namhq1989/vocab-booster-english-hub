package nlp

import (
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
	"github.com/namhq1989/vocab-booster-utilities/language"
)

type SentenceAnalysisResult struct {
	Translated   language.Multilingual `json:"translated"`
	MainWord     MainWord              `json:"main_word"`
	PosTags      []PosTag              `json:"pos_tags"`
	Sentiment    Sentiment             `json:"sentiment"`
	Dependencies []Dependency          `json:"dependencies"`
	Verbs        []Verb                `json:"verbs"`
	Level        string                `json:"level"`
}

type MainWord struct {
	Word string `json:"word"`
	Base string `json:"base"`
	Pos  string `json:"pos"`
}

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
	DepRel string `json:"deprel"`
	Head   string `json:"head"`
}

type Verb struct {
	Base                string `json:"base"`
	Past                string `json:"past"`
	PastParticiple      string `json:"past_participle"`
	Gerund              string `json:"gerund"`
	ThirdPersonSingular string `json:"third_person_singular_present"`
}

func (n NLP) AnalyzeSentence(_ *appcontext.AppContext, sentence, term string) (result *SentenceAnalysisResult, err error) {
	_, err = n.httpClient.R().
		SetBody(map[string]interface{}{
			"sentence": sentence,
			"term":     term,
		}).
		SetResult(&result).
		Post("/analyze-sentence")
	return
}
