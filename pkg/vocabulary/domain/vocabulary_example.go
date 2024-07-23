package domain

import (
	"time"

	"github.com/namhq1989/vocab-booster-english-hub/internal/database"
	apperrors "github.com/namhq1989/vocab-booster-english-hub/internal/utils/error"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
	"github.com/namhq1989/vocab-booster-utilities/language"
)

type VocabularyExampleRepository interface {
	FindVocabularyExamplesByVocabularyID(ctx *appcontext.AppContext, vocabularyID string) ([]VocabularyExample, error)
	CreateVocabularyExamples(ctx *appcontext.AppContext, examples []VocabularyExample) error
	UpdateVocabularyExample(ctx *appcontext.AppContext, example VocabularyExample) error
}

type VocabularyExample struct {
	ID           string
	VocabularyID string
	Audio        string
	Content      language.Multilingual
	MainWord     VocabularyMainWord
	PosTags      []PosTag
	Sentiment    Sentiment
	Dependencies []Dependency
	Verbs        []Verb
	Level        SentenceLevel
	CreatedAt    time.Time
}

type VocabularyMainWord struct {
	Word       string
	Base       string
	Pos        PartOfSpeech
	Definition language.Multilingual
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

func (d *VocabularyExample) SetAudio(audio string) error {
	if audio == "" {
		return apperrors.Vocabulary.InvalidAudioName
	}
	d.Audio = audio
	return nil
}

func (d *VocabularyExample) SetContent(content language.Multilingual) error {
	if content.IsEmpty() {
		return apperrors.Vocabulary.InvalidExampleContent
	}

	d.Content = content
	return nil
}

func (d *VocabularyExample) SetMainWordData(word, base, pos string, definition language.Multilingual) error {
	dPos := ToPartOfSpeech(pos)
	if !dPos.IsValid() {
		return apperrors.Vocabulary.InvalidPartOfSpeech
	}

	if base == "" {
		return apperrors.Vocabulary.InvalidTerm
	}

	if word == "" {
		return apperrors.Vocabulary.InvalidTerm
	}

	if definition.IsEmpty() {
		return apperrors.Vocabulary.InvalidDefinition
	}

	d.MainWord.Pos = dPos
	d.MainWord.Base = base
	d.MainWord.Word = word
	d.MainWord.Definition = definition
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

func (d *VocabularyExample) SetLevel(level string) error {
	dLevel := ToSentenceLevel(level)
	if !dLevel.IsValid() {
		return apperrors.Vocabulary.InvalidExampleLevel
	}

	d.Level = dLevel
	return nil
}
