package infrastructure

import (
	"github.com/namhq1989/vocab-booster-english-hub/core/appcontext"
	"github.com/namhq1989/vocab-booster-english-hub/internal/queue"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/domain"
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
