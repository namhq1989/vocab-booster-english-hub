package exercise

import (
	"github.com/namhq1989/vocab-booster-english-hub/core/appcontext"
	"github.com/namhq1989/vocab-booster-english-hub/internal/monolith"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/exercise/application"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/exercise/grpc"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/exercise/infrastructure"
)

type Module struct{}

func (Module) Name() string {
	return "EXERCISE"
}

func (Module) Startup(ctx *appcontext.AppContext, mono monolith.Monolith) error {
	var (
		exerciseRepository = infrastructure.NewExerciseRepository(mono.Database())

		// app
		app = application.New(exerciseRepository)
	)

	// grpc server
	if err := grpc.RegisterServer(ctx, mono.RPC(), app); err != nil {
		return err
	}

	return nil
}
