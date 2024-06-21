package infrastructure

import (
	"github.com/namhq1989/vocab-booster-english-hub/core/appcontext"
	"github.com/namhq1989/vocab-booster-english-hub/core/language"
	"github.com/namhq1989/vocab-booster-english-hub/internal/genproto/exercisepb"
)

type ExerciseHub struct {
	client exercisepb.ExerciseServiceClient
}

func NewExerciseHub(client exercisepb.ExerciseServiceClient) ExerciseHub {
	return ExerciseHub{
		client: client,
	}
}

func (r ExerciseHub) CreateExercise(ctx *appcontext.AppContext, vocabularyExampleID, content, vocabulary, correctAnswer string, translated language.TranslatedLanguages, options []string) error {
	_, err := r.client.NewExercise(ctx.Context(), &exercisepb.NewExerciseRequest{
		VocabularyExampleId: vocabularyExampleID,
		Content:             content,
		Translated: &exercisepb.TranslatedLanguages{
			Vi: translated.Vi,
		},
		Vocabulary:    vocabulary,
		CorrectAnswer: correctAnswer,
		Options:       options,
	})

	return err
}
