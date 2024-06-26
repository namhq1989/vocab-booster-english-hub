package hub

import (
	"github.com/namhq1989/vocab-booster-english-hub/core/appcontext"
	apperrors "github.com/namhq1989/vocab-booster-english-hub/core/error"
	"github.com/namhq1989/vocab-booster-english-hub/internal/genproto/vocabularypb"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/domain"
)

type CreateUserVocabularyCollectionHandler struct {
	userVocabularyCollectionRepository domain.UserVocabularyCollectionRepository
}

func NewCreateUserVocabularyCollectionHandler(userVocabularyCollectionRepository domain.UserVocabularyCollectionRepository) CreateUserVocabularyCollectionHandler {
	return CreateUserVocabularyCollectionHandler{
		userVocabularyCollectionRepository: userVocabularyCollectionRepository,
	}
}

func (h CreateUserVocabularyCollectionHandler) CreateUserVocabularyCollection(ctx *appcontext.AppContext, req *vocabularypb.CreateUserVocabularyCollectionRequest) (*vocabularypb.CreateUserVocabularyCollectionResponse, error) {
	ctx.Logger().Info("[hub] new create user vocabulary collection request", appcontext.Fields{"userID": req.GetUserId(), "name": req.GetName()})

	ctx.Logger().Text("create collection model")
	collection, err := domain.NewUserVocabularyCollection(req.GetUserId(), req.GetName())
	if err != nil {
		ctx.Logger().Error("failed to create collection model", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("count total collections of user")
	total, err := h.userVocabularyCollectionRepository.CountTotalUserVocabularyCollectionsByUserID(ctx, req.GetUserId())
	if err != nil {
		ctx.Logger().Error("failed to count total collections of user", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Info("check the cap of collections", appcontext.Fields{"totalCreated": total})
	if !collection.CanCreate(int(total)) {
		ctx.Logger().ErrorText("the cap of collections is reached")
		return nil, apperrors.UserVocabularyCollection.MaxCollectionsReached
	}

	ctx.Logger().Text("persist collection in db")
	if err = h.userVocabularyCollectionRepository.CreateUserVocabularyCollection(ctx, *collection); err != nil {
		ctx.Logger().Error("failed to persist collection in db", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("done create user vocabulary collection request")
	return &vocabularypb.CreateUserVocabularyCollectionResponse{
		Id: collection.ID,
	}, nil
}
