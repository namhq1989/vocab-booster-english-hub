package mapping

import (
	"github.com/namhq1989/vocab-booster-english-hub/internal/database/gen/vocab-booster/public/model"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/exercise/domain"
)

type UserExerciseStatusMapper struct{}

func (UserExerciseStatusMapper) FromModelToDomain(status model.UserExerciseStatuses) (*domain.UserExerciseStatus, error) {
	var result = &domain.UserExerciseStatus{
		ID:            status.ID,
		ExerciseID:    status.ExerciseID,
		UserID:        status.UserID,
		CorrectStreak: int(status.CorrectStreak),
		AnswerCount:   int(status.AnswerCount),
		IsFavorite:    status.IsFavorite,
		IsMastered:    status.IsMastered,
		UpdatedAt:     status.UpdatedAt,
		NextReviewAt:  status.NextReviewAt,
	}

	return result, nil
}

func (UserExerciseStatusMapper) FromDomainToModel(status domain.UserExerciseStatus) (*model.UserExerciseStatuses, error) {
	var result = &model.UserExerciseStatuses{
		ID:            status.ID,
		ExerciseID:    status.ExerciseID,
		UserID:        status.UserID,
		CorrectStreak: int32(status.CorrectStreak),
		AnswerCount:   int32(status.AnswerCount),
		IsFavorite:    status.IsFavorite,
		IsMastered:    status.IsMastered,
		UpdatedAt:     status.UpdatedAt,
		NextReviewAt:  status.NextReviewAt,
	}

	return result, nil
}
