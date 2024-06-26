package hub

import (
	"github.com/namhq1989/vocab-booster-english-hub/core/appcontext"
	apperrors "github.com/namhq1989/vocab-booster-english-hub/core/error"
	"github.com/namhq1989/vocab-booster-english-hub/internal/genproto/vocabularypb"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/domain"
)

type RemoveVocabularyFromCollectionHandler struct {
	vocabularyRepository              domain.VocabularyRepository
	collectionRepository              domain.CollectionRepository
	collectionAndVocabularyRepository domain.CollectionAndVocabularyRepository
}

func NewRemoveVocabularyFromCollectionHandler(
	vocabularyRepository domain.VocabularyRepository,
	collectionRepository domain.CollectionRepository,
	collectionAndVocabularyRepository domain.CollectionAndVocabularyRepository,
) RemoveVocabularyFromCollectionHandler {
	return RemoveVocabularyFromCollectionHandler{
		vocabularyRepository:              vocabularyRepository,
		collectionRepository:              collectionRepository,
		collectionAndVocabularyRepository: collectionAndVocabularyRepository,
	}
}

func (h RemoveVocabularyFromCollectionHandler) RemoveVocabularyFromCollection(ctx *appcontext.AppContext, req *vocabularypb.RemoveVocabularyFromCollectionRequest) (*vocabularypb.RemoveVocabularyFromCollectionResponse, error) {
	ctx.Logger().Info("[hub] new remove vocabulary from collection request", appcontext.Fields{"userID": req.GetUserId(), "collectionID": req.GetCollectionId(), "vocabulary": req.GetVocabulary()})

	ctx.Logger().Text("find vocabulary in db")
	vocabulary, err := h.vocabularyRepository.FindVocabularyByTerm(ctx, req.GetVocabulary())
	if err != nil {
		ctx.Logger().Error("failed to find vocabulary in db", err, appcontext.Fields{})
		return nil, err
	}
	if vocabulary == nil {
		ctx.Logger().ErrorText("vocabulary not found in db")
		return nil, apperrors.Vocabulary.VocabularyNotFound
	}

	ctx.Logger().Text("find collection in db")
	collection, err := h.collectionRepository.FindCollectionByID(ctx, req.GetCollectionId())
	if err != nil {
		ctx.Logger().Error("failed to find collection in db", err, appcontext.Fields{})
		return nil, err
	}
	if collection == nil {
		ctx.Logger().Error("collection not found in db", err, appcontext.Fields{})
		return nil, apperrors.Collection.CollectionNotFound
	}

	ctx.Logger().Text("check collection's owner")
	if !collection.IsOwner(req.GetUserId()) {
		ctx.Logger().ErrorText("the collection is not owned by the user, respond")
		return nil, apperrors.Common.Forbidden
	}

	ctx.Logger().Text("check collection vocabulary existence")
	cav, err := h.collectionAndVocabularyRepository.FindCollectionAndVocabulary(ctx, collection.ID, vocabulary.ID)
	if err != nil {
		ctx.Logger().Error("failed to find collection vocabulary in db", err, appcontext.Fields{})
		return nil, err
	}
	if cav == nil {
		ctx.Logger().Text("vocabulary not found in collection, respond")

		ctx.Logger().Text("done remove vocabulary from collection request")
		return &vocabularypb.RemoveVocabularyFromCollectionResponse{}, nil
	}

	ctx.Logger().Text("vocabulary found in collection, delete it")
	if err = h.collectionAndVocabularyRepository.DeleteCollectionAndVocabulary(ctx, *cav); err != nil {
		ctx.Logger().Error("failed to delete collection vocabulary in db", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("update collection in db")
	collection.DecreaseNumOfVocabulary()
	if err = h.collectionRepository.UpdateCollection(ctx, *collection); err != nil {
		ctx.Logger().Error("failed to persist collection in db", err, appcontext.Fields{})
	}

	ctx.Logger().Text("done remove vocabulary from collection request")
	return &vocabularypb.RemoveVocabularyFromCollectionResponse{}, nil
}
