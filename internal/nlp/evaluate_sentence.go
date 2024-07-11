package nlp

import (
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
	"github.com/namhq1989/vocab-booster-utilities/language"
)

type EvaluateSentenceResult struct {
	IsEnglish           bool                         `json:"is_english"`
	IsVocabularyCorrect bool                         `json:"is_vocabulary_correct"`
	IsTenseCorrect      bool                         `json:"is_tense_correct"`
	Sentiment           Sentiment                    `json:"sentiment"`
	Translated          language.TranslatedLanguages `json:"translated"`
	Clauses             []EvaluateSentenceClause     `json:"clauses"`
}

type EvaluateSentenceClause struct {
	Clause         string `json:"clause"`
	Tense          string `json:"tense"`
	IsPrimaryTense bool   `json:"is_primary_tense"`
}

func (n NLP) EvaluateSentence(_ *appcontext.AppContext, sentence, tense string, vocabulary []string) (result *EvaluateSentenceResult, err error) {
	_, err = n.httpClient.R().
		SetBody(map[string]interface{}{
			"sentence":   sentence,
			"vocabulary": vocabulary,
			"tense":      tense,
		}).
		SetResult(&result).
		Post("/evaluate-sentence")
	return
}
