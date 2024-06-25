package dto

import (
	"github.com/namhq1989/vocab-booster-english-hub/internal/genproto/exercisepb"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/exercise/domain"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ConvertUserExerciseFromDomainToGrpc(exercises []domain.UserExercise) []*exercisepb.UserExercise {
	var result = make([]*exercisepb.UserExercise, len(exercises))

	for index, exercise := range exercises {
		result[index] = &exercisepb.UserExercise{
			Id:            exercise.ID,
			Level:         exercise.Level.String(),
			Audio:         exercise.Audio,
			Vocabulary:    exercise.Vocabulary,
			Content:       exercise.Content,
			Translated:    exercise.Translated,
			CorrectAnswer: exercise.CorrectAnswer,
			Options:       exercise.Options,
			CorrectStreak: int32(exercise.CorrectStreak),
			IsFavorite:    exercise.IsFavorite,
			IsMastered:    exercise.IsMastered,
			NextReviewAt:  timestamppb.New(exercise.NextReviewAt),
		}
	}

	return result
}
