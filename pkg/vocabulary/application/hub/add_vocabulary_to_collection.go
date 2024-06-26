package hub

import (
	"github.com/namhq1989/vocab-booster-english-hub/core/appcontext"
	apperrors "github.com/namhq1989/vocab-booster-english-hub/core/error"
	"github.com/namhq1989/vocab-booster-english-hub/internal/genproto/vocabularypb"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/domain"
)

type AddVocabularyToCollectionHandler struct {
	vocabularyRepository              domain.VocabularyRepository
	collectionRepository              domain.CollectionRepository
	collectionAndVocabularyRepository domain.CollectionAndVocabularyRepository
}

func NewAddVocabularyToCollectionHandler(
	vocabularyRepository domain.VocabularyRepository,
	collectionRepository domain.CollectionRepository,
	collectionAndVocabularyRepository domain.CollectionAndVocabularyRepository,
) AddVocabularyToCollectionHandler {
	return AddVocabularyToCollectionHandler{
		vocabularyRepository:              vocabularyRepository,
		collectionRepository:              collectionRepository,
		collectionAndVocabularyRepository: collectionAndVocabularyRepository,
	}
}

func (h AddVocabularyToCollectionHandler) AddVocabularyToCollection(ctx *appcontext.AppContext, req *vocabularypb.AddVocabularyToCollectionRequest) (*vocabularypb.AddVocabularyToCollectionResponse, error) {
	ctx.Logger().Info("[hub] new add vocabulary to collection request", appcontext.Fields{"userID": req.GetUserId(), "collectionID": req.GetCollectionId(), "vocabulary": req.GetVocabulary()})

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
	if cav != nil {
		ctx.Logger().Text("vocabulary already exists in collection, respond")

		ctx.Logger().Text("done add vocabulary to collection request")
		return &vocabularypb.AddVocabularyToCollectionResponse{}, nil
	}

	ctx.Logger().Text("vocabulary not found in collection, create new one")
	cav, err = domain.NewCollectionAndVocabulary(collection.ID, vocabulary.ID, req.GetVocabulary())
	if err != nil {
		ctx.Logger().Error("failed to create collection vocabulary model", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("persist collection vocabulary in db")
	if err = h.collectionAndVocabularyRepository.CreateCollectionAndVocabulary(ctx, *cav); err != nil {
		ctx.Logger().Error("failed to persist collection vocabulary in db", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("update collection in db")
	collection.IncreaseNumOfVocabulary()
	if err = h.collectionRepository.UpdateCollection(ctx, *collection); err != nil {
		ctx.Logger().Error("failed to persist collection in db", err, appcontext.Fields{})
	}

	ctx.Logger().Text("done add vocabulary to collection request")
	return &vocabularypb.AddVocabularyToCollectionResponse{}, nil
}
