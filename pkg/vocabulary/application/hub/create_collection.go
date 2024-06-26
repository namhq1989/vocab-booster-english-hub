package hub

import (
	"github.com/namhq1989/vocab-booster-english-hub/core/appcontext"
	apperrors "github.com/namhq1989/vocab-booster-english-hub/core/error"
	"github.com/namhq1989/vocab-booster-english-hub/internal/genproto/vocabularypb"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/domain"
)

type CreateCollectionHandler struct {
	CollectionRepository domain.CollectionRepository
}

func NewCreateCollectionHandler(CollectionRepository domain.CollectionRepository) CreateCollectionHandler {
	return CreateCollectionHandler{
		CollectionRepository: CollectionRepository,
	}
}

func (h CreateCollectionHandler) CreateCollection(ctx *appcontext.AppContext, req *vocabularypb.CreateCollectionRequest) (*vocabularypb.CreateCollectionResponse, error) {
	ctx.Logger().Info("[hub] new create user collection request", appcontext.Fields{"userID": req.GetUserId(), "name": req.GetName(), "description": req.GetDescription()})

	ctx.Logger().Text("create collection model")
	collection, err := domain.NewCollection(req.GetUserId(), req.GetName(), req.GetDescription())
	if err != nil {
		ctx.Logger().Error("failed to create collection model", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("count total collections of user")
	total, err := h.CollectionRepository.CountTotalCollectionsByUserID(ctx, req.GetUserId())
	if err != nil {
		ctx.Logger().Error("failed to count total collections of user", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Info("check the cap of collections", appcontext.Fields{"totalCreated": total})
	if !collection.CanCreate(int(total)) {
		ctx.Logger().ErrorText("the cap of collections is reached")
		return nil, apperrors.Collection.MaxCollectionsReached
	}

	ctx.Logger().Text("persist collection in db")
	if err = h.CollectionRepository.CreateCollection(ctx, *collection); err != nil {
		ctx.Logger().Error("failed to persist collection in db", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("done create user collection request")
	return &vocabularypb.CreateCollectionResponse{
		Id: collection.ID,
	}, nil
}
