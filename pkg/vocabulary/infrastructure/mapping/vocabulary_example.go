package mapping

import (
	"github.com/goccy/go-json"
	"github.com/namhq1989/vocab-booster-english-hub/internal/database/gen/vocab-booster/public/model"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/domain"
	"github.com/namhq1989/vocab-booster-utilities/language"
)

type VocabularyExampleMapper struct{}

func (VocabularyExampleMapper) FromModelToDomain(example model.VocabularyExamples) (*domain.VocabularyExample, error) {
	var result = &domain.VocabularyExample{
		ID:           example.ID,
		VocabularyID: example.VocabularyID,
		Audio:        example.Audio,
		Content:      language.Multilingual{},
		MainWord:     domain.VocabularyMainWord{},
		PosTags:      make([]domain.PosTag, 0),
		Sentiment:    domain.Sentiment{},
		Dependencies: make([]domain.Dependency, 0),
		Verbs:        make([]domain.Verb, 0),
		Level:        domain.ToSentenceLevel(example.Level),
		CreatedAt:    example.CreatedAt,
	}

	if err := json.Unmarshal([]byte(example.Content), &result.Content); err != nil {
		return nil, err
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
		Content:      "",
		MainWord:     "",
		PosTags:      "",
		Sentiment:    "",
		Dependencies: "",
		Verbs:        "",
		Level:        example.Level.String(),
		CreatedAt:    example.CreatedAt,
	}

	if data, err := json.Marshal(example.Content); err != nil {
		return nil, err
	} else {
		result.Content = string(data)
	}

	mainWord := VocabularyMainWord{
		Word: example.MainWord.Word,
		Base: example.MainWord.Base,
	}
	if data, err := json.Marshal(mainWord); err != nil {
		return nil, err
	} else {
		result.MainWord = string(data)
	}

	posTags := make([]PosTag, 0)
	for _, posTag := range example.PosTags {
		posTags = append(posTags, PosTag{
			Word:  posTag.Word,
			Value: posTag.Value.String(),
			Level: posTag.Level,
		})
	}
	if data, err := json.Marshal(posTags); err != nil {
		return nil, err
	} else {
		result.PosTags = string(data)
	}

	sentiment := Sentiment{
		Polarity:     example.Sentiment.Polarity,
		Subjectivity: example.Sentiment.Subjectivity,
	}
	if data, err := json.Marshal(sentiment); err != nil {
		return nil, err
	} else {
		result.Sentiment = string(data)
	}

	dependencies := make([]Dependency, 0)
	for _, dependency := range example.Dependencies {
		dependencies = append(dependencies, Dependency{
			Word:   dependency.Word,
			DepRel: dependency.DepRel,
			Head:   dependency.Head,
		})
	}
	if data, err := json.Marshal(dependencies); err != nil {
		return nil, err
	} else {
		result.Dependencies = string(data)
	}

	verbs := make([]Verb, 0)
	for _, verb := range example.Verbs {
		verbs = append(verbs, Verb{
			Base:                verb.Base,
			Past:                verb.Past,
			PastParticiple:      verb.PastParticiple,
			Gerund:              verb.Gerund,
			ThirdPersonSingular: verb.ThirdPersonSingular,
		})
	}
	if data, err := json.Marshal(verbs); err != nil {
		return nil, err
	} else {
		result.Verbs = string(data)
	}

	return result, nil

}
