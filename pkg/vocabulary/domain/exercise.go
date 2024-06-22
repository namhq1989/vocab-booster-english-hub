package domain

import (
	"github.com/namhq1989/vocab-booster-english-hub/core/appcontext"
	"github.com/namhq1989/vocab-booster-english-hub/core/language"
)

type ExerciseHub interface {
	CreateExercise(ctx *appcontext.AppContext, vocabularyExampleID, level, content, vocabulary, correctAnswer string, translated language.TranslatedLanguages, options []string) error
	UpdateExerciseAudio(ctx *appcontext.AppContext, vocabularyExampleID, audio string) error
}
