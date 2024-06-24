package dto

import (
	"github.com/namhq1989/vocab-booster-english-hub/internal/genproto/vocabularypb"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/domain"
)

func ConvertVocabularyFromDomainToGrpc(vocabulary domain.Vocabulary, examples []domain.VocabularyExample) *vocabularypb.Vocabulary {
	partsOfSpeech := make([]string, 0)
	for _, pos := range vocabulary.PartsOfSpeech {
		partsOfSpeech = append(partsOfSpeech, pos.String())
	}

	exampleList := make([]*vocabularypb.VocabularyExample, 0)
	for _, example := range examples {
		posTags := make([]*vocabularypb.PosTag, 0)
		for _, posTag := range example.PosTags {
			posTags = append(posTags, &vocabularypb.PosTag{
				Word:  posTag.Word,
				Value: posTag.Value.String(),
			})
		}

		dependencies := make([]*vocabularypb.Dependency, 0)
		for _, dependency := range example.Dependencies {
			dependencies = append(dependencies, &vocabularypb.Dependency{
				Word:   dependency.Word,
				DepRel: dependency.DepRel,
				Head:   dependency.Head,
			})
		}

		verbs := make([]*vocabularypb.Verb, 0)
		for _, verb := range example.Verbs {
			verbs = append(verbs, &vocabularypb.Verb{
				Base:                verb.Base,
				Past:                verb.Past,
				PastParticiple:      verb.PastParticiple,
				Gerund:              verb.Gerund,
				ThirdPersonSingular: verb.ThirdPersonSingular,
			})
		}

		exampleList = append(exampleList, &vocabularypb.VocabularyExample{
			Id:           example.ID,
			VocabularyId: example.VocabularyID,
			Content:      example.Content,
			Translated: &vocabularypb.TranslatedLanguages{
				Vi: example.Translated.Vi,
			},
			MainWord: &vocabularypb.VocabularyMainWord{
				Word:       example.MainWord.Word,
				Base:       example.MainWord.Base,
				Pos:        example.MainWord.Pos.String(),
				Definition: example.MainWord.Definition,
				Translated: &vocabularypb.TranslatedLanguages{
					Vi: example.MainWord.Translated.Vi,
				},
			},
			PosTags: posTags,
			Sentiment: &vocabularypb.Sentiment{
				Polarity:     example.Sentiment.Polarity,
				Subjectivity: example.Sentiment.Subjectivity,
			},
			Dependencies: dependencies,
			Verbs:        verbs,
		})
	}

	result := &vocabularypb.Vocabulary{
		Id:            vocabulary.ID,
		AuthorId:      vocabulary.AuthorID,
		Term:          vocabulary.Term,
		PartsOfSpeech: partsOfSpeech,
		Ipa:           vocabulary.IPA,
		Audio:         vocabulary.Audio,
		Synonyms:      vocabulary.Synonyms,
		Antonyms:      vocabulary.Antonyms,
		Examples:      exampleList,
	}

	return result
}
