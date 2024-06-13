package dto

import (
	"github.com/namhq1989/vocab-booster-english-hub/internal/genproto/vocabularypb"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/domain"
)

func ConvertVocabularyFromDomainToGrpc(vocabulary domain.Vocabulary, examples []domain.VocabularyExample) *vocabularypb.Vocabulary {
	partsOfSpeech := make([]string, 0)
	for _, pos := range vocabulary.Data.PartsOfSpeech {
		partsOfSpeech = append(partsOfSpeech, pos.String())
	}

	exampleList := make([]*vocabularypb.VocabularyExample, 0)
	for _, example := range examples {
		exampleList = append(exampleList, &vocabularypb.VocabularyExample{
			Id:         example.ID,
			English:    example.English,
			Vietnamese: example.Vietnamese,
			Pos:        example.POS.String(),
			Definition: example.Definition,
			Word:       example.Word,
		})
	}

	result := &vocabularypb.Vocabulary{
		Id:       vocabulary.ID,
		AuthorId: vocabulary.AuthorID,
		Term:     vocabulary.Term,
		Data: &vocabularypb.VocabularyData{
			PartsOfSpeech: partsOfSpeech,
			Ipa:           vocabulary.Data.IPA,
			AudioName:     vocabulary.Data.AudioName,
			Synonyms:      vocabulary.Data.Synonyms,
			Antonyms:      vocabulary.Data.Antonyms,
		},
		Examples: exampleList,
	}

	return result
}
