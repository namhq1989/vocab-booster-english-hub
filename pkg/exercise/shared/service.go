package shared

import "github.com/namhq1989/vocab-booster-english-hub/pkg/exercise/domain"

type Service struct {
	exerciseCollectionRepository domain.ExerciseCollectionRepository
	cachingRepository            domain.CachingRepository
}

func NewService(
	exerciseCollectionRepository domain.ExerciseCollectionRepository,
	cachingRepository domain.CachingRepository,
) Service {
	return Service{
		exerciseCollectionRepository: exerciseCollectionRepository,
		cachingRepository:            cachingRepository,
	}
}
