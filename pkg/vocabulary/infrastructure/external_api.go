package infrastructure

import (
	"slices"

	"github.com/namhq1989/vocab-booster-english-hub/internal/externalapi"
	"github.com/namhq1989/vocab-booster-english-hub/internal/nlp"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/domain"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
	"github.com/namhq1989/vocab-booster-utilities/language"
)

type ExternalAPIRepository struct {
	ea  externalapi.Operations
	nlp nlp.Operations
}

func NewExternalAPIRepository(ea externalapi.Operations, nlp nlp.Operations) ExternalAPIRepository {
	return ExternalAPIRepository{
		ea:  ea,
		nlp: nlp,
	}
}

func (r ExternalAPIRepository) SearchTermWithDatamuse(ctx *appcontext.AppContext, term string) (*domain.DatamuseSearchTermResult, error) {
	apiResult, err := r.ea.SearchTermWitDatamuse(ctx, term)
	if err != nil {
		return nil, err
	}

	var result = domain.DatamuseSearchTermResult{
		Synonyms:      apiResult.Synonyms,
		Antonyms:      apiResult.Antonyms,
		Frequency:     apiResult.Frequency,
		Ipa:           apiResult.Ipa,
		PartsOfSpeech: make([]string, 0),
		Definitions:   make([]domain.VocabularyDefinition, 0),
	}

	for _, def := range apiResult.Definitions {
		multilingual, transErr := r.translateDefinition(ctx, def.Definition)
		if transErr != nil {
			ctx.Logger().Error("failed to translate definition", transErr, appcontext.Fields{"definition": def.Definition})
			continue
		}

		pos := domain.MappingPos(def.Pos)
		if !slices.Contains(result.PartsOfSpeech, pos) {
			result.PartsOfSpeech = append(result.PartsOfSpeech, pos)
		}

		result.Definitions = append(result.Definitions, domain.VocabularyDefinition{
			Pos:        domain.ToPartOfSpeech(pos),
			Definition: *multilingual,
		})
	}

	return &result, nil
}

func (r ExternalAPIRepository) translateDefinition(ctx *appcontext.AppContext, definition string) (*language.Multilingual, error) {
	return r.nlp.Translate(ctx, definition)
}
