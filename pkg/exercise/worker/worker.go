package worker

import (
	"context"
	"fmt"

	"github.com/hibiken/asynq"
	"github.com/namhq1989/vocab-booster-english-hub/internal/queue"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/exercise/domain"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type (
	Handlers interface {
		UpdateUserExerciseCollectionStats(ctx *appcontext.AppContext, payload domain.QueueUpdateUserExerciseCollectionStatsPayload) error
	}
	Cronjob interface {
		UpdateExerciseCollectionStats(ctx *appcontext.AppContext, _ domain.QueueUpdateExerciseCollectionStatsPayload) error
	}
	Instance interface {
		Handlers
		Cronjob
	}

	workerHandlers struct {
		UpdateUserExerciseCollectionStatsHandler
	}
	workerCronjob struct {
		UpdateExerciseCollectionStatsHandler
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
	exerciseRepository domain.ExerciseRepository,
	exerciseCollectionRepository domain.ExerciseCollectionRepository,
	userExerciseCollectionStatusRepository domain.UserExerciseCollectionStatusRepository,
	cachingRepository domain.CachingRepository,
	service domain.Service,
) Worker {
	return Worker{
		queue: queue,

		workerHandlers: workerHandlers{
			UpdateUserExerciseCollectionStatsHandler: NewUpdateUserExerciseCollectionStatsHandler(
				userExerciseCollectionStatusRepository,
				cachingRepository,
				service,
			),
		},
		workerCronjob: workerCronjob{
			UpdateExerciseCollectionStatsHandler: NewUpdateExerciseCollectionStatsHandler(
				exerciseRepository,
				exerciseCollectionRepository,
				cachingRepository,
				service,
			),
		},
	}
}

func (w Worker) Start() {
	w.addCronjob()

	server := w.queue.GetServer()

	server.HandleFunc(w.queue.GenerateTypename(queue.TypeNames.UpdateUserExerciseCollectionStats), func(bgCtx context.Context, t *asynq.Task) error {
		return queue.ProcessTask[domain.QueueUpdateUserExerciseCollectionStatsPayload](bgCtx, t, queue.ParsePayload[domain.QueueUpdateUserExerciseCollectionStatsPayload], w.UpdateUserExerciseCollectionStats)
	})

	server.HandleFunc(w.queue.GenerateTypename(queue.TypeNames.UpdateExerciseCollectionStats), func(bgCtx context.Context, t *asynq.Task) error {
		return queue.ProcessTask[domain.QueueUpdateExerciseCollectionStatsPayload](bgCtx, t, queue.ParsePayload[domain.QueueUpdateExerciseCollectionStatsPayload], w.UpdateExerciseCollectionStats)
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
				Task:       w.queue.GenerateTypename(queue.TypeNames.UpdateExerciseCollectionStats),
				CronSpec:   "@every 30m",
				Payload:    domain.QueueUpdateExerciseCollectionStatsPayload{},
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
