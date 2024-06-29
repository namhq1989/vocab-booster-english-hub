package vocabulary

import (
	"github.com/namhq1989/vocab-booster-english-hub/core/appcontext"
	"github.com/namhq1989/vocab-booster-english-hub/internal/grpcclient"
	"github.com/namhq1989/vocab-booster-english-hub/internal/monolith"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/application"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/grpc"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/infrastructure"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/worker"
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
		vocabularyRepository              = infrastructure.NewVocabularyRepository(mono.Database())
		vocabularyExampleRepository       = infrastructure.NewVocabularyExampleRepository(mono.Database())
		vocabularyScrapingItemRepository  = infrastructure.NewVocabularyScrapingItemRepository(mono.Database())
		communitySentenceRepository       = infrastructure.NewCommunitySentenceRepository(mono.Database())
		communitySentenceDraftRepository  = infrastructure.NewCommunitySentenceDraftRepository(mono.Database())
		verbConjugationRepository         = infrastructure.NewVerbConjugationRepository(mono.Database())
		collectionRepository              = infrastructure.NewCollectionRepository(mono.Database())
		collectionAndVocabularyRepository = infrastructure.NewCollectionAndVocabularyRepository(mono.Database())

		aiRepository      = infrastructure.NewAIRepository(mono.AI(), mono.NLP())
		scraperRepository = infrastructure.NewScraperRepository(mono.Scraper())
		ttsRepository     = infrastructure.NewTTSRepository(mono.TTS())
		nlpRepository     = infrastructure.NewNlpRepository(mono.NLP())
		queueRepository   = infrastructure.NewQueueRepository(mono.Queue())
		cachingRepository = infrastructure.NewCachingRepository(mono.Caching())

		exerciseHub = infrastructure.NewExerciseHub(exerciseGRPCClient)

		// app
		app = application.New(
			vocabularyRepository,
			vocabularyExampleRepository,
			communitySentenceRepository,
			communitySentenceDraftRepository,
			collectionRepository,
			collectionAndVocabularyRepository,
			aiRepository,
			scraperRepository,
			ttsRepository,
			nlpRepository,
			queueRepository,
			cachingRepository,
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
		queueRepository,
		ttsRepository,
		aiRepository,
		scraperRepository,
		nlpRepository,
		exerciseHub,
	)
	w.Start()

	return nil
}
