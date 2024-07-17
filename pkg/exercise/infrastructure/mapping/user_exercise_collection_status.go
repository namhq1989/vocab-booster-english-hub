package mapping

import (
	"github.com/namhq1989/vocab-booster-english-hub/internal/database/gen/vocab-booster/public/model"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/exercise/domain"
)

type UserExerciseCollectionStatusMapper struct{}

func (UserExerciseCollectionStatusMapper) FromModelToDomain(status model.UserExerciseCollectionStatus) (*domain.UserExerciseCollectionStatus, error) {
	var result = &domain.UserExerciseCollectionStatus{
		ID:                  status.ID,
		UserID:              status.UserID,
		CollectionID:        status.CollectionID,
		InteractedExercises: int(status.InteractedExercises),
	}

	return result, nil
}

func (UserExerciseCollectionStatusMapper) FromDomainToModel(status domain.UserExerciseCollectionStatus) (*model.UserExerciseCollectionStatus, error) {
	var result = &model.UserExerciseCollectionStatus{
		ID:                  status.ID,
		UserID:              status.UserID,
		CollectionID:        status.CollectionID,
		InteractedExercises: int32(status.InteractedExercises),
	}

	return result, nil
}
