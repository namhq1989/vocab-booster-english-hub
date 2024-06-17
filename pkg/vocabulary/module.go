package vocabulary

import (
	"github.com/namhq1989/vocab-booster-english-hub/core/appcontext"
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
	var (
		vocabularyRepository           = infrastructure.NewVocabularyRepository(mono.Database())
		vocabularyExampleRepository    = infrastructure.NewVocabularyExampleRepository(mono.Database())
		vocabularyScrapeItemRepository = infrastructure.NewVocabularyScrapeItemRepository(mono.Database())
		verbConjugationRepository      = infrastructure.NewVerbConjugationRepository(mono.Database())
		aiRepository                   = infrastructure.NewAIRepository(mono.AI())
		scraperRepository              = infrastructure.NewScraperRepository(mono.Scraper())
		ttsRepository                  = infrastructure.NewTTSRepository(mono.TTS())
		nlpRepository                  = infrastructure.NewNlpRepository(mono.NLP())
		queueRepository                = infrastructure.NewQueueRepository(mono.Queue())

		// app
		app = application.New(
			vocabularyRepository,
			vocabularyExampleRepository,
			aiRepository,
			scraperRepository,
			ttsRepository,
			nlpRepository,
			queueRepository,
		)
	)

	// grpc server
	if err := grpc.RegisterServer(ctx, app, mono.RPC()); err != nil {
		return err
	}

	// worker
	w := worker.New(
		mono.Queue(),
		vocabularyRepository,
		vocabularyExampleRepository,
		vocabularyScrapeItemRepository,
		verbConjugationRepository,
		queueRepository,
		ttsRepository,
		aiRepository,
		scraperRepository,
		nlpRepository,
	)
	w.Start()

	return nil
}
