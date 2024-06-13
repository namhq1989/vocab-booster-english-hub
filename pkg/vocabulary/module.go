package vocabulary

import (
	"github.com/namhq1989/vocab-booster-english-hub/core/appcontext"
	"github.com/namhq1989/vocab-booster-english-hub/internal/monolith"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/application"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/grpc"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/infrastructure"
)

type Module struct{}

func (Module) Name() string {
	return "VOCABULARY"
}

func (Module) Startup(ctx *appcontext.AppContext, mono monolith.Monolith) error {
	var (
		vocabularyRepository        = infrastructure.NewVocabularyRepository(mono.Database())
		vocabularyExampleRepository = infrastructure.NewVocabularyExampleRepository(mono.Database())
		aiRepository                = infrastructure.NewAIRepository(mono.AI())
		scraperRepository           = infrastructure.NewScraperRepository(mono.Scraper())
		ttsRepository               = infrastructure.NewTTSRepository(mono.TTS())

		// app
		app = application.New(vocabularyRepository, vocabularyExampleRepository, aiRepository, scraperRepository, ttsRepository)
	)

	// grpc server
	if err := grpc.RegisterServer(ctx, app, mono.RPC()); err != nil {
		return err
	}

	return nil
}
