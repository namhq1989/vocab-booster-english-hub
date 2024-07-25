package dto

import (
	"github.com/namhq1989/vocab-booster-english-hub/internal/genproto/vocabularypb"
	"github.com/namhq1989/vocab-booster-english-hub/internal/utils/staticfiles"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/domain"
)

func ConvertVocabularyFromDomainToGrpc(vocabulary domain.Vocabulary) *vocabularypb.Vocabulary {
	definitions := make([]*vocabularypb.VocabularyDefinition, 0)
	for _, def := range vocabulary.Definitions {
		definitions = append(definitions, &vocabularypb.VocabularyDefinition{
			Pos:        def.Pos.String(),
			Definition: ConvertMultilingualToGrpcData(def.Definition),
		})
	}

	partsOfSpeech := make([]string, 0)
	for _, pos := range vocabulary.PartsOfSpeech {
		partsOfSpeech = append(partsOfSpeech, pos.String())
	}

	result := &vocabularypb.Vocabulary{
		Id:            vocabulary.ID,
		AuthorId:      vocabulary.AuthorID,
		Term:          vocabulary.Term,
		Definitions:   definitions,
		PartsOfSpeech: partsOfSpeech,
		Ipa:           vocabulary.Ipa,
		Audio:         staticfiles.GetVocabularyEndpoint(vocabulary.Audio),
		Synonyms:      vocabulary.Synonyms,
		Antonyms:      vocabulary.Antonyms,
	}

	return result
}
