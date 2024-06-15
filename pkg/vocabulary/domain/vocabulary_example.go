package domain

import (
	"time"

	"github.com/namhq1989/vocab-booster-english-hub/core/appcontext"
	apperrors "github.com/namhq1989/vocab-booster-english-hub/core/error"
	"github.com/namhq1989/vocab-booster-english-hub/internal/database"
)

type VocabularyExampleRepository interface {
	FindVocabularyExamplesByVocabularyID(ctx *appcontext.AppContext, vocabularyID string) ([]VocabularyExample, error)
	CreateVocabularyExamples(ctx *appcontext.AppContext, examples []VocabularyExample) error
}

type VocabularyExample struct {
	ID           string
	VocabularyID string
	FromLang     string
	ToLang       string
	Pos          PartOfSpeech
	ToDefinition string
	Word         string
	CreatedAt    time.Time
	PosTags      []PosTag
	Sentiment    Sentiment
	Dependencies []Dependency
	Verbs        []Verb
}

func NewVocabularyExample(vocabularyID string) (*VocabularyExample, error) {
	if vocabularyID == "" {
		return nil, apperrors.Vocabulary.InvalidVocabularyID
	}

	return &VocabularyExample{
		ID:           database.NewStringID(),
		VocabularyID: vocabularyID,
		CreatedAt:    time.Now(),
	}, nil
}

func (d *VocabularyExample) SetContent(fromLang, toLang string) error {
	if fromLang == "" || toLang == "" {
		return apperrors.Vocabulary.InvalidExampleLanguage
	}

	d.FromLang = fromLang
	d.ToLang = toLang
	return nil
}

func (d *VocabularyExample) SetWordData(word, toDefinition, pos string) error {

	dPos := ToPartOfSpeech(pos)
	if !dPos.IsValid() {
		return apperrors.Vocabulary.InvalidPartOfSpeech
	}

	if word == "" {
		return apperrors.Vocabulary.InvalidTerm
	}

	if toDefinition == "" {
		return apperrors.Vocabulary.InvalidDefinition
	}

	d.Pos = dPos
	d.ToDefinition = toDefinition
	d.Word = word
	return nil
}

func (d *VocabularyExample) SetPosTags(posTags []PosTag) error {
	if len(posTags) == 0 {
		return apperrors.Vocabulary.InvalidExamplePosTags
	}

	d.PosTags = posTags
	return nil
}

func (d *VocabularyExample) SetSentiment(polarity, subjectivity float64) error {
	d.Sentiment.Polarity = polarity
	d.Sentiment.Subjectivity = subjectivity
	return nil
}

func (d *VocabularyExample) SetDependencies(deps []Dependency) error {
	d.Dependencies = deps
	return nil
}

func (d *VocabularyExample) SetVerbs(verbs []Verb) error {
	d.Verbs = verbs
	return nil
}
