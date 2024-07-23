package hub

import (
	"github.com/namhq1989/vocab-booster-english-hub/internal/genproto/exercisepb"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/exercise/domain"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/exercise/dto"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type GetUserReadyForReviewExercisesHandler struct {
	userExerciseStatusRepository domain.UserExerciseStatusRepository
}

func NewGetUserReadyForReviewExercisesHandler(
	userExerciseStatusRepository domain.UserExerciseStatusRepository,
) GetUserReadyForReviewExercisesHandler {
	return GetUserReadyForReviewExercisesHandler{
		userExerciseStatusRepository: userExerciseStatusRepository,
	}
}

func (h GetUserReadyForReviewExercisesHandler) GetUserReadyForReviewExercises(ctx *appcontext.AppContext, req *exercisepb.GetUserReadyForReviewExercisesRequest) (*exercisepb.GetUserReadyForReviewExercisesResponse, error) {
	ctx.Logger().Info("[hub] new get user ready for review exercises request", appcontext.Fields{"userID": req.GetUserId(), "lang": req.GetLang(), "timezone": req.GetTimezone()})

	ctx.Logger().Text("new user exercise filter")
	filter, err := domain.NewUserExerciseFilter(req.GetUserId(), "", req.GetLang(), req.GetTimezone())
	if err != nil {
		ctx.Logger().Error("failed to create new user exercise filter", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("find exercises in db")
	exercises, err := h.userExerciseStatusRepository.FindUserReadyToReviewExercises(ctx, *filter)
	if err != nil {
		ctx.Logger().Error("failed to find exercises in db", err, appcontext.Fields{})
		return nil, err
	}

	result := dto.ConvertUserExercisesFromDomainToGrpc(exercises, req.GetLang())
	ctx.Logger().Text("done get user ready to review exercises request")

	return &exercisepb.GetUserReadyForReviewExercisesResponse{
		Exercises: result,
	}, nil
}
