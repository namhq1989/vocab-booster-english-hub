package hub

import (
	"github.com/namhq1989/vocab-booster-english-hub/internal/genproto/vocabularypb"
	apperrors "github.com/namhq1989/vocab-booster-english-hub/internal/utils/error"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/domain"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type BookmarkVocabularyHandler struct {
	vocabularyRepository               domain.VocabularyRepository
	userBookmarkedVocabularyRepository domain.UserBookmarkedVocabularyRepository
}

func NewBookmarkVocabularyHandler(vocabularyRepository domain.VocabularyRepository, userBookmarkedVocabularyRepository domain.UserBookmarkedVocabularyRepository) BookmarkVocabularyHandler {
	return BookmarkVocabularyHandler{
		vocabularyRepository:               vocabularyRepository,
		userBookmarkedVocabularyRepository: userBookmarkedVocabularyRepository,
	}
}

func (h BookmarkVocabularyHandler) BookmarkVocabulary(ctx *appcontext.AppContext, req *vocabularypb.BookmarkVocabularyRequest) (*vocabularypb.BookmarkVocabularyResponse, error) {
	ctx.Logger().Info("[hub] new bookmark vocabulary request", appcontext.Fields{"userID": req.GetUserId(), "vocabularyID": req.GetVocabularyId()})

	ctx.Logger().Text("find vocabulary in db")
	vocabulary, err := h.vocabularyRepository.FindVocabularyByID(ctx, req.GetVocabularyId())
	if err != nil {
		ctx.Logger().Error("failed to find vocabulary in db", err, appcontext.Fields{})
		return nil, err
	}
	if vocabulary == nil {
		ctx.Logger().ErrorText("vocabulary not found")
		return nil, apperrors.Vocabulary.VocabularyNotFound
	}

	ctx.Logger().Text("find bookmarked document in db")
	ubv, err := h.userBookmarkedVocabularyRepository.FindBookmarkedVocabulary(ctx, req.GetUserId(), req.GetVocabularyId())
	if err != nil {
		ctx.Logger().Error("failed to find bookmarked document in db", err, appcontext.Fields{})
		return nil, err
	}

	isBookmarked := false

	if ubv == nil {
		ctx.Logger().Text("user not bookmarked vocabulary yet, create new bookmarked document")
		ubv, err = domain.NewUserBookmarkedVocabulary(req.GetUserId(), req.GetVocabularyId())
		if err != nil {
			ctx.Logger().Error("failed to create new bookmarked document", err, appcontext.Fields{})
			return nil, err
		}

		ctx.Logger().Text("persist bookmarked document in db")
		err = h.userBookmarkedVocabularyRepository.CreateUserBookmarkedVocabulary(ctx, *ubv)
		if err != nil {
			ctx.Logger().Error("failed to persist bookmarked document in db", err, appcontext.Fields{})
			return nil, err
		}

		isBookmarked = true
	} else {
		ctx.Logger().Text("user already bookmarked sentence, delete bookmarked document")
		if err = h.userBookmarkedVocabularyRepository.DeleteUserBookmarkedVocabulary(ctx, *ubv); err != nil {
			ctx.Logger().Error("failed to delete bookmarked document in db", err, appcontext.Fields{})
			return nil, err
		}
	}

	ctx.Logger().Text("done bookmark vocabulary request")
	return &vocabularypb.BookmarkVocabularyResponse{IsBookmarked: isBookmarked}, nil
}
