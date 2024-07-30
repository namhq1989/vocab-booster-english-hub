package dto

import (
	"github.com/namhq1989/vocab-booster-english-hub/internal/genproto/vocabularypb"
	"github.com/namhq1989/vocab-booster-english-hub/internal/utils/staticfiles"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/domain"
)

func ConvertVocabularyFromDomainToGrpc(vocabulary domain.Vocabulary, examples []domain.VocabularyExample) *vocabularypb.Vocabulary {
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

	exampleBriefs := make([]*vocabularypb.VocabularyExampleBrief, 0)
	for _, example := range examples {
		exampleBriefs = append(exampleBriefs, &vocabularypb.VocabularyExampleBrief{
			Id:      example.ID,
			Content: ConvertMultilingualToGrpcData(example.Content),
			Audio:   staticfiles.GetExampleEndpoint(example.Audio),
			MainWord: &vocabularypb.VocabularyMainWord{
				Word: example.MainWord.Word,
				Base: example.MainWord.Base,
				Pos:  example.MainWord.Pos.String(),
			},
		})
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
		Examples:      exampleBriefs,
	}

	return result
}

func ConvertVocabularyBriefFromDomainToGrpc(vocabulary domain.Vocabulary) *vocabularypb.VocabularyBrief {
	partsOfSpeech := make([]string, 0)
	for _, pos := range vocabulary.PartsOfSpeech {
		partsOfSpeech = append(partsOfSpeech, pos.String())
	}

	result := &vocabularypb.VocabularyBrief{
		Id:            vocabulary.ID,
		Term:          vocabulary.Term,
		PartsOfSpeech: partsOfSpeech,
		Ipa:           vocabulary.Ipa,
		Audio:         staticfiles.GetVocabularyEndpoint(vocabulary.Audio),
	}

	return result
}
