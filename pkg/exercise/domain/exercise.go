package domain

import (
	"slices"
	"time"

	"github.com/namhq1989/vocab-booster-english-hub/internal/database"
	apperrors "github.com/namhq1989/vocab-booster-english-hub/internal/utils/error"
	"github.com/namhq1989/vocab-booster-english-hub/internal/utils/manipulation"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
	"github.com/namhq1989/vocab-booster-utilities/language"
)

type ExerciseRepository interface {
	FindExerciseByID(ctx *appcontext.AppContext, exerciseID string) (*Exercise, error)
	FindExerciseByVocabularyExampleID(ctx *appcontext.AppContext, exampleID string) (*Exercise, error)
	CreateExercise(ctx *appcontext.AppContext, exercise Exercise) error
	UpdateExercise(ctx *appcontext.AppContext, exercise Exercise) error
	PickRandomExercisesForUser(ctx *appcontext.AppContext, filter UserExerciseFilter) ([]UserExercise, error)
	CountExercisesByCriteria(ctx *appcontext.AppContext, criteria string, ts time.Time) (int64, error)
}

type Exercise struct {
	ID                  string
	VocabularyExampleID string
	Level               ExerciseLevel
	Audio               string
	Vocabulary          string
	Content             language.Multilingual
	CorrectAnswer       string
	Options             []string
	CreatedAt           time.Time
}

func NewExercise(vocabularyExampleID, level string, content language.Multilingual, vocabulary, correctAnswer string, options []string) (*Exercise, error) {
	if vocabularyExampleID == "" {
		return nil, apperrors.Exercise.InvalidExerciseID
	}

	dLevel := ToExerciseLevel(level)
	if !dLevel.IsValid() {
		return nil, apperrors.Exercise.InvalidLevel
	}

	if content.IsEmpty() {
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
		Level:               dLevel,
		Audio:               "",
		Content:             content,
		Vocabulary:          vocabulary,
		CorrectAnswer:       correctAnswer,
		Options:             options,
		CreatedAt:           manipulation.NowUTC(),
	}, nil
}

func (e *Exercise) SetAudio(audio string) error {
	if audio == "" {
		return apperrors.Exercise.InvalidAudio
	}

	e.Audio = audio
	return nil
}
