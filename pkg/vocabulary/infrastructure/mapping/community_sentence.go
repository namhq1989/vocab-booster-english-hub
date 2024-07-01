package mapping

import (
	"github.com/goccy/go-json"
	"github.com/namhq1989/vocab-booster-english-hub/core/language"
	"github.com/namhq1989/vocab-booster-english-hub/internal/database/gen/vocab-booster/public/model"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/domain"
)

type CommunitySentenceMapper struct{}

func (CommunitySentenceMapper) FromModelToDomain(sentence model.CommunitySentences) (*domain.CommunitySentence, error) {
	var result = &domain.CommunitySentence{
		ID:                 sentence.ID,
		UserID:             sentence.UserID,
		VocabularyID:       sentence.VocabularyID,
		Content:            sentence.Content,
		RequiredVocabulary: sentence.RequiredVocabulary,
		RequiredTense:      domain.ToTense(sentence.RequiredTense),
		Translated:         language.TranslatedLanguages{},
		Clauses:            make([]domain.SentenceClause, 0),
		PosTags:            make([]domain.PosTag, 0),
		Sentiment:          domain.Sentiment{},
		Dependencies:       make([]domain.Dependency, 0),
		Verbs:              make([]domain.Verb, 0),
		Level:              domain.ToSentenceLevel(sentence.Level),
		StatsLike:          int(sentence.StatsLike),
		CreatedAt:          sentence.CreatedAt,
	}

	if err := json.Unmarshal([]byte(sentence.Translated), &result.Translated); err != nil {
		return nil, err
	}

	if err := json.Unmarshal([]byte(sentence.Clauses), &result.Clauses); err != nil {
		return nil, err
	}

	if err := json.Unmarshal([]byte(sentence.Sentiment), &result.Sentiment); err != nil {
		return nil, err
	}

	if err := json.Unmarshal([]byte(sentence.PosTags), &result.PosTags); err != nil {
		return nil, err
	}

	if err := json.Unmarshal([]byte(sentence.Dependencies), &result.Dependencies); err != nil {
		return nil, err
	}

	if err := json.Unmarshal([]byte(sentence.Verbs), &result.Verbs); err != nil {
		return nil, err
	}

	return result, nil
}

func (CommunitySentenceMapper) FromDomainToModel(sentence domain.CommunitySentence) (*model.CommunitySentences, error) {
	var result = &model.CommunitySentences{
		ID:                 sentence.ID,
		UserID:             sentence.UserID,
		VocabularyID:       sentence.VocabularyID,
		Content:            sentence.Content,
		RequiredVocabulary: sentence.RequiredVocabulary,
		RequiredTense:      sentence.RequiredTense.String(),
		Translated:         "",
		Sentiment:          "",
		Clauses:            "",
		StatsLike:          int32(sentence.StatsLike),
		CreatedAt:          sentence.CreatedAt,
		PosTags:            "",
		Dependencies:       "",
		Verbs:              "",
		Level:              sentence.Level.String(),
	}

	if data, err := json.Marshal(sentence.Translated); err != nil {
		return nil, err
	} else {
		result.Translated = string(data)
	}

	clauses := make([]SentenceClause, 0)
	for _, clause := range sentence.Clauses {
		clauses = append(clauses, SentenceClause{
			Clause:         clause.Clause,
			Tense:          clause.Tense.String(),
			IsPrimaryTense: clause.IsPrimaryTense,
		})
	}
	if data, err := json.Marshal(clauses); err != nil {
		return nil, err
	} else {
		result.Clauses = string(data)
	}

	posTags := make([]PosTag, 0)
	for _, posTag := range sentence.PosTags {
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
		Polarity:     sentence.Sentiment.Polarity,
		Subjectivity: sentence.Sentiment.Subjectivity,
	}
	if data, err := json.Marshal(sentiment); err != nil {
		return nil, err
	} else {
		result.Sentiment = string(data)
	}

	dependencies := make([]Dependency, 0)
	for _, dependency := range sentence.Dependencies {
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
	for _, verb := range sentence.Verbs {
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

//
// EXTENDED
//

type ExtendedCommunitySentence struct {
	Sentence model.CommunitySentences `alias:"cs"`
	IsLiked  bool                     `alias:"csl.is_liked"`
}

type ExtendedCommunitySentenceMapper struct{}

func (ExtendedCommunitySentenceMapper) FromModelToDomain(sentence ExtendedCommunitySentence, lang string) (*domain.ExtendedCommunitySentence, error) {
	var sentenceMapper = CommunitySentenceMapper{}
	communitySentence, err := sentenceMapper.FromModelToDomain(sentence.Sentence)
	if err != nil {
		return nil, err
	}

	return &domain.ExtendedCommunitySentence{
		CommunitySentence: *communitySentence,
		Translated:        communitySentence.Translated.GetLanguageValue(lang),
		IsLiked:           sentence.IsLiked,
	}, nil
}
