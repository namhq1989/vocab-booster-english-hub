package nlp

import (
	"github.com/namhq1989/vocab-booster-english-hub/core/appcontext"
	"github.com/namhq1989/vocab-booster-english-hub/core/language"
)

type SentenceAnalysisResult struct {
	Translated   language.TranslatedLanguages `json:"translated"`
	PosTags      []PosTag                     `json:"pos_tags"`
	Sentiment    Sentiment                    `json:"sentiment"`
	Dependencies []Dependency                 `json:"dependencies"`
	Verbs        []Verb                       `json:"verbs"`
}

type PosTag struct {
	Word  string `json:"word"`
	Value string `json:"value"`
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

func (n NLP) AnalyzeSentence(_ *appcontext.AppContext, sentence string) (result *SentenceAnalysisResult, err error) {
	_, err = n.httpClient.R().
		SetBody(map[string]interface{}{
			"sentence": sentence,
		}).
		SetResult(&result).
		Post("/analyze-sentence")
	return
}
