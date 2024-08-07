package hub

import (
	"github.com/namhq1989/vocab-booster-english-hub/internal/genproto/vocabularypb"
	apperrors "github.com/namhq1989/vocab-booster-english-hub/internal/utils/error"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/domain"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/dto"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type GetCommunitySentenceHandler struct {
	communitySentenceRepository domain.CommunitySentenceRepository
}

func NewGetCommunitySentenceHandler(
	communitySentenceRepository domain.CommunitySentenceRepository,
) GetCommunitySentenceHandler {
	return GetCommunitySentenceHandler{
		communitySentenceRepository: communitySentenceRepository,
	}
}

func (h GetCommunitySentenceHandler) GetCommunitySentence(ctx *appcontext.AppContext, req *vocabularypb.GetCommunitySentenceRequest) (*vocabularypb.GetCommunitySentenceResponse, error) {
	ctx.Logger().Info("[hub] new get community sentence request", appcontext.Fields{"userID": req.GetUserId(), "sentenceID": req.GetSentenceId()})

	ctx.Logger().Text("find sentence in db")
	sentence, err := h.communitySentenceRepository.FindCommunitySentenceWithUserID(ctx, req.GetSentenceId(), req.GetUserId())
	if err != nil {
		ctx.Logger().Error("failed to query sentence in db", err, appcontext.Fields{})
		return nil, err
	}
	if sentence == nil {
		ctx.Logger().ErrorText("sentence not found")
		return nil, apperrors.Vocabulary.InvalidSentence
	}

	ctx.Logger().Text("convert to grpc data")
	result := &vocabularypb.GetCommunitySentenceResponse{
		Sentence: nil,
	}
	result.Sentence = dto.ConvertCommunitySentenceFromDomainToGrpc(*sentence)

	ctx.Logger().Text("done get community sentence request")
	return result, nil
}
