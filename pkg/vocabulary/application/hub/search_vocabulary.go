package hub

import (
	"github.com/namhq1989/vocab-booster-english-hub/internal/genproto/vocabularypb"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/domain"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/dto"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type SearchVocabularyHandler struct {
	userBookmarkedVocabularyRepository domain.UserBookmarkedVocabularyRepository
	service                            domain.Service
}

func NewSearchVocabularyHandler(
	userBookmarkedVocabularyRepository domain.UserBookmarkedVocabularyRepository,
	service domain.Service,
) SearchVocabularyHandler {
	return SearchVocabularyHandler{
		userBookmarkedVocabularyRepository: userBookmarkedVocabularyRepository,
		service:                            service,
	}
}

func (h SearchVocabularyHandler) SearchVocabulary(ctx *appcontext.AppContext, req *vocabularypb.SearchVocabularyRequest) (*vocabularypb.SearchVocabularyResponse, error) {
	ctx.Logger().Info("new search vocabulary request", appcontext.Fields{"term": req.GetTerm()})
	var result = &vocabularypb.SearchVocabularyResponse{
		Found:       false,
		Suggestions: make([]string, 0),
		Vocabulary:  nil,
	}

	ctx.Logger().Text("find vocabulary")
	vocabulary, err := h.service.FindVocabulary(ctx, req.GetTerm())
	if err != nil {
		ctx.Logger().Error("failed to find vocabulary", err, appcontext.Fields{})
		return nil, err
	}
	if vocabulary == nil {
		ctx.Logger().ErrorText("vocabulary not found")
		return result, nil
	}
	ctx.Logger().Text("vocabulary found, find examples")
	examples, examplesErr := h.service.FindVocabularyExamples(ctx, vocabulary.ID)
	if examplesErr != nil {
		ctx.Logger().Error("failed to find vocabulary examples", examplesErr, appcontext.Fields{"vocabularyID": vocabulary.ID})
		return nil, examplesErr
	}

	ctx.Logger().Text("check bookmarked")
	ubv, bookmarkedErr := h.userBookmarkedVocabularyRepository.FindBookmarkedVocabulary(ctx, req.GetPerformerId(), vocabulary.ID)
	if bookmarkedErr != nil {
		ctx.Logger().Error("failed to check bookmarked vocabulary", bookmarkedErr, appcontext.Fields{"vocabularyID": vocabulary.ID})
	}
	isBookmarked := ubv != nil

	result.Found = true
	result.Vocabulary = dto.ConvertVocabularyFromDomainToGrpc(*vocabulary, examples, isBookmarked)

	ctx.Logger().Text("done search vocabulary request")
	return result, nil
}
