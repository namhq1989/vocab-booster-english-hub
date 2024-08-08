package hub

import (
	"github.com/namhq1989/vocab-booster-english-hub/internal/genproto/vocabularypb"
	"github.com/namhq1989/vocab-booster-english-hub/internal/utils/pagetoken"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/domain"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/dto"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type GetUserCommunitySentenceDraftsHandler struct {
	communitySentenceDraftRepository domain.CommunitySentenceDraftRepository
}

func NewGetUserCommunitySentenceDraftsHandler(
	communitySentenceDraftRepository domain.CommunitySentenceDraftRepository,
) GetUserCommunitySentenceDraftsHandler {
	return GetUserCommunitySentenceDraftsHandler{
		communitySentenceDraftRepository: communitySentenceDraftRepository,
	}
}

func (h GetUserCommunitySentenceDraftsHandler) GetUserCommunitySentenceDrafts(ctx *appcontext.AppContext, req *vocabularypb.GetUserCommunitySentenceDraftsRequest) (*vocabularypb.GetUserCommunitySentenceDraftsResponse, error) {
	ctx.Logger().Info("[hub] new get user draft community sentences request", appcontext.Fields{
		"userID": req.GetUserId(), "vocabularyID": req.GetVocabularyId(),
		"pageToken": req.GetPageToken(),
	})

	ctx.Logger().Text("create filter")
	filter, err := domain.NewCommunitySentenceDraftFilter(req.GetUserId(), req.GetVocabularyId(), req.GetPageToken())
	if err != nil {
		ctx.Logger().Error("failed to create filter", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("find draft sentences in db")
	sentences, err := h.communitySentenceDraftRepository.FindUserCommunitySentenceDrafts(ctx, *filter)
	if err != nil {
		ctx.Logger().Error("failed to query sentences in db", err, appcontext.Fields{})
		return nil, err
	}

	var (
		totalSentences = len(sentences)
		result         = &vocabularypb.GetUserCommunitySentenceDraftsResponse{
			Sentences:     make([]*vocabularypb.CommunitySentenceDraft, 0),
			NextPageToken: "",
		}
	)

	if totalSentences == 0 {
		ctx.Logger().Text("no sentences found, respond")
		return result, nil
	}

	ctx.Logger().Text("convert response data")
	for _, sentence := range sentences {
		result.Sentences = append(result.Sentences, dto.ConvertCommunitySentenceDraftFromDomainToGrpc(sentence))
	}
	result.NextPageToken = pagetoken.NewWithTimestamp(sentences[totalSentences-1].CreatedAt)

	ctx.Logger().Text("done get user draft community sentences request")
	return result, nil
}
