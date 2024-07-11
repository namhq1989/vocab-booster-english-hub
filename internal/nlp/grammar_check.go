package nlp

import (
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
	"github.com/namhq1989/vocab-booster-utilities/language"
)

type GrammarCheckResult struct {
	Errors []GrammarCheckError `json:"errors"`
}

type GrammarCheckError struct {
	Message     string                       `json:"message"`
	Translated  language.TranslatedLanguages `json:"translated"`
	Segment     string                       `json:"segment"`
	Replacement string                       `json:"replacement"`
}

func (n NLP) GrammarCheck(_ *appcontext.AppContext, sentence string) (result *GrammarCheckResult, err error) {
	_, err = n.httpClient.R().
		SetBody(map[string]interface{}{
			"sentence": sentence,
		}).
		SetResult(&result).
		Post("/grammar-check")
	return
}
