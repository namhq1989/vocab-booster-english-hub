package hub

import (
	"github.com/namhq1989/vocab-booster-english-hub/core/appcontext"
	"github.com/namhq1989/vocab-booster-english-hub/internal/genproto/exercisepb"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/exercise/domain"
)

type CountUserReadyToReviewExercisesHandler struct {
	userExerciseStatusRepository domain.UserExerciseStatusRepository
}

func NewCountUserReadyToReviewExercisesHandler(
	userExerciseStatusRepository domain.UserExerciseStatusRepository,
) CountUserReadyToReviewExercisesHandler {
	return CountUserReadyToReviewExercisesHandler{
		userExerciseStatusRepository: userExerciseStatusRepository,
	}
}

func (h CountUserReadyToReviewExercisesHandler) CountUserReadyToReviewExercises(ctx *appcontext.AppContext, req *exercisepb.CountUserReadyToReviewExercisesRequest) (*exercisepb.CountUserReadyToReviewExercisesResponse, error) {
	ctx.Logger().Info("[hub] new count user ready to review exercises request", appcontext.Fields{"userID": req.GetUserId()})

	ctx.Logger().Text("count in db")
	total, err := h.userExerciseStatusRepository.CountUserReadyToReviewExercises(ctx, req.GetUserId())
	if err != nil {
		ctx.Logger().Error("failed to count in db", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("done count user ready to review exercises request")
	return &exercisepb.CountUserReadyToReviewExercisesResponse{Total: int32(total)}, nil
}
