package vocabulary

import (
	"github.com/namhq1989/vocab-booster-english-hub/internal/grpcclient"
	"github.com/namhq1989/vocab-booster-english-hub/internal/monolith"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/application"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/grpc"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/infrastructure"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/shared"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/worker"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type Module struct{}

func (Module) Name() string {
	return "VOCABULARY"
}

func (Module) Startup(ctx *appcontext.AppContext, mono monolith.Monolith) error {
	exerciseGRPCClient, err := grpcclient.NewExerciseClient(ctx, mono.Config().GRPCPort)
	if err != nil {
		return err
	}

	var (
		vocabularyRepository               = infrastructure.NewVocabularyRepository(mono.Database())
		vocabularyExampleRepository        = infrastructure.NewVocabularyExampleRepository(mono.Database())
		vocabularyScrapingItemRepository   = infrastructure.NewVocabularyScrapingItemRepository(mono.Database())
		userBookmarkedVocabularyRepository = infrastructure.NewUserBookmarkedVocabularyRepository(mono.Database())
		communitySentenceRepository        = infrastructure.NewCommunitySentenceRepository(mono.Database())
		communitySentenceDraftRepository   = infrastructure.NewCommunitySentenceDraftRepository(mono.Database())
		communitySentenceLikeRepository    = infrastructure.NewCommunitySentenceLikeRepository(mono.Database())
		verbConjugationRepository          = infrastructure.NewVerbConjugationRepository(mono.Database())
		wordOfTheDayRepository             = infrastructure.NewWordOfTheDayRepository(mono.Database())

		aiRepository          = infrastructure.NewAIRepository(mono.AI(), mono.NLP())
		externalApiRepository = infrastructure.NewExternalAPIRepository(mono.ExternalAPI(), mono.NLP())
		scraperRepository     = infrastructure.NewScraperRepository(mono.Scraper())
		ttsRepository         = infrastructure.NewTTSRepository(mono.TTS())
		nlpRepository         = infrastructure.NewNlpRepository(mono.NLP())
		queueRepository       = infrastructure.NewQueueRepository(mono.Queue())
		cachingRepository     = infrastructure.NewCachingRepository(mono.Caching())

		exerciseHub = infrastructure.NewExerciseHub(exerciseGRPCClient)

		service = shared.NewService(
			vocabularyRepository,
			vocabularyExampleRepository,
			aiRepository,
			externalApiRepository,
			scraperRepository,
			ttsRepository,
			nlpRepository,
			queueRepository,
			cachingRepository,
		)

		// app
		app = application.New(
			vocabularyRepository,
			userBookmarkedVocabularyRepository,
			wordOfTheDayRepository,
			communitySentenceRepository,
			communitySentenceDraftRepository,
			communitySentenceLikeRepository,
			nlpRepository,
			cachingRepository,
			service,
		)
	)

	// grpc server
	if err = grpc.RegisterServer(ctx, mono.RPC(), app); err != nil {
		return err
	}

	// worker
	w := worker.New(
		mono.Queue(),
		vocabularyRepository,
		vocabularyExampleRepository,
		vocabularyScrapingItemRepository,
		verbConjugationRepository,
		wordOfTheDayRepository,
		queueRepository,
		ttsRepository,
		aiRepository,
		exerciseHub,
		service,
	)
	w.Start()

	return nil
}
