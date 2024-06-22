package application

import (
	"github.com/namhq1989/vocab-booster-english-hub/core/appcontext"
	"github.com/namhq1989/vocab-booster-english-hub/internal/genproto/exercisepb"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/exercise/application/hub"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/exercise/domain"
)

type (
	Hubs interface {
		NewExercise(ctx *appcontext.AppContext, req *exercisepb.NewExerciseRequest) (*exercisepb.NewExerciseResponse, error)
		UpdateExerciseAudio(ctx *appcontext.AppContext, req *exercisepb.UpdateExerciseAudioRequest) (*exercisepb.UpdateExerciseAudioResponse, error)
		AnswerExercise(ctx *appcontext.AppContext, req *exercisepb.AnswerExerciseRequest) (*exercisepb.AnswerExerciseResponse, error)
	}
	App interface {
		Hubs
	}

	appHubHandler struct {
		hub.NewExerciseHandler
		hub.UpdateExerciseAudioHandler
		hub.AnswerExerciseHandler
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
		},
	}
}
