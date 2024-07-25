package mapping

import (
	"github.com/namhq1989/vocab-booster-english-hub/internal/database/gen/vocab-booster/public/model"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/domain"
)

type VocabularyMapper struct{}

func (VocabularyMapper) FromModelToDomain(vocab model.Vocabularies) (*domain.Vocabulary, error) {
	var result = &domain.Vocabulary{
		ID:            vocab.ID,
		AuthorID:      vocab.AuthorID,
		Term:          vocab.Term,
		PartsOfSpeech: make([]domain.PartOfSpeech, 0),
		Ipa:           vocab.Ipa,
		Audio:         vocab.Audio,
		Synonyms:      vocab.Synonyms,
		Antonyms:      vocab.Antonyms,
		Frequency:     vocab.Frequency,
		CreatedAt:     vocab.CreatedAt,
		UpdatedAt:     vocab.UpdatedAt,
	}

	for _, pos := range vocab.PartsOfSpeech {
		result.PartsOfSpeech = append(result.PartsOfSpeech, domain.PartOfSpeech(pos))
	}

	return result, nil
}

func (VocabularyMapper) FromDomainToModel(vocab domain.Vocabulary) (*model.Vocabularies, error) {
	var result = &model.Vocabularies{
		ID:            vocab.ID,
		AuthorID:      vocab.AuthorID,
		Ipa:           vocab.Ipa,
		Term:          vocab.Term,
		PartsOfSpeech: make([]string, 0),
		Audio:         vocab.Audio,
		Synonyms:      vocab.Synonyms,
		Antonyms:      vocab.Antonyms,
		Frequency:     vocab.Frequency,
		CreatedAt:     vocab.CreatedAt,
		UpdatedAt:     vocab.UpdatedAt,
	}

	for _, pos := range vocab.PartsOfSpeech {
		result.PartsOfSpeech = append(result.PartsOfSpeech, pos.String())
	}

	return result, nil
}
