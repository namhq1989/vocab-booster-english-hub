package grpc

import (
	"context"

	apperrors "github.com/namhq1989/vocab-booster-english-hub/internal/utils/error"

	"github.com/namhq1989/vocab-booster-english-hub/internal/genproto/exercisepb"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/exercise/application"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
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
	resp, err := s.app.NewExercise(appcontext.NewGRPC(bgCtx), req)
	if err != nil {
		return nil, apperrors.ToGrpcError(bgCtx, err)
	}
	return resp, nil
}

func (s server) UpdateExerciseAudio(bgCtx context.Context, req *exercisepb.UpdateExerciseAudioRequest) (*exercisepb.UpdateExerciseAudioResponse, error) {
	resp, err := s.app.UpdateExerciseAudio(appcontext.NewGRPC(bgCtx), req)
	if err != nil {
		return nil, apperrors.ToGrpcError(bgCtx, err)
	}
	return resp, nil
}

func (s server) AnswerExercise(bgCtx context.Context, req *exercisepb.AnswerExerciseRequest) (*exercisepb.AnswerExerciseResponse, error) {
	resp, err := s.app.AnswerExercise(appcontext.NewGRPC(bgCtx), req)
	if err != nil {
		return nil, apperrors.ToGrpcError(bgCtx, err)
	}
	return resp, nil
}

func (s server) GetUserExercises(bgCtx context.Context, req *exercisepb.GetUserExercisesRequest) (*exercisepb.GetUserExercisesResponse, error) {
	resp, err := s.app.GetUserExercises(appcontext.NewGRPC(bgCtx), req)
	if err != nil {
		return nil, apperrors.ToGrpcError(bgCtx, err)
	}
	return resp, nil
}

func (s server) CountUserReadyToReviewExercises(bgCtx context.Context, req *exercisepb.CountUserReadyToReviewExercisesRequest) (*exercisepb.CountUserReadyToReviewExercisesResponse, error) {
	resp, err := s.app.CountUserReadyToReviewExercises(appcontext.NewGRPC(bgCtx), req)
	if err != nil {
		return nil, apperrors.ToGrpcError(bgCtx, err)
	}
	return resp, nil
}

func (s server) GetUserReadyToReviewExercises(bgCtx context.Context, req *exercisepb.GetUserReadyToReviewExercisesRequest) (*exercisepb.GetUserReadyToReviewExercisesResponse, error) {
	resp, err := s.app.GetUserReadyToReviewExercises(appcontext.NewGRPC(bgCtx), req)
	if err != nil {
		return nil, apperrors.ToGrpcError(bgCtx, err)
	}
	return resp, nil
}

func (s server) ChangeExerciseFavorite(bgCtx context.Context, req *exercisepb.ChangeExerciseFavoriteRequest) (*exercisepb.ChangeExerciseFavoriteResponse, error) {
	resp, err := s.app.ChangeExerciseFavorite(appcontext.NewGRPC(bgCtx), req)
	if err != nil {
		return nil, apperrors.ToGrpcError(bgCtx, err)
	}
	return resp, nil
}

func (s server) GetUserFavoriteExercises(bgCtx context.Context, req *exercisepb.GetUserFavoriteExercisesRequest) (*exercisepb.GetUserFavoriteExercisesResponse, error) {
	resp, err := s.app.GetUserFavoriteExercises(appcontext.NewGRPC(bgCtx), req)
	if err != nil {
		return nil, apperrors.ToGrpcError(bgCtx, err)
	}
	return resp, nil
}

func (s server) GetUserStats(bgCtx context.Context, req *exercisepb.GetUserStatsRequest) (*exercisepb.GetUserStatsResponse, error) {
	resp, err := s.app.GetUserStats(appcontext.NewGRPC(bgCtx), req)
	if err != nil {
		return nil, apperrors.ToGrpcError(bgCtx, err)
	}
	return resp, nil
}

func (s server) GetExerciseCollections(bgCtx context.Context, req *exercisepb.GetExerciseCollectionsRequest) (*exercisepb.GetExerciseCollectionsResponse, error) {
	resp, err := s.app.GetExerciseCollections(appcontext.NewGRPC(bgCtx), req)
	if err != nil {
		return nil, apperrors.ToGrpcError(bgCtx, err)
	}
	return resp, nil
}

func (s server) GetUserRecentExercisesChart(bgCtx context.Context, req *exercisepb.GetUserRecentExercisesChartRequest) (*exercisepb.GetUserRecentExercisesChartResponse, error) {
	resp, err := s.app.GetUserRecentExercisesChart(appcontext.NewGRPC(bgCtx), req)
	if err != nil {
		return nil, apperrors.ToGrpcError(bgCtx, err)
	}
	return resp, nil
}
