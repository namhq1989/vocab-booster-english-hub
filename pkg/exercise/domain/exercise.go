package domain

import (
	"slices"
	"time"

	"github.com/namhq1989/vocab-booster-english-hub/core/language"

	"github.com/namhq1989/vocab-booster-english-hub/core/appcontext"
	apperrors "github.com/namhq1989/vocab-booster-english-hub/core/error"
	"github.com/namhq1989/vocab-booster-english-hub/internal/database"
)

type ExerciseRepository interface {
	CreateExercise(ctx *appcontext.AppContext, exercise Exercise) error
}

type Exercise struct {
	ID                  string
	VocabularyExampleID string
	Vocabulary          string
	Content             string
	Translated          language.TranslatedLanguages
	CorrectAnswer       string
	Options             []string
	CreatedAt           time.Time
}

func NewExercise(vocabularyExampleID, content, vocabulary, correctAnswer string, translated language.TranslatedLanguages, options []string) (*Exercise, error) {
	if vocabularyExampleID == "" {
		return nil, apperrors.Exercise.InvalidExerciseID
	}

	if content == "" {
		return nil, apperrors.Exercise.InvalidContent
	}

	if vocabulary == "" {
		return nil, apperrors.Exercise.InvalidVocabulary
	}

	if correctAnswer == "" {
		return nil, apperrors.Exercise.InvalidCorrectAnswer
	}

	if len(options) == 0 {
		return nil, apperrors.Exercise.InvalidOptions
	}

	if !slices.Contains(options, correctAnswer) {
		return nil, apperrors.Exercise.InvalidOptions
	}

	return &Exercise{
		ID:                  database.NewStringID(),
		VocabularyExampleID: vocabularyExampleID,
		Content:             content,
		Translated:          translated,
		Vocabulary:          vocabulary,
		CorrectAnswer:       correctAnswer,
		Options:             options,
		CreatedAt:           time.Now(),
	}, nil
}
