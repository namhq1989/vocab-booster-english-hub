package worker

import (
	"context"
	"fmt"

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
		AddOtherVocabularyToScrapingQueue(ctx *appcontext.AppContext, payload domain.QueueAddOtherVocabularyToScrapingQueuePayload) error
	}
	Cronjob interface {
		AutoScrapingVocabulary(ctx *appcontext.AppContext, payload domain.QueueAutoScrapingVocabularyPayload) error
	}
	Instance interface {
		Handlers
		Cronjob
	}

	workerHandlers struct {
		NewVocabularyCreatedHandler
		NewVocabularyExampleCreatedHandler
		CreateVocabularyExampleAudioHandler
		CreateVerbConjugationHandler
		AddOtherVocabularyToScrapingQueueHandler
	}
	workerCronjob struct {
		AutoScrapingVocabularyHandler
	}
	Worker struct {
		queue queue.Operations
		workerHandlers
		workerCronjob
	}
)

var _ Instance = (*Worker)(nil)

func New(
	queue queue.Operations,
	vocabularyRepository domain.VocabularyRepository,
	vocabularyExampleRepository domain.VocabularyExampleRepository,
	vocabularyScrapingItemRepository domain.VocabularyScrapingItemRepository,
	verbConjugationRepository domain.VerbConjugationRepository,
	queueRepository domain.QueueRepository,
	ttsRepository domain.TTSRepository,
	aiRepository domain.AIRepository,
	scraperRepository domain.ScraperRepository,
	nlpRepository domain.NlpRepository,
	exerciseHub domain.ExerciseHub,
) Worker {
	return Worker{
		queue: queue,
		workerHandlers: workerHandlers{
			NewVocabularyCreatedHandler: NewNewVocabularyCreatedHandler(vocabularyRepository),
			NewVocabularyExampleCreatedHandler: NewNewVocabularyExampleCreatedHandler(
				vocabularyRepository,
				vocabularyExampleRepository,
				queueRepository,
				exerciseHub,
			),
			CreateVocabularyExampleAudioHandler: NewCreateVocabularyExampleAudioHandler(
				vocabularyExampleRepository,
				ttsRepository,
				exerciseHub,
			),
			CreateVerbConjugationHandler: NewCreateVerbConjugationHandler(verbConjugationRepository),
			AddOtherVocabularyToScrapingQueueHandler: NewAddOtherVocabularyToScrapingQueueHandler(
				vocabularyRepository,
				vocabularyScrapingItemRepository,
			),
		},
		workerCronjob: workerCronjob{
			AutoScrapingVocabularyHandler: NewAutoScrapingVocabularyHandler(
				vocabularyRepository,
				vocabularyExampleRepository,
				vocabularyScrapingItemRepository,
				aiRepository,
				scraperRepository,
				ttsRepository,
				nlpRepository,
				queueRepository,
			),
		},
	}
}

func (w Worker) Start() {
	w.addCronjob()

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

	server.HandleFunc(w.queue.GenerateTypename(queue.TypeNames.AddOtherVocabularyToScrapingQueue), func(bgCtx context.Context, t *asynq.Task) error {
		return queue.ProcessTask[domain.QueueAddOtherVocabularyToScrapingQueuePayload](bgCtx, t, queue.ParsePayload[domain.QueueAddOtherVocabularyToScrapingQueuePayload], w.AddOtherVocabularyToScrapingQueue)
	})

	server.HandleFunc(w.queue.GenerateTypename(queue.TypeNames.AutoScrapingVocabulary), func(bgCtx context.Context, t *asynq.Task) error {
		return queue.ProcessTask[domain.QueueAutoScrapingVocabularyPayload](bgCtx, t, queue.ParsePayload[domain.QueueAutoScrapingVocabularyPayload], w.AutoScrapingVocabulary)
	})
}

type cronjobData struct {
	Task       string      `json:"task"`
	CronSpec   string      `json:"cronSpec"`
	Payload    interface{} `json:"payload"`
	RetryTimes int         `json:"retryTimes"`
}

func (w Worker) addCronjob() {
	var (
		ctx  = appcontext.NewWorker(context.Background())
		jobs = []cronjobData{
			{
				Task:       w.queue.GenerateTypename(queue.TypeNames.AutoScrapingVocabulary),
				CronSpec:   "@every 1m",
				Payload:    domain.QueueAutoScrapingVocabularyPayload{},
				RetryTimes: 1,
			},
		}
	)

	for _, job := range jobs {
		entryID, err := w.queue.ScheduleTask(job.Task, job.Payload, job.CronSpec, job.RetryTimes)
		if err != nil {
			ctx.Logger().Error("error when initializing cronjob", err, appcontext.Fields{"job": job})
			panic(err)
		}

		ctx.Logger().Info(fmt.Sprintf("[cronjob] cronjob '%s' initialize successfully with cronSpec '%s' and retryTimes '%d'", job.Task, job.CronSpec, job.RetryTimes), appcontext.Fields{
			"entryId": entryID,
		})
	}
}
