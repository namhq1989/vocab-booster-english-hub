package hub

import (
	"github.com/namhq1989/vocab-booster-english-hub/internal/genproto/exercisepb"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/exercise/domain"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type GetUserStatsHandler struct {
	userExerciseStatusRepository domain.UserExerciseStatusRepository
}

func NewGetUserStatsHandler(
	userExerciseStatusRepository domain.UserExerciseStatusRepository,
) GetUserStatsHandler {
	return GetUserStatsHandler{
		userExerciseStatusRepository: userExerciseStatusRepository,
	}
}

func (h GetUserStatsHandler) GetUserStats(ctx *appcontext.AppContext, req *exercisepb.GetUserStatsRequest) (*exercisepb.GetUserStatsResponse, error) {
	ctx.Logger().Info("[hub] new get user stats request", appcontext.Fields{"userID": req.GetUserId(), "timezone": req.GetTimezone()})

	ctx.Logger().Text("find in db")
	stats, err := h.userExerciseStatusRepository.FindUserStats(ctx, req.GetUserId(), req.GetTimezone())
	if err != nil {
		ctx.Logger().Error("failed to find in db", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("done get user stats request")
	return &exercisepb.GetUserStatsResponse{
		Mastered:         int32(stats.Mastered),
		WaitingForReview: int32(stats.WaitingForReview),
	}, nil
}
