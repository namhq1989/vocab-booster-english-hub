package mapping

import (
	"github.com/goccy/go-json"
	"github.com/namhq1989/vocab-booster-english-hub/core/language"
	"github.com/namhq1989/vocab-booster-english-hub/internal/database/gen/vocab-booster/public/model"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/domain"
)

type VocabularyExampleMapper struct{}

func (VocabularyExampleMapper) FromModelToDomain(example model.VocabularyExamples) (*domain.VocabularyExample, error) {
	var result = &domain.VocabularyExample{
		ID:           example.ID,
		VocabularyID: example.VocabularyID,
		Audio:        example.Audio,
		Content:      example.Content,
		Translated:   language.TranslatedLanguages{},
		MainWord:     domain.VocabularyMainWord{},
		PosTags:      make([]domain.PosTag, 0),
		Sentiment:    domain.Sentiment{},
		Dependencies: make([]domain.Dependency, 0),
		Verbs:        make([]domain.Verb, 0),
		Level:        domain.ToSentenceLevel(example.Level),
		CreatedAt:    example.CreatedAt,
	}

	if err := json.Unmarshal([]byte(example.PosTags), &result.PosTags); err != nil {
		return nil, err
	}

	if err := json.Unmarshal([]byte(example.Dependencies), &result.Dependencies); err != nil {
		return nil, err
	}

	if err := json.Unmarshal([]byte(example.Verbs), &result.Verbs); err != nil {
		return nil, err
	}

	if err := json.Unmarshal([]byte(example.Level), &result.Level); err != nil {
		return nil, err
	}

	if err := json.Unmarshal([]byte(example.Translated), &result.Translated); err != nil {
		return nil, err
	}

	if err := json.Unmarshal([]byte(example.MainWord), &result.MainWord); err != nil {
		return nil, err
	}

	if err := json.Unmarshal([]byte(example.Sentiment), &result.Sentiment); err != nil {
		return nil, err
	}

	return result, nil
}

func (VocabularyExampleMapper) FromDomainToModel(example domain.VocabularyExample) (*model.VocabularyExamples, error) {
	var result = &model.VocabularyExamples{
		ID:           example.ID,
		VocabularyID: example.VocabularyID,
		Audio:        example.Audio,
		Content:      example.Content,
		Translated:   "",
		MainWord:     "",
		PosTags:      "",
		Sentiment:    "",
		Dependencies: "",
		Verbs:        "",
		Level:        example.Level.String(),
		CreatedAt:    example.CreatedAt,
	}

	if data, err := json.Marshal(example.Translated); err != nil {
		return nil, err
	} else {
		result.Translated = string(data)
	}

	if data, err := json.Marshal(example.MainWord); err != nil {
		return nil, err
	} else {
		result.MainWord = string(data)
	}

	if data, err := json.Marshal(example.PosTags); err != nil {
		return nil, err
	} else {
		result.PosTags = string(data)
	}

	if data, err := json.Marshal(example.Sentiment); err != nil {
		return nil, err
	} else {
		result.Sentiment = string(data)
	}

	if data, err := json.Marshal(example.Dependencies); err != nil {
		return nil, err
	} else {
		result.Dependencies = string(data)
	}

	if data, err := json.Marshal(example.Verbs); err != nil {
		return nil, err
	} else {
		result.Verbs = string(data)
	}

	return result, nil

}
