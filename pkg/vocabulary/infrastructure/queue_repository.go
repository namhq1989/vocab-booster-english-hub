package infrastructure

import (
	"github.com/namhq1989/vocab-booster-english-hub/internal/queue"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/domain"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type QueueRepository struct {
	queue queue.Operations
}

func NewQueueRepository(queue queue.Operations) QueueRepository {
	return QueueRepository{
		queue: queue,
	}
}

func (r QueueRepository) NewVocabularyCreated(ctx *appcontext.AppContext, payload domain.QueueNewVocabularyCreatedPayload) error {
	return queue.EnqueueTask(ctx, r.queue, queue.TypeNames.NewVocabularyCreated, payload, -1)
}

func (r QueueRepository) NewVocabularyExampleCreated(ctx *appcontext.AppContext, payload domain.QueueNewVocabularyExampleCreatedPayload) error {
	return queue.EnqueueTask(ctx, r.queue, queue.TypeNames.NewVocabularyExampleCreated, payload, -1)
}

func (r QueueRepository) CreateVocabularyExampleAudio(ctx *appcontext.AppContext, payload domain.QueueCreateVocabularyExampleAudioPayload) error {
	return queue.EnqueueTask(ctx, r.queue, queue.TypeNames.CreateVocabularyExampleAudio, payload, -1)
}

func (r QueueRepository) CreateVerbConjugation(ctx *appcontext.AppContext, payload domain.QueueCreateVerbConjugationPayload) error {
	return queue.EnqueueTask(ctx, r.queue, queue.TypeNames.CreateVerbConjugation, payload, -1)
}

func (r QueueRepository) AddOtherVocabularyToScrapingQueue(ctx *appcontext.AppContext, payload domain.QueueAddOtherVocabularyToScrapingQueuePayload) error {
	return queue.EnqueueTask(ctx, r.queue, queue.TypeNames.AddOtherVocabularyToScrapingQueue, payload, -1)
}
