package infrastructure

import (
	"github.com/namhq1989/vocab-booster-english-hub/internal/queue"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/exercise/domain"
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

func (r QueueRepository) UpdateUserExerciseCollectionStats(ctx *appcontext.AppContext, payload domain.QueueUpdateUserExerciseCollectionStatsPayload) error {
	return queue.EnqueueTask(ctx, r.queue, queue.TypeNames.UpdateUserExerciseCollectionStats, payload, -1)
}
