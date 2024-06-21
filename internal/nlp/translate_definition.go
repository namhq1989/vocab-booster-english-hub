package nlp

import (
	"github.com/namhq1989/vocab-booster-english-hub/core/appcontext"
)

type DefinitionTranslationResult struct {
	Vi string `json:"vi"`
}

func (n NLP) TranslateDefinition(_ *appcontext.AppContext, definition string) (result *DefinitionTranslationResult, err error) {
	_, err = n.httpClient.R().
		SetBody(map[string]interface{}{
			"definition": definition,
		}).
		SetResult(&result).
		Post("/translate-definition")
	return
}
