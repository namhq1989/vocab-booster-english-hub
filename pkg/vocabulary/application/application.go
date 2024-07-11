package application

import (
	"github.com/namhq1989/vocab-booster-english-hub/internal/genproto/vocabularypb"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/application/hub"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/domain"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type (
	Hubs interface {
		SearchVocabulary(ctx *appcontext.AppContext, req *vocabularypb.SearchVocabularyRequest) (*vocabularypb.SearchVocabularyResponse, error)

		CreateCollection(ctx *appcontext.AppContext, req *vocabularypb.CreateCollectionRequest) (*vocabularypb.CreateCollectionResponse, error)
		UpdateCollection(ctx *appcontext.AppContext, req *vocabularypb.UpdateCollectionRequest) (*vocabularypb.UpdateCollectionResponse, error)
		AddVocabularyToCollection(ctx *appcontext.AppContext, req *vocabularypb.AddVocabularyToCollectionRequest) (*vocabularypb.AddVocabularyToCollectionResponse, error)
		RemoveVocabularyFromCollection(ctx *appcontext.AppContext, req *vocabularypb.RemoveVocabularyFromCollectionRequest) (*vocabularypb.RemoveVocabularyFromCollectionResponse, error)

		CreateCommunitySentenceDraft(ctx *appcontext.AppContext, req *vocabularypb.CreateCommunitySentenceDraftRequest) (*vocabularypb.CreateCommunitySentenceDraftResponse, error)
		UpdateCommunitySentenceDraft(ctx *appcontext.AppContext, req *vocabularypb.UpdateCommunitySentenceDraftRequest) (*vocabularypb.UpdateCommunitySentenceDraftResponse, error)
		PromoteCommunitySentenceDraft(ctx *appcontext.AppContext, req *vocabularypb.PromoteCommunitySentenceDraftRequest) (*vocabularypb.PromoteCommunitySentenceDraftResponse, error)
		LikeCommunitySentence(ctx *appcontext.AppContext, req *vocabularypb.LikeCommunitySentenceRequest) (*vocabularypb.LikeCommunitySentenceResponse, error)
		GetVocabularyCommunitySentences(ctx *appcontext.AppContext, req *vocabularypb.GetVocabularyCommunitySentencesRequest) (*vocabularypb.GetVocabularyCommunitySentencesResponse, error)
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

		hub.CreateCommunitySentenceDraftHandler
		hub.UpdateCommunitySentenceDraftHandler
		hub.PromoteCommunitySentenceDraftHandler
		hub.LikeCommunitySentenceHandler
		hub.GetVocabularyCommunitySentencesHandler
	}
	Application struct {
		appHubHandler
	}
)

var _ App = (*Application)(nil)

func New(
	vocabularyRepository domain.VocabularyRepository,
	vocabularyExampleRepository domain.VocabularyExampleRepository,
	communitySentenceRepository domain.CommunitySentenceRepository,
	communitySentenceDraftRepository domain.CommunitySentenceDraftRepository,
	communitySentenceLikeRepository domain.CommunitySentenceLikeRepository,
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

			CreateCommunitySentenceDraftHandler: hub.NewCreateCommunitySentenceDraftHandler(
				vocabularyRepository,
				communitySentenceDraftRepository,
				nlpRepository,
			),
			UpdateCommunitySentenceDraftHandler: hub.NewUpdateCommunitySentenceDraftHandler(
				communitySentenceDraftRepository,
				nlpRepository,
			),
			PromoteCommunitySentenceDraftHandler: hub.NewPromoteCommunitySentenceDraftHandler(
				communitySentenceRepository,
				communitySentenceDraftRepository,
				nlpRepository,
			),
			LikeCommunitySentenceHandler: hub.NewLikeCommunitySentenceHandler(
				communitySentenceRepository,
				communitySentenceLikeRepository,
			),
			GetVocabularyCommunitySentencesHandler: hub.NewGetVocabularyCommunitySentencesHandler(
				communitySentenceRepository,
			),
		},
	}
}
