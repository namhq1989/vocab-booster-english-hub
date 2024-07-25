package exercise

import (
	"github.com/namhq1989/vocab-booster-english-hub/internal/monolith"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/exercise/application"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/exercise/domain"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/exercise/grpc"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/exercise/infrastructure"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/exercise/shared"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/exercise/worker"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type Module struct{}

func (Module) Name() string {
	return "EXERCISE"
}

func (m Module) Startup(ctx *appcontext.AppContext, mono monolith.Monolith) error {
	var (
		exerciseRepository                      = infrastructure.NewExerciseRepository(mono.Database())
		exerciseCollectionRepository            = infrastructure.NewExerciseCollectionRepository(mono.Database())
		userExerciseStatusRepository            = infrastructure.NewUserExerciseStatusRepository(mono.Database())
		userExerciseCollectionStatusRepository  = infrastructure.NewUserExerciseCollectionStatusRepository(mono.Database())
		userExerciseInteractedHistoryRepository = infrastructure.NewUserExerciseInteractedHistoryRepository(mono.Database())
		queueRepository                         = infrastructure.NewQueueRepository(mono.Queue())
		cachingRepository                       = infrastructure.NewCachingRepository(mono.Caching())

		service = shared.NewService(exerciseCollectionRepository, cachingRepository)

		// app
		app = application.New(
			exerciseRepository,
			userExerciseStatusRepository,
			exerciseCollectionRepository,
			userExerciseInteractedHistoryRepository,
			cachingRepository,
			queueRepository,
			service,
		)
	)

	// grpc server
	if err := grpc.RegisterServer(ctx, mono.RPC(), app); err != nil {
		return err
	}

	// worker
	w := worker.New(
		mono.Queue(),
		exerciseRepository,
		exerciseCollectionRepository,
		userExerciseCollectionStatusRepository,
		userExerciseInteractedHistoryRepository,
		cachingRepository,
		service,
	)
	w.Start()

	m.init(ctx, exerciseCollectionRepository, cachingRepository)

	return nil
}

func (m Module) init(
	ctx *appcontext.AppContext,
	exerciseCollectionRepository domain.ExerciseCollectionRepository,
	cachingRepository domain.CachingRepository,
) {
	m.createExerciseCollections(ctx, exerciseCollectionRepository, cachingRepository)
}

func (Module) createExerciseCollections(
	ctx *appcontext.AppContext,
	exerciseCollectionRepository domain.ExerciseCollectionRepository,
	cachingRepository domain.CachingRepository,
) {
	for _, collection := range domain.SystemCollections {
		if err := exerciseCollectionRepository.UpsertExerciseCollection(ctx, collection); err != nil {
			panic(err)
		}
	}

	if err := cachingRepository.SetExerciseCollections(ctx, domain.SystemCollections); err != nil {
		panic(err)
	}
}
