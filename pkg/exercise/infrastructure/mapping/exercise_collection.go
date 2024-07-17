package mapping

import (
	"github.com/goccy/go-json"
	"github.com/namhq1989/vocab-booster-english-hub/internal/database/gen/vocab-booster/public/model"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/exercise/domain"
	"github.com/namhq1989/vocab-booster-utilities/language"
)

type ExerciseCollectionMapper struct{}

func (ExerciseCollectionMapper) FromModelToDomain(collection model.ExerciseCollections) (*domain.ExerciseCollection, error) {
	var result = &domain.ExerciseCollection{
		ID:                 collection.ID,
		Name:               collection.Name,
		Slug:               collection.Slug,
		Translated:         language.TranslatedLanguages{},
		Criteria:           collection.Criteria,
		IsFromSystem:       collection.IsFromSystem,
		Image:              collection.Image,
		Order:              int(collection.Order),
		StatsExercises:     int(collection.StatsExercises),
		LastStatsUpdatedAt: collection.LastStatsUpdatedAt,
	}

	if err := json.Unmarshal([]byte(collection.Translated), &result.Translated); err != nil {
		return nil, err
	}

	return result, nil
}

func (ExerciseCollectionMapper) FromDomainToModel(collection domain.ExerciseCollection) (*model.ExerciseCollections, error) {
	var result = &model.ExerciseCollections{
		ID:                 collection.ID,
		Name:               collection.Name,
		Slug:               collection.Slug,
		Translated:         "",
		Criteria:           collection.Criteria,
		IsFromSystem:       collection.IsFromSystem,
		Image:              collection.Image,
		Order:              int32(collection.Order),
		StatsExercises:     int32(collection.StatsExercises),
		LastStatsUpdatedAt: collection.LastStatsUpdatedAt,
	}

	if data, err := json.Marshal(collection.Translated); err != nil {
		return nil, err
	} else {
		result.Translated = string(data)
	}

	return result, nil
}
