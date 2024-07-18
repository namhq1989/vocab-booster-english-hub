package mapping

import (
	"github.com/namhq1989/vocab-booster-english-hub/pkg/exercise/domain"
)

type UserAggregatedExercise struct {
	Date     string
	Exercise int64
}

type UserAggregatedExerciseMapper struct{}

func (UserAggregatedExerciseMapper) FromModelToDomain(uae UserAggregatedExercise) (*domain.UserAggregatedExercise, error) {
	return &domain.UserAggregatedExercise{
		Date:     uae.Date,
		Exercise: uae.Exercise,
	}, nil
}
