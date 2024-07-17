package application

import (
	"github.com/namhq1989/vocab-booster-english-hub/internal/genproto/exercisepb"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/exercise/application/hub"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/exercise/domain"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type (
	Hubs interface {
		NewExercise(ctx *appcontext.AppContext, req *exercisepb.NewExerciseRequest) (*exercisepb.NewExerciseResponse, error)
		UpdateExerciseAudio(ctx *appcontext.AppContext, req *exercisepb.UpdateExerciseAudioRequest) (*exercisepb.UpdateExerciseAudioResponse, error)
		AnswerExercise(ctx *appcontext.AppContext, req *exercisepb.AnswerExerciseRequest) (*exercisepb.AnswerExerciseResponse, error)
		GetUserExercises(ctx *appcontext.AppContext, req *exercisepb.GetUserExercisesRequest) (*exercisepb.GetUserExercisesResponse, error)
		CountUserReadyToReviewExercises(ctx *appcontext.AppContext, req *exercisepb.CountUserReadyToReviewExercisesRequest) (*exercisepb.CountUserReadyToReviewExercisesResponse, error)
		GetUserReadyToReviewExercises(ctx *appcontext.AppContext, req *exercisepb.GetUserReadyToReviewExercisesRequest) (*exercisepb.GetUserReadyToReviewExercisesResponse, error)
		ChangeExerciseFavorite(ctx *appcontext.AppContext, req *exercisepb.ChangeExerciseFavoriteRequest) (*exercisepb.ChangeExerciseFavoriteResponse, error)
		GetUserFavoriteExercises(ctx *appcontext.AppContext, req *exercisepb.GetUserFavoriteExercisesRequest) (*exercisepb.GetUserFavoriteExercisesResponse, error)
		GetUserStats(ctx *appcontext.AppContext, req *exercisepb.GetUserStatsRequest) (*exercisepb.GetUserStatsResponse, error)
		GetExerciseCollections(ctx *appcontext.AppContext, req *exercisepb.GetExerciseCollectionsRequest) (*exercisepb.GetExerciseCollectionsResponse, error)
	}
	App interface {
		Hubs
	}

	appHubHandler struct {
		hub.NewExerciseHandler
		hub.UpdateExerciseAudioHandler
		hub.AnswerExerciseHandler
		hub.GetUserExercisesHandler
		hub.CountUserReadyToReviewExercisesHandler
		hub.GetUserReadyToReviewExercisesHandler
		hub.ChangeExerciseFavoriteHandler
		hub.GetUserFavoriteExercisesHandler
		hub.GetUserStatsHandler
		hub.GetExerciseCollectionsHandler
	}
	Application struct {
		appHubHandler
	}
)

var _ App = (*Application)(nil)

func New(
	exerciseRepository domain.ExerciseRepository,
	userExerciseStatusRepository domain.UserExerciseStatusRepository,
	exerciseCollectionRepository domain.ExerciseCollectionRepository,
	cachingRepository domain.CachingRepository,
	queueRepository domain.QueueRepository,
	service domain.Service,
) *Application {
	return &Application{
		appHubHandler: appHubHandler{
			NewExerciseHandler:         hub.NewNewExerciseHandler(exerciseRepository),
			UpdateExerciseAudioHandler: hub.NewUpdateExerciseAudioHandler(exerciseRepository),
			AnswerExerciseHandler: hub.NewAnswerExerciseHandler(
				exerciseRepository,
				userExerciseStatusRepository,
				queueRepository,
			),
			GetUserExercisesHandler:                hub.NewGetUserExercisesHandler(exerciseRepository, service),
			CountUserReadyToReviewExercisesHandler: hub.NewCountUserReadyToReviewExercisesHandler(userExerciseStatusRepository),
			GetUserReadyToReviewExercisesHandler:   hub.NewGetUserReadyToReviewExercisesHandler(userExerciseStatusRepository),
			ChangeExerciseFavoriteHandler:          hub.NewChangeExerciseFavoriteHandler(userExerciseStatusRepository),
			GetUserFavoriteExercisesHandler:        hub.NewGetUserFavoriteExercisesHandler(userExerciseStatusRepository),
			GetUserStatsHandler:                    hub.NewGetUserStatsHandler(userExerciseStatusRepository),
			GetExerciseCollectionsHandler:          hub.NewGetExerciseCollectionsHandler(exerciseCollectionRepository, cachingRepository),
		},
	}
}
