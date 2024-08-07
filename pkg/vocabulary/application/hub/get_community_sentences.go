package hub

import (
	"github.com/namhq1989/vocab-booster-english-hub/internal/genproto/vocabularypb"
	"github.com/namhq1989/vocab-booster-english-hub/internal/utils/pagetoken"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/domain"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/dto"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type GetCommunitySentencesHandler struct {
	communitySentenceRepository domain.CommunitySentenceRepository
}

func NewGetCommunitySentencesHandler(
	communitySentenceRepository domain.CommunitySentenceRepository,
) GetCommunitySentencesHandler {
	return GetCommunitySentencesHandler{
		communitySentenceRepository: communitySentenceRepository,
	}
}

func (h GetCommunitySentencesHandler) GetCommunitySentences(ctx *appcontext.AppContext, req *vocabularypb.GetCommunitySentencesRequest) (*vocabularypb.GetCommunitySentencesResponse, error) {
	ctx.Logger().Info("[hub] new get vocabulary community sentences request", appcontext.Fields{"userID": req.GetUserId(), "vocabularyID": req.GetVocabularyId(), "lang": req.GetLang(), "pageToken": req.GetPageToken()})

	ctx.Logger().Text("new vocabulary community sentences filter")
	filter, err := domain.NewVocabularyCommunitySentenceFilter(req.GetUserId(), req.GetVocabularyId(), req.GetLang(), req.GetPageToken())
	if err != nil {
		ctx.Logger().Error("failed to create new vocabulary community sentences filter", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("find sentences in db")
	sentences, err := h.communitySentenceRepository.FindCommunitySentences(ctx, *filter)
	if err != nil {
		ctx.Logger().Error("failed to query sentences in db", err, appcontext.Fields{})
		return nil, err
	}

	var (
		totalSentences = len(sentences)
		result         = &vocabularypb.GetCommunitySentencesResponse{
			Sentences:     make([]*vocabularypb.CommunitySentenceBrief, 0),
			NextPageToken: "",
		}
	)

	if totalSentences == 0 {
		ctx.Logger().Text("no sentences found, respond")
		return result, nil
	}

	ctx.Logger().Text("convert to grpc data")
	for _, sentence := range sentences {
		result.Sentences = append(result.Sentences, dto.ConvertCommunitySentenceBriefFromDomainToGrpc(sentence))
	}
	result.NextPageToken = pagetoken.NewWithTimestamp(sentences[totalSentences-1].CreatedAt)

	ctx.Logger().Text("done get vocabulary community sentences request")
	return result, nil
}
