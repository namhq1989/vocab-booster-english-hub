package hub

import (
	"github.com/namhq1989/vocab-booster-english-hub/internal/genproto/vocabularypb"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/domain"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/dto"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type GetVocabularyCommunitySentencesHandler struct {
	communitySentenceRepository domain.CommunitySentenceRepository
}

func NewGetVocabularyCommunitySentencesHandler(
	communitySentenceRepository domain.CommunitySentenceRepository,
) GetVocabularyCommunitySentencesHandler {
	return GetVocabularyCommunitySentencesHandler{
		communitySentenceRepository: communitySentenceRepository,
	}
}

func (h GetVocabularyCommunitySentencesHandler) GetVocabularyCommunitySentences(ctx *appcontext.AppContext, req *vocabularypb.GetVocabularyCommunitySentencesRequest) (*vocabularypb.GetVocabularyCommunitySentencesResponse, error) {
	ctx.Logger().Info("[hub] new get vocabulary community sentences request", appcontext.Fields{"userID": req.GetUserId(), "vocabularyID": req.GetVocabularyId(), "lang": req.GetLang(), "pageToken": req.GetPageToken()})

	ctx.Logger().Text("new vocabulary community sentences filter")
	filter, err := domain.NewVocabularyCommunitySentenceFilter(req.GetUserId(), req.GetVocabularyId(), req.GetLang(), req.GetPageToken())
	if err != nil {
		ctx.Logger().Error("failed to create new vocabulary community sentences filter", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("find sentences in db")
	sentences, err := h.communitySentenceRepository.FindVocabularyCommunitySentences(ctx, *filter)
	if err != nil {
		ctx.Logger().Error("failed to query sentences in db", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("convert to grpc data")
	result := dto.ConvertCommunitySentencesFromDomainToGrpc(sentences)

	ctx.Logger().Text("done get vocabulary community sentences request")
	return &vocabularypb.GetVocabularyCommunitySentencesResponse{
		Sentences: result,
	}, nil
}
