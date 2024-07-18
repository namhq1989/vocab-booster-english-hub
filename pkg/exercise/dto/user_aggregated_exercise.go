package dto

import (
	"github.com/namhq1989/vocab-booster-english-hub/internal/genproto/exercisepb"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/exercise/domain"
)

func ConvertUserAggregatedExercisesFromDomainToGrpc(aggregatedExercises []domain.UserAggregatedExercise) []*exercisepb.UserAggregatedExercise {
	var result = make([]*exercisepb.UserAggregatedExercise, len(aggregatedExercises))

	for index, uae := range aggregatedExercises {
		result[index] = &exercisepb.UserAggregatedExercise{
			Date:     uae.Date,
			Exercise: uae.Exercise,
		}
	}

	return result
}
