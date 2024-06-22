package domain

import (
	"reflect"
	"time"

	"github.com/namhq1989/vocab-booster-english-hub/core/appcontext"
	apperrors "github.com/namhq1989/vocab-booster-english-hub/core/error"
	"github.com/namhq1989/vocab-booster-english-hub/core/language"
	"github.com/namhq1989/vocab-booster-english-hub/internal/database"
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
	Content      string
	Translated   language.TranslatedLanguages
	MainWord     VocabularyMainWord
	CreatedAt    time.Time
	PosTags      []PosTag
	Sentiment    Sentiment
	Dependencies []Dependency
	Verbs        []Verb
	Level        SentenceLevel
}

type VocabularyMainWord struct {
	Word       string
	Base       string
	Pos        PartOfSpeech
	Definition string
	Translated language.TranslatedLanguages
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

func (d *VocabularyExample) SetContent(content string, translated language.TranslatedLanguages) error {
	if content == "" {
		return apperrors.Vocabulary.InvalidExampleContent
	}

	// check for each translated language
	val := reflect.ValueOf(translated)
	for i := 0; i < val.NumField(); i++ {
		if val.Field(i).String() == "" {
			return apperrors.Vocabulary.InvalidExampleTranslatedLanguages
		}
	}

	d.Content = content
	d.Translated = translated
	return nil
}

func (d *VocabularyExample) SetMainWordData(word, base, definition, pos string, translated language.TranslatedLanguages) error {
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

	if definition == "" {
		return apperrors.Vocabulary.InvalidDefinition
	}

	// check for each translated language
	val := reflect.ValueOf(translated)
	for i := 0; i < val.NumField(); i++ {
		if val.Field(i).String() == "" {
			return apperrors.Vocabulary.InvalidExampleTranslatedLanguages
		}
	}

	d.MainWord.Pos = dPos
	d.MainWord.Base = base
	d.MainWord.Definition = definition
	d.MainWord.Word = word
	d.MainWord.Translated = translated
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
