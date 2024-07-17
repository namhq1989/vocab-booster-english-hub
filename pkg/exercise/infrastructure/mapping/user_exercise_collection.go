package mapping

import (
	"github.com/goccy/go-json"
	"github.com/namhq1989/vocab-booster-english-hub/internal/database/gen/vocab-booster/public/model"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/exercise/domain"
	"github.com/namhq1989/vocab-booster-utilities/language"
)

type UserExerciseCollection struct {
	Collection model.ExerciseCollections          `alias:"ec"`
	Status     model.UserExerciseCollectionStatus `alias:"uecs"`
}

type UserExerciseCollectionMapper struct{}

func (UserExerciseCollectionMapper) FromModelToDomain(uec UserExerciseCollection) (*domain.UserExerciseCollection, error) {
	var result = &domain.UserExerciseCollection{
		ID:              uec.Collection.ID,
		Name:            uec.Collection.Name,
		Slug:            uec.Collection.Slug,
		Translated:      language.TranslatedLanguages{},
		Image:           uec.Collection.Image,
		Order:           int(uec.Collection.Order),
		StatsExercises:  int(uec.Collection.StatsExercises),
		StatsInteracted: int(uec.Status.InteractedExercises),
	}

	if err := json.Unmarshal([]byte(uec.Collection.Translated), &result.Translated); err != nil {
		return nil, err
	}

	return result, nil
}
