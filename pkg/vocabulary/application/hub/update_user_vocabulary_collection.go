package hub

import (
	"github.com/namhq1989/vocab-booster-english-hub/core/appcontext"
	apperrors "github.com/namhq1989/vocab-booster-english-hub/core/error"
	"github.com/namhq1989/vocab-booster-english-hub/internal/genproto/vocabularypb"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/domain"
)

type UpdateUserVocabularyCollectionHandler struct {
	userVocabularyCollectionRepository domain.UserVocabularyCollectionRepository
}

func NewUpdateUserVocabularyCollectionHandler(userVocabularyCollectionRepository domain.UserVocabularyCollectionRepository) UpdateUserVocabularyCollectionHandler {
	return UpdateUserVocabularyCollectionHandler{
		userVocabularyCollectionRepository: userVocabularyCollectionRepository,
	}
}

func (h UpdateUserVocabularyCollectionHandler) UpdateUserVocabularyCollection(ctx *appcontext.AppContext, req *vocabularypb.UpdateUserVocabularyCollectionRequest) (*vocabularypb.UpdateUserVocabularyCollectionResponse, error) {
	ctx.Logger().Info("[hub] new update user vocabulary collection request", appcontext.Fields{"collectionID": req.GetCollectionId(), "name": req.GetName()})

	ctx.Logger().Text("find collection in db")
	collection, err := h.userVocabularyCollectionRepository.FindUserVocabularyCollectionByID(ctx, req.GetCollectionId())
	if err != nil {
		ctx.Logger().Error("failed to find collection in db", err, appcontext.Fields{})
		return nil, err
	}
	if collection == nil {
		ctx.Logger().Error("collection not found in db", err, appcontext.Fields{})
		return nil, apperrors.UserVocabularyCollection.CollectionNotFound
	}

	ctx.Logger().Text("update collection data")
	if err = collection.SetName(req.GetName()); err != nil {
		ctx.Logger().Error("failed to update collection name", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("update collection in db")
	if err = h.userVocabularyCollectionRepository.UpdateUserVocabularyCollection(ctx, *collection); err != nil {
		ctx.Logger().Error("failed to persist collection in db", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("done update user vocabulary collection request")
	return &vocabularypb.UpdateUserVocabularyCollectionResponse{
		Id: collection.ID,
	}, nil
}
