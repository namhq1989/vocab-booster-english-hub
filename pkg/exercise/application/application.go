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
	}
	Application struct {
		appHubHandler
	}
)

var _ App = (*Application)(nil)

func New(
	exerciseRepository domain.ExerciseRepository,
	userExerciseStatusRepository domain.UserExerciseStatusRepository,
) *Application {
	return &Application{
		appHubHandler: appHubHandler{
			NewExerciseHandler: hub.NewNewExerciseHandler(exerciseRepository),
			UpdateExerciseAudioHandler: hub.NewUpdateExerciseAudioHandler(
				exerciseRepository,
			),
			AnswerExerciseHandler: hub.NewAnswerExerciseHandler(
				exerciseRepository,
				userExerciseStatusRepository,
			),
			GetUserExercisesHandler: hub.NewGetUserExercisesHandler(
				exerciseRepository,
			),
			CountUserReadyToReviewExercisesHandler: hub.NewCountUserReadyToReviewExercisesHandler(
				userExerciseStatusRepository,
			),
			GetUserReadyToReviewExercisesHandler: hub.NewGetUserReadyToReviewExercisesHandler(
				userExerciseStatusRepository,
			),
			ChangeExerciseFavoriteHandler: hub.NewChangeExerciseFavoriteHandler(
				userExerciseStatusRepository,
			),
			GetUserFavoriteExercisesHandler: hub.NewGetUserFavoriteExercisesHandler(
				userExerciseStatusRepository,
			),
		},
	}
}
