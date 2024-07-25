package infrastructure

import (
	"github.com/namhq1989/vocab-booster-english-hub/internal/genproto/exercisepb"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
	"github.com/namhq1989/vocab-booster-utilities/language"
)

type ExerciseHub struct {
	client exercisepb.ExerciseServiceClient
}

func NewExerciseHub(client exercisepb.ExerciseServiceClient) ExerciseHub {
	return ExerciseHub{
		client: client,
	}
}

func (r ExerciseHub) CreateExercise(ctx *appcontext.AppContext, vocabularyExampleID, level string, frequency float64, content language.Multilingual, vocabulary, correctAnswer string, options []string) error {
	_, err := r.client.NewExercise(ctx.Context(), &exercisepb.NewExerciseRequest{
		VocabularyExampleId: vocabularyExampleID,
		Level:               level,
		Frequency:           frequency,
		Content: &exercisepb.Multilingual{
			English:    content.English,
			Vietnamese: content.Vietnamese,
		},
		Vocabulary:    vocabulary,
		CorrectAnswer: correctAnswer,
		Options:       options,
	})

	return err
}

func (r ExerciseHub) UpdateExerciseAudio(ctx *appcontext.AppContext, vocabularyExampleID, audio string) error {
	_, err := r.client.UpdateExerciseAudio(ctx.Context(), &exercisepb.UpdateExerciseAudioRequest{
		VocabularyExampleId: vocabularyExampleID,
		Audio:               audio,
	})

	return err
}
