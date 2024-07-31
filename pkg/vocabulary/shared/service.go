package shared

import "github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/domain"

type Service struct {
	vocabularyRepository        domain.VocabularyRepository
	vocabularyExampleRepository domain.VocabularyExampleRepository
	aiRepository                domain.AIRepository
	externalApiRepository       domain.ExternalApiRepository
	scraperRepository           domain.ScraperRepository
	ttsRepository               domain.TTSRepository
	nlpRepository               domain.NlpRepository
	queueRepository             domain.QueueRepository
	cachingRepository           domain.CachingRepository
}

func NewService(
	vocabularyRepository domain.VocabularyRepository,
	vocabularyExampleRepository domain.VocabularyExampleRepository,
	aiRepository domain.AIRepository,
	externalApiRepository domain.ExternalApiRepository,
	scraperRepository domain.ScraperRepository,
	ttsRepository domain.TTSRepository,
	nlpRepository domain.NlpRepository,
	queueRepository domain.QueueRepository,
	cachingRepository domain.CachingRepository,
) Service {
	return Service{
		vocabularyRepository:        vocabularyRepository,
		vocabularyExampleRepository: vocabularyExampleRepository,
		aiRepository:                aiRepository,
		externalApiRepository:       externalApiRepository,
		scraperRepository:           scraperRepository,
		ttsRepository:               ttsRepository,
		nlpRepository:               nlpRepository,
		queueRepository:             queueRepository,
		cachingRepository:           cachingRepository,
	}
}
