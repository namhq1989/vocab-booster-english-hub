package mapping

import (
	"github.com/goccy/go-json"
	"github.com/namhq1989/vocab-booster-english-hub/internal/database/gen/vocab-booster/public/model"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/domain"
	"github.com/namhq1989/vocab-booster-utilities/language"
)

type VocabularyMapper struct{}

type VocabularyDefinition struct {
	Pos        string                `json:"pos"`
	Definition language.Multilingual `json:"definition"`
}

func (VocabularyMapper) FromModelToDomain(vocab model.Vocabularies) (*domain.Vocabulary, error) {
	var result = &domain.Vocabulary{
		ID:            vocab.ID,
		AuthorID:      vocab.AuthorID,
		Term:          vocab.Term,
		Definitions:   make([]domain.VocabularyDefinition, 0),
		PartsOfSpeech: make([]domain.PartOfSpeech, 0),
		Ipa:           vocab.Ipa,
		Audio:         vocab.Audio,
		Synonyms:      vocab.Synonyms,
		Antonyms:      vocab.Antonyms,
		Frequency:     vocab.Frequency,
		CreatedAt:     vocab.CreatedAt,
		UpdatedAt:     vocab.UpdatedAt,
	}

	if len(result.Synonyms) == 1 && result.Synonyms[0] == "" {
		result.Synonyms = make([]string, 0)
	}

	if len(result.Antonyms) == 1 && result.Antonyms[0] == "" {
		result.Antonyms = make([]string, 0)
	}

	if err := json.Unmarshal([]byte(vocab.Definitions), &result.Definitions); err != nil {
		return nil, err
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
		Definitions:   "",
		PartsOfSpeech: make([]string, 0),
		Audio:         vocab.Audio,
		Synonyms:      vocab.Synonyms,
		Antonyms:      vocab.Antonyms,
		Frequency:     vocab.Frequency,
		CreatedAt:     vocab.CreatedAt,
		UpdatedAt:     vocab.UpdatedAt,
	}

	definitions := make([]VocabularyDefinition, 0)
	for _, def := range vocab.Definitions {
		definitions = append(definitions, VocabularyDefinition{
			Pos:        def.Pos.String(),
			Definition: def.Definition,
		})
	}
	if data, err := json.Marshal(definitions); err != nil {
		return nil, err
	} else {
		result.Definitions = string(data)
	}

	for _, pos := range vocab.PartsOfSpeech {
		result.PartsOfSpeech = append(result.PartsOfSpeech, pos.String())
	}

	return result, nil
}
