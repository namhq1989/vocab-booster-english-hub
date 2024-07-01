package hub

import (
	"github.com/namhq1989/vocab-booster-english-hub/core/appcontext"
	"github.com/namhq1989/vocab-booster-english-hub/internal/genproto/exercisepb"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/exercise/domain"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/exercise/dto"
)

type GetUserExercisesHandler struct {
	exerciseRepository domain.ExerciseRepository
}

func NewGetUserExercisesHandler(exerciseRepository domain.ExerciseRepository) GetUserExercisesHandler {
	return GetUserExercisesHandler{
		exerciseRepository: exerciseRepository,
	}
}

func (h GetUserExercisesHandler) GetUserExercises(ctx *appcontext.AppContext, req *exercisepb.GetUserExercisesRequest) (*exercisepb.GetUserExercisesResponse, error) {
	ctx.Logger().Info("[hub] new get user exercises request", appcontext.Fields{"userID": req.GetUserId(), "level": req.GetLevel(), "lang": req.GetLang()})

	ctx.Logger().Text("new user exercise filter")
	filter, err := domain.NewUserExerciseFilter(req.GetUserId(), req.GetLevel(), req.GetLang())
	if err != nil {
		ctx.Logger().Error("failed to create new user exercise filter", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("pick random exercises for user")
	exercises, err := h.exerciseRepository.PickRandomExercisesForUser(ctx, *filter)
	if err != nil {
		ctx.Logger().Error("failed to pick random exercises for user", err, appcontext.Fields{})
		return nil, err
	}

	result := dto.ConvertUserExercisesFromDomainToGrpc(exercises)
	ctx.Logger().Text("done get user exercises request")

	return &exercisepb.GetUserExercisesResponse{
		Exercises: result,
	}, nil
}
