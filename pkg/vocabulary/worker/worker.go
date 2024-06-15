package worker

import (
	"context"

	"github.com/hibiken/asynq"
	"github.com/namhq1989/vocab-booster-english-hub/core/appcontext"
	"github.com/namhq1989/vocab-booster-english-hub/internal/queue"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/domain"
)

type (
	Handlers interface {
		NewVocabularyCreated(ctx *appcontext.AppContext, payload domain.QueueNewVocabularyCreatedPayload) error
	}
	Instance interface {
		Handlers
	}

	workerHandlers struct {
		NewVocabularyCreatedHandler
	}
	Worker struct {
		queue queue.Operations
		workerHandlers
	}
)

var _ Instance = (*Worker)(nil)

func New(
	queue queue.Operations,
	vocabularyRepository domain.VocabularyRepository,
	vocabularyExampleRepository domain.VocabularyExampleRepository,
	ttsRepository domain.TTSRepository,
) Worker {
	return Worker{
		queue: queue,
		workerHandlers: workerHandlers{
			NewVocabularyCreatedHandler: NewNewVocabularyCreatedHandler(vocabularyRepository),
		},
	}
}

func (w Worker) Start() {
	server := w.queue.GetServer()

	server.HandleFunc(w.queue.GenerateTypename(queue.TypeNames.NewVocabularyCreated), func(bgCtx context.Context, t *asynq.Task) error {
		return queue.ProcessTask[domain.QueueNewVocabularyCreatedPayload](bgCtx, t, queue.ParsePayload[domain.QueueNewVocabularyCreatedPayload], w.NewVocabularyCreated)
	})
}
