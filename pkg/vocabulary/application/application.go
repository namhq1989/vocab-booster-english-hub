package application

import (
	"github.com/namhq1989/vocab-booster-english-hub/core/appcontext"
	"github.com/namhq1989/vocab-booster-english-hub/internal/genproto/vocabularypb"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/application/hub"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/domain"
)

type (
	Hubs interface {
		SearchVocabulary(ctx *appcontext.AppContext, req *vocabularypb.SearchVocabularyRequest) (*vocabularypb.SearchVocabularyResponse, error)
	}
	App interface {
		Hubs
	}

	appHubHandler struct {
		hub.SearchVocabularyHandler
	}
	Application struct {
		appHubHandler
	}
)

var _ App = (*Application)(nil)

func New(
	vocabularyRepository domain.VocabularyRepository,
	vocabularyExampleRepository domain.VocabularyExampleRepository,
	aiRepository domain.AIRepository,
	scraperRepository domain.ScraperRepository,
	ttsRepository domain.TTSRepository,
) *Application {
	return &Application{
		appHubHandler: appHubHandler{
			hub.NewSearchVocabularyHandler(vocabularyRepository, vocabularyExampleRepository, aiRepository, scraperRepository, ttsRepository),
		},
	}
}
