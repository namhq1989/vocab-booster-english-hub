package shared

import "github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/domain"

type Service struct {
	vocabularyRepository        domain.VocabularyRepository
	vocabularyExampleRepository domain.VocabularyExampleRepository
	cachingRepository           domain.CachingRepository
}

func NewService(
	vocabularyRepository domain.VocabularyRepository,
	vocabularyExampleRepository domain.VocabularyExampleRepository,
	cachingRepository domain.CachingRepository,
) Service {
	return Service{
		vocabularyRepository:        vocabularyRepository,
		vocabularyExampleRepository: vocabularyExampleRepository,
		cachingRepository:           cachingRepository,
	}
}
