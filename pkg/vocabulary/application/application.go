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
		BookmarkVocabulary(ctx *appcontext.AppContext, req *vocabularypb.BookmarkVocabularyRequest) (*vocabularypb.BookmarkVocabularyResponse, error)
		GetUserBookmarkedVocabularies(ctx *appcontext.AppContext, req *vocabularypb.GetUserBookmarkedVocabulariesRequest) (*vocabularypb.GetUserBookmarkedVocabulariesResponse, error)
		GetWordOfTheDay(ctx *appcontext.AppContext, req *vocabularypb.GetWordOfTheDayRequest) (*vocabularypb.GetWordOfTheDayResponse, error)

		CreateCommunitySentenceDraft(ctx *appcontext.AppContext, req *vocabularypb.CreateCommunitySentenceDraftRequest) (*vocabularypb.CreateCommunitySentenceDraftResponse, error)
		UpdateCommunitySentenceDraft(ctx *appcontext.AppContext, req *vocabularypb.UpdateCommunitySentenceDraftRequest) (*vocabularypb.UpdateCommunitySentenceDraftResponse, error)
		PromoteCommunitySentenceDraft(ctx *appcontext.AppContext, req *vocabularypb.PromoteCommunitySentenceDraftRequest) (*vocabularypb.PromoteCommunitySentenceDraftResponse, error)
		LikeCommunitySentence(ctx *appcontext.AppContext, req *vocabularypb.LikeCommunitySentenceRequest) (*vocabularypb.LikeCommunitySentenceResponse, error)
		GetCommunitySentences(ctx *appcontext.AppContext, req *vocabularypb.GetCommunitySentencesRequest) (*vocabularypb.GetCommunitySentencesResponse, error)
		GetCommunitySentence(ctx *appcontext.AppContext, req *vocabularypb.GetCommunitySentenceRequest) (*vocabularypb.GetCommunitySentenceResponse, error)
		GetUserDraftCommunitySentences(ctx *appcontext.AppContext, req *vocabularypb.GetUserDraftCommunitySentencesRequest) (*vocabularypb.GetUserDraftCommunitySentencesResponse, error)
	}
	App interface {
		Hubs
	}

	appHubHandler struct {
		hub.SearchVocabularyHandler
		hub.BookmarkVocabularyHandler
		hub.GetUserBookmarkedVocabulariesHandler
		hub.GetWordOfTheDayHandler

		hub.CreateCommunitySentenceDraftHandler
		hub.UpdateCommunitySentenceDraftHandler
		hub.PromoteCommunitySentenceDraftHandler
		hub.LikeCommunitySentenceHandler
		hub.GetCommunitySentencesHandler
		hub.GetCommunitySentenceHandler
		hub.GetUserDraftCommunitySentencesHandler
	}
	Application struct {
		appHubHandler
	}
)

var _ App = (*Application)(nil)

func New(
	vocabularyRepository domain.VocabularyRepository,
	userBookmarkedVocabularyRepository domain.UserBookmarkedVocabularyRepository,
	wordOfTheDayRepository domain.WordOfTheDayRepository,
	communitySentenceRepository domain.CommunitySentenceRepository,
	communitySentenceDraftRepository domain.CommunitySentenceDraftRepository,
	communitySentenceLikeRepository domain.CommunitySentenceLikeRepository,
	nlpRepository domain.NlpRepository,
	cachingRepository domain.CachingRepository,
	service domain.Service,
) *Application {
	return &Application{
		appHubHandler: appHubHandler{
			SearchVocabularyHandler: hub.NewSearchVocabularyHandler(
				userBookmarkedVocabularyRepository,
				service,
			),

			BookmarkVocabularyHandler: hub.NewBookmarkVocabularyHandler(
				vocabularyRepository,
				userBookmarkedVocabularyRepository,
			),

			GetUserBookmarkedVocabulariesHandler: hub.NewGetUserBookmarkedVocabulariesHandler(
				userBookmarkedVocabularyRepository,
			),

			GetWordOfTheDayHandler: hub.NewGetWordOfTheDayHandler(
				wordOfTheDayRepository,
				cachingRepository,
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
				vocabularyRepository,
				communitySentenceRepository,
				communitySentenceDraftRepository,
				nlpRepository,
			),
			LikeCommunitySentenceHandler: hub.NewLikeCommunitySentenceHandler(
				communitySentenceRepository,
				communitySentenceLikeRepository,
			),
			GetCommunitySentencesHandler: hub.NewGetCommunitySentencesHandler(
				communitySentenceRepository,
			),
			GetCommunitySentenceHandler: hub.NewGetCommunitySentenceHandler(
				communitySentenceRepository,
			),
			GetUserDraftCommunitySentencesHandler: hub.NewGetUserDraftCommunitySentencesHandler(
				communitySentenceDraftRepository,
			),
		},
	}
}
