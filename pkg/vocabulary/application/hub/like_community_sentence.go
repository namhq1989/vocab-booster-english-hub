package hub

import (
	"github.com/namhq1989/vocab-booster-english-hub/internal/genproto/vocabularypb"
	apperrors "github.com/namhq1989/vocab-booster-english-hub/internal/utils/error"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/domain"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type LikeCommunitySentenceHandler struct {
	communitySentenceRepository     domain.CommunitySentenceRepository
	communitySentenceLikeRepository domain.CommunitySentenceLikeRepository
}

func NewLikeCommunitySentenceHandler(
	communitySentenceRepository domain.CommunitySentenceRepository,
	communitySentenceLikeRepository domain.CommunitySentenceLikeRepository,
) LikeCommunitySentenceHandler {
	return LikeCommunitySentenceHandler{
		communitySentenceRepository:     communitySentenceRepository,
		communitySentenceLikeRepository: communitySentenceLikeRepository,
	}
}

func (h LikeCommunitySentenceHandler) LikeCommunitySentence(ctx *appcontext.AppContext, req *vocabularypb.LikeCommunitySentenceRequest) (*vocabularypb.LikeCommunitySentenceResponse, error) {
	ctx.Logger().Info("[hub] new like community sentence request", appcontext.Fields{"userID": req.GetUserId(), "sentenceID": req.GetSentenceId()})

	ctx.Logger().Text("find sentence in db")
	sentence, err := h.communitySentenceRepository.FindCommunitySentenceByID(ctx, req.GetSentenceId())
	if err != nil {
		ctx.Logger().Error("failed to find sentence in db", err, appcontext.Fields{})
		return nil, err
	}
	if sentence == nil {
		ctx.Logger().ErrorText("sentence not found")
		return nil, apperrors.Vocabulary.InvalidSentence
	}

	ctx.Logger().Text("find like document in db")
	like, err := h.communitySentenceLikeRepository.FindCommunitySentenceLike(ctx, req.GetUserId(), req.GetSentenceId())
	if err != nil {
		ctx.Logger().Error("failed to find like document in db", err, appcontext.Fields{})
		return nil, err
	}

	var isLiked = false

	if like == nil {
		ctx.Logger().Text("user not liked sentence yet, create new like document")
		like, err = domain.NewCommunitySentenceLike(req.GetUserId(), req.GetSentenceId())
		if err != nil {
			ctx.Logger().Error("failed to create new like document", err, appcontext.Fields{})
			return nil, err
		}

		ctx.Logger().Text("persist like document in db")
		err = h.communitySentenceLikeRepository.CreateCommunitySentenceLike(ctx, *like)
		if err != nil {
			ctx.Logger().Error("failed to persist like document in db", err, appcontext.Fields{})
			return nil, err
		}

		isLiked = true
	} else {
		ctx.Logger().Text("user already liked sentence, delete like document")
		if err = h.communitySentenceLikeRepository.DeleteCommunitySentenceLike(ctx, *like); err != nil {
			ctx.Logger().Error("failed to delete like document in db", err, appcontext.Fields{})
			return nil, err
		}
	}

	ctx.Logger().Text("update sentence like stats")
	if isLiked {
		sentence.IncreaseStatsLike()
	} else {
		sentence.DecreaseStatsLike()
	}
	if err = h.communitySentenceRepository.UpdateCommunitySentence(ctx, *sentence); err != nil {
		ctx.Logger().Error("failed to update sentence in db", err, appcontext.Fields{})
	}

	ctx.Logger().Text("done like community sentence request")
	return &vocabularypb.LikeCommunitySentenceResponse{IsLiked: isLiked}, nil
}
