package hub

import (
	"github.com/namhq1989/vocab-booster-english-hub/internal/genproto/vocabularypb"
	apperrors "github.com/namhq1989/vocab-booster-english-hub/internal/utils/error"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/domain"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type UpdateCollectionHandler struct {
	CollectionRepository domain.CollectionRepository
}

func NewUpdateCollectionHandler(CollectionRepository domain.CollectionRepository) UpdateCollectionHandler {
	return UpdateCollectionHandler{
		CollectionRepository: CollectionRepository,
	}
}

func (h UpdateCollectionHandler) UpdateCollection(ctx *appcontext.AppContext, req *vocabularypb.UpdateCollectionRequest) (*vocabularypb.UpdateCollectionResponse, error) {
	ctx.Logger().Info("[hub] new update collection request", appcontext.Fields{"collectionID": req.GetCollectionId(), "name": req.GetName(), "description": req.GetDescription()})

	ctx.Logger().Text("find collection in db")
	collection, err := h.CollectionRepository.FindCollectionByID(ctx, req.GetCollectionId())
	if err != nil {
		ctx.Logger().Error("failed to find collection in db", err, appcontext.Fields{})
		return nil, err
	}
	if collection == nil {
		ctx.Logger().Error("collection not found in db", err, appcontext.Fields{})
		return nil, apperrors.Collection.CollectionNotFound
	}

	ctx.Logger().Text("update collection data")
	if err = collection.SetName(req.GetName()); err != nil {
		ctx.Logger().Error("failed to update collection name", err, appcontext.Fields{})
		return nil, err
	}
	if err = collection.SetDescription(req.GetDescription()); err != nil {
		ctx.Logger().Error("failed to update collection description", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("update collection in db")
	if err = h.CollectionRepository.UpdateCollection(ctx, *collection); err != nil {
		ctx.Logger().Error("failed to persist collection in db", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("done update collection request")
	return &vocabularypb.UpdateCollectionResponse{
		Id: collection.ID,
	}, nil
}
