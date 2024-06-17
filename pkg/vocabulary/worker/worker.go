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
		NewVocabularyExampleCreated(ctx *appcontext.AppContext, payload domain.QueueNewVocabularyExampleCreatedPayload) error
		CreateVocabularyExampleAudio(ctx *appcontext.AppContext, payload domain.QueueCreateVocabularyExampleAudioPayload) error
		CreateVerbConjugation(ctx *appcontext.AppContext, payload domain.QueueCreateVerbConjugationPayload) error
		AddOtherVocabularyToScrapeQueue(ctx *appcontext.AppContext, payload domain.QueueAddOtherVocabularyToScrapeQueuePayload) error
	}
	Instance interface {
		Handlers
	}

	workerHandlers struct {
		NewVocabularyCreatedHandler
		NewVocabularyExampleCreatedHandler
		CreateVocabularyExampleAudioHandler
		CreateVerbConjugationHandler
		AddOtherVocabularyToScrapeQueueHandler
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
	vocabularyScrapeItemRepository domain.VocabularyScrapeItemRepository,
	verbConjugationRepository domain.VerbConjugationRepository,
	queueRepository domain.QueueRepository,
	ttsRepository domain.TTSRepository,
) Worker {
	return Worker{
		queue: queue,
		workerHandlers: workerHandlers{
			NewVocabularyCreatedHandler: NewNewVocabularyCreatedHandler(vocabularyRepository),
			NewVocabularyExampleCreatedHandler: NewNewVocabularyExampleCreatedHandler(
				vocabularyExampleRepository,
				queueRepository,
			),
			CreateVocabularyExampleAudioHandler: NewCreateVocabularyExampleAudioHandler(
				vocabularyExampleRepository,
				ttsRepository,
			),
			CreateVerbConjugationHandler: NewCreateVerbConjugationHandler(verbConjugationRepository),
			AddOtherVocabularyToScrapeQueueHandler: NewAddOtherVocabularyToScrapeQueueHandler(
				vocabularyRepository,
				vocabularyScrapeItemRepository,
			),
		},
	}
}

func (w Worker) Start() {
	server := w.queue.GetServer()

	server.HandleFunc(w.queue.GenerateTypename(queue.TypeNames.NewVocabularyCreated), func(bgCtx context.Context, t *asynq.Task) error {
		return queue.ProcessTask[domain.QueueNewVocabularyCreatedPayload](bgCtx, t, queue.ParsePayload[domain.QueueNewVocabularyCreatedPayload], w.NewVocabularyCreated)
	})

	server.HandleFunc(w.queue.GenerateTypename(queue.TypeNames.NewVocabularyExampleCreated), func(bgCtx context.Context, t *asynq.Task) error {
		return queue.ProcessTask[domain.QueueNewVocabularyExampleCreatedPayload](bgCtx, t, queue.ParsePayload[domain.QueueNewVocabularyExampleCreatedPayload], w.NewVocabularyExampleCreated)
	})

	server.HandleFunc(w.queue.GenerateTypename(queue.TypeNames.CreateVocabularyExampleAudio), func(bgCtx context.Context, t *asynq.Task) error {
		return queue.ProcessTask[domain.QueueCreateVocabularyExampleAudioPayload](bgCtx, t, queue.ParsePayload[domain.QueueCreateVocabularyExampleAudioPayload], w.CreateVocabularyExampleAudio)
	})

	server.HandleFunc(w.queue.GenerateTypename(queue.TypeNames.CreateVerbConjugation), func(bgCtx context.Context, t *asynq.Task) error {
		return queue.ProcessTask[domain.QueueCreateVerbConjugationPayload](bgCtx, t, queue.ParsePayload[domain.QueueCreateVerbConjugationPayload], w.CreateVerbConjugation)
	})

	server.HandleFunc(w.queue.GenerateTypename(queue.TypeNames.AddOtherVocabularyToScrapeQueue), func(bgCtx context.Context, t *asynq.Task) error {
		return queue.ProcessTask[domain.QueueAddOtherVocabularyToScrapeQueuePayload](bgCtx, t, queue.ParsePayload[domain.QueueAddOtherVocabularyToScrapeQueuePayload], w.AddOtherVocabularyToScrapeQueue)
	})
}
