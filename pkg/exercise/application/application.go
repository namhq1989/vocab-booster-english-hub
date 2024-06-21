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
	}
	App interface {
		Hubs
	}

	appHubHandler struct {
		hub.NewExerciseHandler
	}
	Application struct {
		appHubHandler
	}
)

var _ App = (*Application)(nil)

func New(
	exerciseRepository domain.ExerciseRepository,
) *Application {
	return &Application{
		appHubHandler: appHubHandler{
			NewExerciseHandler: hub.NewNewExerciseHandler(exerciseRepository),
		},
	}
}
