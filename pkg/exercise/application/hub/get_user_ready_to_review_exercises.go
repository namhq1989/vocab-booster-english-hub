package hub

import (
	"github.com/namhq1989/vocab-booster-english-hub/core/appcontext"
	"github.com/namhq1989/vocab-booster-english-hub/internal/genproto/exercisepb"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/exercise/domain"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/exercise/dto"
)

type GetUserReadyToReviewExercisesHandler struct {
	userExerciseStatusRepository domain.UserExerciseStatusRepository
}

func NewGetUserReadyToReviewExercisesHandler(
	userExerciseStatusRepository domain.UserExerciseStatusRepository,
) GetUserReadyToReviewExercisesHandler {
	return GetUserReadyToReviewExercisesHandler{
		userExerciseStatusRepository: userExerciseStatusRepository,
	}
}

func (h GetUserReadyToReviewExercisesHandler) GetUserReadyToReviewExercises(ctx *appcontext.AppContext, req *exercisepb.GetUserReadyToReviewExercisesRequest) (*exercisepb.GetUserReadyToReviewExercisesResponse, error) {
	ctx.Logger().Info("[hub] new get user ready to review exercises request", appcontext.Fields{"userID": req.GetUserId(), "level": req.GetLevel(), "lang": req.GetLang()})

	ctx.Logger().Text("new user exercise filter")
	filter, err := domain.NewUserExerciseFilter(req.GetUserId(), req.GetLevel(), req.GetLang())
	if err != nil {
		ctx.Logger().Error("failed to create new user exercise filter", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("find exercises in db")
	exercises, err := h.userExerciseStatusRepository.GetUserReadyToReviewExercises(ctx, *filter)
	if err != nil {
		ctx.Logger().Error("failed to find exercises in db", err, appcontext.Fields{})
		return nil, err
	}

	result := dto.ConvertUserExerciseFromDomainToGrpc(exercises)
	ctx.Logger().Text("done get user ready to review exercises request")

	return &exercisepb.GetUserReadyToReviewExercisesResponse{
		Exercises: result,
	}, nil
}
