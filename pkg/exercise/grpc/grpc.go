package grpc

import (
	"context"

	"github.com/namhq1989/vocab-booster-english-hub/core/appcontext"
	"github.com/namhq1989/vocab-booster-english-hub/internal/genproto/exercisepb"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/exercise/application"
	"google.golang.org/grpc"
)

type server struct {
	app application.App
	exercisepb.UnimplementedExerciseServiceServer
}

var _ exercisepb.ExerciseServiceServer = (*server)(nil)

func RegisterServer(_ *appcontext.AppContext, registrar grpc.ServiceRegistrar, app application.App) error {
	exercisepb.RegisterExerciseServiceServer(registrar, server{app: app})
	return nil
}

func (s server) NewExercise(bgCtx context.Context, req *exercisepb.NewExerciseRequest) (*exercisepb.NewExerciseResponse, error) {
	return s.app.NewExercise(appcontext.NewGRPC(bgCtx), req)
}
