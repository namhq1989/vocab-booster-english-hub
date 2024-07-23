package hub

import (
	"github.com/namhq1989/vocab-booster-english-hub/internal/genproto/exercisepb"
	apperrors "github.com/namhq1989/vocab-booster-english-hub/internal/utils/error"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/exercise/domain"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/exercise/dto"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type GetUserExercisesHandler struct {
	exerciseRepository domain.ExerciseRepository
	service            domain.Service
}

func NewGetUserExercisesHandler(exerciseRepository domain.ExerciseRepository, service domain.Service) GetUserExercisesHandler {
	return GetUserExercisesHandler{
		exerciseRepository: exerciseRepository,
		service:            service,
	}
}

func (h GetUserExercisesHandler) GetUserExercises(ctx *appcontext.AppContext, req *exercisepb.GetUserExercisesRequest) (*exercisepb.GetUserExercisesResponse, error) {
	ctx.Logger().Info("[hub] new get user exercises request", appcontext.Fields{
		"userID": req.GetUserId(), "collectionID": req.GetCollectionId(), "lang": req.GetLang(),
	})

	ctx.Logger().Text("get collection by id")
	collection, err := h.service.FindExerciseCollectionByID(ctx, req.GetCollectionId())
	if err != nil {
		ctx.Logger().Error("failed to get collection by id", err, appcontext.Fields{})
		return nil, err
	}
	if collection == nil {
		ctx.Logger().Text("collection not found")
		return nil, apperrors.Collection.CollectionNotFound
	}

	ctx.Logger().Text("new user exercise filter")
	filter, err := domain.NewUserExerciseFilter(req.GetUserId(), collection.Criteria, req.GetLang(), "")
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
