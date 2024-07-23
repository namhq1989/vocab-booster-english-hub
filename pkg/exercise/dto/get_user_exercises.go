package dto

import (
	"github.com/namhq1989/vocab-booster-english-hub/internal/genproto/exercisepb"
	"github.com/namhq1989/vocab-booster-english-hub/internal/utils/staticfiles"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/exercise/domain"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ConvertUserExercisesFromDomainToGrpc(exercises []domain.UserExercise, lang string) []*exercisepb.UserExercise {
	var result = make([]*exercisepb.UserExercise, len(exercises))

	for index, exercise := range exercises {
		content := exercise.Content.GetLocalized(lang)

		result[index] = &exercisepb.UserExercise{
			Id:            exercise.ID,
			Level:         exercise.Level.String(),
			Audio:         staticfiles.GetExampleEndpoint(exercise.Audio),
			Vocabulary:    exercise.Vocabulary,
			Content:       ConvertMultilingualToGrpcData(content),
			CorrectAnswer: exercise.CorrectAnswer,
			Options:       exercise.Options,
			CorrectStreak: int32(exercise.CorrectStreak),
			IsFavorite:    exercise.IsFavorite,
			IsMastered:    exercise.IsMastered,
			UpdatedAt:     timestamppb.New(exercise.UpdatedAt),
			NextReviewAt:  timestamppb.New(exercise.NextReviewAt),
		}
	}

	return result
}
