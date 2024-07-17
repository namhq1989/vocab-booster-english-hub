package hub

import (
	"github.com/namhq1989/vocab-booster-english-hub/internal/genproto/exercisepb"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/exercise/domain"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/exercise/dto"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type GetExerciseCollectionsHandler struct {
	service domain.Service
}

func NewGetExerciseCollectionsHandler(
	service domain.Service,
) GetExerciseCollectionsHandler {
	return GetExerciseCollectionsHandler{
		service: service,
	}
}

func (h GetExerciseCollectionsHandler) GetExerciseCollections(ctx *appcontext.AppContext, req *exercisepb.GetExerciseCollectionsRequest) (*exercisepb.GetExerciseCollectionsResponse, error) {
	ctx.Logger().Info("[hub] new get exercise collections request", appcontext.Fields{"userID": req.GetUserId(), "lang": req.GetLang()})

	ctx.Logger().Text("find in service")
	collections, err := h.service.FindExerciseCollections(ctx)
	if err != nil {
		ctx.Logger().Error("failed to find in service", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("convert to response data")
	result := dto.ConvertExerciseCollectionsFromDomainToGrpc(collections, req.GetLang())

	ctx.Logger().Text("done get exercise collections request")
	return &exercisepb.GetExerciseCollectionsResponse{Collections: result}, nil
}
