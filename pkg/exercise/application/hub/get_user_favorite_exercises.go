package hub

import (
	"github.com/namhq1989/vocab-booster-english-hub/internal/genproto/exercisepb"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/exercise/domain"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/exercise/dto"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type GetUserFavoriteExercisesHandler struct {
	userExerciseStatusRepository domain.UserExerciseStatusRepository
}

func NewGetUserFavoriteExercisesHandler(userExerciseStatusRepository domain.UserExerciseStatusRepository) GetUserFavoriteExercisesHandler {
	return GetUserFavoriteExercisesHandler{
		userExerciseStatusRepository: userExerciseStatusRepository,
	}
}

func (h GetUserFavoriteExercisesHandler) GetUserFavoriteExercises(ctx *appcontext.AppContext, req *exercisepb.GetUserFavoriteExercisesRequest) (*exercisepb.GetUserFavoriteExercisesResponse, error) {
	ctx.Logger().Info("[hub] new get user favorite exercises request", appcontext.Fields{"userID": req.GetUserId(), "lang": req.GetLang(), "pageToken": req.GetPageToken()})

	ctx.Logger().Text("new user favorite exercises filter")
	filter, err := domain.NewUserFavoriteExerciseFilter(req.GetUserId(), req.GetLang(), req.GetPageToken())
	if err != nil {
		ctx.Logger().Error("failed to create new user favorite exercises filter", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("find exercises in db")
	exercises, err := h.userExerciseStatusRepository.FindUserFavoriteExercises(ctx, *filter)
	if err != nil {
		ctx.Logger().Error("failed to query exercises in db", err, appcontext.Fields{})
		return nil, err
	}

	result := dto.ConvertUserExercisesFromDomainToGrpc(exercises, req.GetLang())
	ctx.Logger().Text("done get user favorite exercises request")

	return &exercisepb.GetUserFavoriteExercisesResponse{
		Exercises: result,
	}, nil
}
