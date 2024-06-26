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

		CreateUserVocabularyCollection(ctx *appcontext.AppContext, req *vocabularypb.CreateUserVocabularyCollectionRequest) (*vocabularypb.CreateUserVocabularyCollectionResponse, error)
		UpdateUserVocabularyCollection(ctx *appcontext.AppContext, req *vocabularypb.UpdateUserVocabularyCollectionRequest) (*vocabularypb.UpdateUserVocabularyCollectionResponse, error)
	}
	App interface {
		Hubs
	}

	appHubHandler struct {
		hub.SearchVocabularyHandler

		hub.CreateUserVocabularyCollectionHandler
		hub.UpdateUserVocabularyCollectionHandler
	}
	Application struct {
		appHubHandler
	}
)

var _ App = (*Application)(nil)

func New(
	vocabularyRepository domain.VocabularyRepository,
	vocabularyExampleRepository domain.VocabularyExampleRepository,
	userVocabularyCollectionRepository domain.UserVocabularyCollectionRepository,
	aiRepository domain.AIRepository,
	scraperRepository domain.ScraperRepository,
	ttsRepository domain.TTSRepository,
	nlpRepository domain.NlpRepository,
	queueRepository domain.QueueRepository,
	cachingRepository domain.CachingRepository,
) *Application {
	return &Application{
		appHubHandler: appHubHandler{
			SearchVocabularyHandler: hub.NewSearchVocabularyHandler(vocabularyRepository, vocabularyExampleRepository, aiRepository, scraperRepository, ttsRepository, nlpRepository, queueRepository, cachingRepository),

			CreateUserVocabularyCollectionHandler: hub.NewCreateUserVocabularyCollectionHandler(userVocabularyCollectionRepository),
			UpdateUserVocabularyCollectionHandler: hub.NewUpdateUserVocabularyCollectionHandler(userVocabularyCollectionRepository),
		},
	}
}
