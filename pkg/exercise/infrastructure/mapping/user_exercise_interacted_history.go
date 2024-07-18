package mapping

import (
	"github.com/namhq1989/vocab-booster-english-hub/internal/database/gen/vocab-booster/public/model"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/exercise/domain"
)

type UserExerciseInteractedHistoryMapper struct{}

func (UserExerciseInteractedHistoryMapper) FromModelToDomain(history model.UserExerciseInteractedHistories) (*domain.UserExerciseInteractedHistory, error) {
	var result = &domain.UserExerciseInteractedHistory{
		ID:         history.ID,
		ExerciseID: history.ExerciseID,
		UserID:     history.UserID,
		Date:       history.Date,
	}

	return result, nil
}

func (UserExerciseInteractedHistoryMapper) FromDomainToModel(history domain.UserExerciseInteractedHistory) (*model.UserExerciseInteractedHistories, error) {
	var result = &model.UserExerciseInteractedHistories{
		ID:         history.ID,
		ExerciseID: history.ExerciseID,
		UserID:     history.UserID,
		Date:       history.Date,
	}

	return result, nil
}
