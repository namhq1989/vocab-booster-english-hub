package dbmodel

import (
	"time"

	apperrors "github.com/namhq1989/vocab-booster-english-hub/core/error"
	"github.com/namhq1989/vocab-booster-english-hub/internal/database"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type VocabularyExample struct {
	ID           primitive.ObjectID `bson:"_id"`
	VocabularyID primitive.ObjectID `bson:"vocabularyId"`
	Audio        string             `bson:"audio"`
	FromLang     string             `bson:"fromLang"`
	ToLang       string             `bson:"toLang"`
	Pos          string             `bson:"pos"`
	ToDefinition string             `bson:"toDefinition"`
	Word         string             `bson:"word"`
	CreatedAt    time.Time          `bson:"createdAt"`
	PosTags      []PosTag           `bson:"posTags"`
	Sentiment    Sentiment          `bson:"sentiment"`
	Dependencies []Dependency       `bson:"dependencies"`
	Verbs        []Verb             `bson:"verbs"`
}

func (m VocabularyExample) ToDomain() domain.VocabularyExample {
	posTags := make([]domain.PosTag, 0)
	for _, posTag := range m.PosTags {
		posTags = append(posTags, domain.PosTag{
			Word:  posTag.Word,
			Value: domain.ToPartOfSpeech(posTag.Value),
		})
	}

	dependencies := make([]domain.Dependency, 0)
	for _, dependency := range m.Dependencies {
		dependencies = append(dependencies, domain.Dependency{
			Word:   dependency.Word,
			DepRel: dependency.DepRel,
			Head:   dependency.Head,
		})
	}

	verbs := make([]domain.Verb, 0)
	for _, verb := range m.Verbs {
		verbs = append(verbs, domain.Verb{
			Base:                verb.Base,
			Past:                verb.Past,
			PastParticiple:      verb.PastParticiple,
			Gerund:              verb.Gerund,
			ThirdPersonSingular: verb.ThirdPersonSingular,
		})
	}

	return domain.VocabularyExample{
		ID:           m.ID.Hex(),
		VocabularyID: m.VocabularyID.Hex(),
		Audio:        m.Audio,
		FromLang:     m.FromLang,
		ToLang:       m.ToLang,
		Pos:          domain.ToPartOfSpeech(m.Pos),
		ToDefinition: m.ToDefinition,
		Word:         m.Word,
		CreatedAt:    m.CreatedAt,
		PosTags:      posTags,
		Sentiment: domain.Sentiment{
			Polarity:     m.Sentiment.Polarity,
			Subjectivity: m.Sentiment.Subjectivity,
		},
		Dependencies: dependencies,
		Verbs:        verbs,
	}
}

func (VocabularyExample) FromDomain(example domain.VocabularyExample) (*VocabularyExample, error) {
	id, err := database.ObjectIDFromString(example.ID)
	if err != nil {
		return nil, apperrors.Common.InvalidID
	}

	vid, err := database.ObjectIDFromString(example.VocabularyID)
	if err != nil {
		return nil, apperrors.Vocabulary.InvalidVocabularyID
	}

	posTags := make([]PosTag, 0)
	for _, posTag := range example.PosTags {
		posTags = append(posTags, PosTag{
			Word:  posTag.Word,
			Value: posTag.Value.String(),
		})
	}

	dependencies := make([]Dependency, 0)
	for _, dependency := range example.Dependencies {
		dependencies = append(dependencies, Dependency{
			Word:   dependency.Word,
			DepRel: dependency.DepRel,
			Head:   dependency.Head,
		})
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

	return &VocabularyExample{
		ID:           id,
		VocabularyID: vid,
		FromLang:     example.FromLang,
		ToLang:       example.ToLang,
		Audio:        example.Audio,
		Pos:          example.Pos.String(),
		ToDefinition: example.ToDefinition,
		Word:         example.Word,
		CreatedAt:    example.CreatedAt,
		PosTags:      posTags,
		Sentiment: Sentiment{
			Polarity:     example.Sentiment.Polarity,
			Subjectivity: example.Sentiment.Subjectivity,
		},
		Dependencies: dependencies,
		Verbs:        verbs,
	}, nil
}
