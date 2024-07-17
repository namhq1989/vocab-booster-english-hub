package dto

import (
	"github.com/namhq1989/vocab-booster-english-hub/internal/genproto/exercisepb"
	"github.com/namhq1989/vocab-booster-english-hub/internal/utils/staticfiles"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/exercise/domain"
)

func ConvertUserExerciseCollectionsFromDomainToGrpc(collections []domain.UserExerciseCollection, lang string) []*exercisepb.ExerciseCollection {
	var result = make([]*exercisepb.ExerciseCollection, len(collections))

	for index, collection := range collections {
		result[index] = &exercisepb.ExerciseCollection{
			Id:              collection.ID,
			Name:            collection.Name,
			Slug:            collection.Slug,
			Translated:      collection.Translated.GetLanguageValue(lang),
			StatsExercises:  int32(collection.StatsExercises),
			StatsInteracted: int32(collection.StatsInteracted),
			Image:           staticfiles.GetExerciseCollectionsEndpoint(collection.Image),
		}
	}

	return result
}
