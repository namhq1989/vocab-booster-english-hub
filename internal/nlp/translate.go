package nlp

import (
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
	"github.com/namhq1989/vocab-booster-utilities/language"
)

func (n NLP) Translate(_ *appcontext.AppContext, term string) (result *language.TranslatedLanguages, err error) {
	_, err = n.httpClient.R().
		SetBody(map[string]interface{}{
			"term": term,
		}).
		SetResult(&result).
		Post("/translate")
	return
}
