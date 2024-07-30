package hub

import (
	"github.com/namhq1989/vocab-booster-english-hub/internal/genproto/vocabularypb"
	"github.com/namhq1989/vocab-booster-english-hub/internal/utils/pagetoken"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/domain"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/dto"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type GetUserBookmarkedVocabulariesHandler struct {
	userBookmarkedVocabularyRepository domain.UserBookmarkedVocabularyRepository
}

func NewGetUserBookmarkedVocabulariesHandler(userBookmarkedVocabularyRepository domain.UserBookmarkedVocabularyRepository) GetUserBookmarkedVocabulariesHandler {
	return GetUserBookmarkedVocabulariesHandler{
		userBookmarkedVocabularyRepository: userBookmarkedVocabularyRepository,
	}
}

func (h GetUserBookmarkedVocabulariesHandler) GetUserBookmarkedVocabularies(ctx *appcontext.AppContext, req *vocabularypb.GetUserBookmarkedVocabulariesRequest) (*vocabularypb.GetUserBookmarkedVocabulariesResponse, error) {
	ctx.Logger().Info("[hub] new get user bookmarked vocabularies request", appcontext.Fields{"userID": req.GetUserId(), "pageToken": req.GetPageToken()})

	ctx.Logger().Text("new user bookmarked vocabularies filter")
	filter, err := domain.NewUserBookmarkedVocabularyFilter(req.GetPageToken())
	if err != nil {
		ctx.Logger().Error("failed to create new user bookmarked vocabularies filter", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("find bookmarked vocabularies in db")
	eubvs, err := h.userBookmarkedVocabularyRepository.FindBookmarkedVocabulariesByUserID(ctx, req.GetUserId(), *filter)
	if err != nil {
		ctx.Logger().Error("failed to query bookmarked vocabularies in db", err, appcontext.Fields{})
		return nil, err
	}

	totalEubvs := len(eubvs)
	if totalEubvs == 0 {
		ctx.Logger().Text("no bookmarked vocabularies found, respond")
		return &vocabularypb.GetUserBookmarkedVocabulariesResponse{
			Vocabularies:  make([]*vocabularypb.VocabularyBrief, 0),
			NextPageToken: "",
		}, nil
	}

	ctx.Logger().Text("convert to grpc data")
	var result = &vocabularypb.GetUserBookmarkedVocabulariesResponse{
		Vocabularies: make([]*vocabularypb.VocabularyBrief, 0),
	}

	for _, ubv := range eubvs {
		result.Vocabularies = append(result.Vocabularies, dto.ConvertVocabularyBriefFromDomainToGrpc(ubv.Vocabulary))
	}

	ctx.Logger().Text("generate page token")
	result.NextPageToken = pagetoken.NewWithTimestamp(eubvs[totalEubvs-1].BookmarkedAt)

	ctx.Logger().Text("done get user bookmarked vocabularies request")
	return result, nil
}
