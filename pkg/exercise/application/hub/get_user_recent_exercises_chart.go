package hub

import (
	"github.com/namhq1989/vocab-booster-english-hub/internal/genproto/exercisepb"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/exercise/domain"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/exercise/dto"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type GetUserRecentExercisesChartHandler struct {
	userExerciseInteractedHistory domain.UserExerciseInteractedHistoryRepository
}

func NewGetUserRecentExercisesChartHandler(
	userExerciseInteractedHistory domain.UserExerciseInteractedHistoryRepository,
) GetUserRecentExercisesChartHandler {
	return GetUserRecentExercisesChartHandler{
		userExerciseInteractedHistory: userExerciseInteractedHistory,
	}
}

func (h GetUserRecentExercisesChartHandler) GetUserRecentExercisesChart(ctx *appcontext.AppContext, req *exercisepb.GetUserRecentExercisesChartRequest) (*exercisepb.GetUserRecentExercisesChartResponse, error) {
	ctx.Logger().Info("[hub] new get user recent exercises chart request", appcontext.Fields{"userID": req.GetUserId(), "timezone": req.GetTimezone(), "from": req.GetFrom().AsTime(), "to": req.GetTo().AsTime()})

	ctx.Logger().Text("find in db")
	uaes, err := h.userExerciseInteractedHistory.AggregateUserExercisesInTimeRange(ctx, req.GetUserId(), req.GetTimezone(), req.GetFrom().AsTime(), req.GetTo().AsTime())
	if err != nil {
		ctx.Logger().Error("failed to find in db", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("convert to response data")
	result := dto.ConvertUserAggregatedExercisesFromDomainToGrpc(uaes)
	ctx.Logger().Text("done get user recent exercises chart request")
	return &exercisepb.GetUserRecentExercisesChartResponse{
		Exercises: result,
	}, nil
}
