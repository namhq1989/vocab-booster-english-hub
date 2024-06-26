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

		CreateCollection(ctx *appcontext.AppContext, req *vocabularypb.CreateCollectionRequest) (*vocabularypb.CreateCollectionResponse, error)
		UpdateCollection(ctx *appcontext.AppContext, req *vocabularypb.UpdateCollectionRequest) (*vocabularypb.UpdateCollectionResponse, error)
		AddVocabularyToCollection(ctx *appcontext.AppContext, req *vocabularypb.AddVocabularyToCollectionRequest) (*vocabularypb.AddVocabularyToCollectionResponse, error)
		RemoveVocabularyFromCollection(ctx *appcontext.AppContext, req *vocabularypb.RemoveVocabularyFromCollectionRequest) (*vocabularypb.RemoveVocabularyFromCollectionResponse, error)
	}
	App interface {
		Hubs
	}

	appHubHandler struct {
		hub.SearchVocabularyHandler

		hub.CreateCollectionHandler
		hub.UpdateCollectionHandler
		hub.AddVocabularyToCollectionHandler
		hub.RemoveVocabularyFromCollectionHandler
	}
	Application struct {
		appHubHandler
	}
)

var _ App = (*Application)(nil)

func New(
	vocabularyRepository domain.VocabularyRepository,
	vocabularyExampleRepository domain.VocabularyExampleRepository,
	collectionRepository domain.CollectionRepository,
	collectionAndVocabularyRepository domain.CollectionAndVocabularyRepository,
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

			CreateCollectionHandler: hub.NewCreateCollectionHandler(collectionRepository),
			UpdateCollectionHandler: hub.NewUpdateCollectionHandler(collectionRepository),
			AddVocabularyToCollectionHandler: hub.NewAddVocabularyToCollectionHandler(
				vocabularyRepository,
				collectionRepository,
				collectionAndVocabularyRepository,
			),
			RemoveVocabularyFromCollectionHandler: hub.NewRemoveVocabularyFromCollectionHandler(
				vocabularyRepository,
				collectionRepository,
				collectionAndVocabularyRepository,
			),
		},
	}
}
