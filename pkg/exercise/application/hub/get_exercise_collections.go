package hub

import (
	"github.com/namhq1989/vocab-booster-english-hub/internal/genproto/exercisepb"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/exercise/domain"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/exercise/dto"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type GetExerciseCollectionsHandler struct {
	exerciseCollectionRepository domain.ExerciseCollectionRepository
	cachingRepository            domain.CachingRepository
}

func NewGetExerciseCollectionsHandler(
	exerciseCollectionRepository domain.ExerciseCollectionRepository,
	cachingRepository domain.CachingRepository,
) GetExerciseCollectionsHandler {
	return GetExerciseCollectionsHandler{
		exerciseCollectionRepository: exerciseCollectionRepository,
		cachingRepository:            cachingRepository,
	}
}

func (h GetExerciseCollectionsHandler) GetExerciseCollections(ctx *appcontext.AppContext, req *exercisepb.GetExerciseCollectionsRequest) (*exercisepb.GetExerciseCollectionsResponse, error) {
	ctx.Logger().Info("[hub] new get exercise collections request", appcontext.Fields{"userID": req.GetUserId(), "lang": req.GetLang()})

	ctx.Logger().Text("find in caching")
	cachingCollections, err := h.cachingRepository.GetUserExerciseCollections(ctx, req.GetUserId())
	if cachingCollections != nil {
		ctx.Logger().Info("got data in caching", appcontext.Fields{"numOfCollections": len(*cachingCollections)})

		ctx.Logger().Text("convert to response data")
		result := dto.ConvertUserExerciseCollectionsFromDomainToGrpc(*cachingCollections, req.GetLang())

		ctx.Logger().Text("done get exercise collections request")
		return &exercisepb.GetExerciseCollectionsResponse{Collections: result}, nil
	} else if err != nil {
		ctx.Logger().Error("failed to find in caching", err, appcontext.Fields{})
	}

	ctx.Logger().Text("find in db")
	collections, err := h.exerciseCollectionRepository.FindUserExerciseCollections(ctx, req.GetUserId())
	if err != nil {
		ctx.Logger().Error("failed to find in service", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("convert to response data")
	result := dto.ConvertUserExerciseCollectionsFromDomainToGrpc(collections, req.GetLang())

	ctx.Logger().Text("done get exercise collections request")
	return &exercisepb.GetExerciseCollectionsResponse{Collections: result}, nil
}
