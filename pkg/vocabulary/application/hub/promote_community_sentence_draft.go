package hub

import (
	"github.com/namhq1989/vocab-booster-english-hub/internal/genproto/vocabularypb"
	apperrors "github.com/namhq1989/vocab-booster-english-hub/internal/utils/error"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/domain"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type PromoteCommunitySentenceDraftHandler struct {
	communitySentenceRepository      domain.CommunitySentenceRepository
	communitySentenceDraftRepository domain.CommunitySentenceDraftRepository
	nlpRepository                    domain.NlpRepository
}

func NewPromoteCommunitySentenceDraftHandler(
	communitySentenceRepository domain.CommunitySentenceRepository,
	communitySentenceDraftRepository domain.CommunitySentenceDraftRepository,
	nlpRepository domain.NlpRepository,
) PromoteCommunitySentenceDraftHandler {
	return PromoteCommunitySentenceDraftHandler{
		communitySentenceRepository:      communitySentenceRepository,
		communitySentenceDraftRepository: communitySentenceDraftRepository,
		nlpRepository:                    nlpRepository,
	}
}

func (h PromoteCommunitySentenceDraftHandler) PromoteCommunitySentenceDraft(ctx *appcontext.AppContext, req *vocabularypb.PromoteCommunitySentenceDraftRequest) (*vocabularypb.PromoteCommunitySentenceDraftResponse, error) {
	ctx.Logger().Info("[hub] new promote community sentence draft request", appcontext.Fields{
		"userID": req.GetUserId(), "sentenceID": req.GetSentenceId(),
	})

	ctx.Logger().Text("find draft sentence in db")
	draftSentence, err := h.communitySentenceDraftRepository.FindCommunitySentenceDraftByID(ctx, req.GetSentenceId())
	if err != nil {
		ctx.Logger().Error("failed to find draft sentence in db", err, appcontext.Fields{})
		return nil, err
	}
	if draftSentence == nil {
		ctx.Logger().ErrorText("draft sentence not found")
		return nil, apperrors.Vocabulary.InvalidSentence
	}
	if !draftSentence.IsCorrect {
		ctx.Logger().ErrorText("this sentence is not correct and cannot be promoted")
		return nil, apperrors.Vocabulary.CannotPromoteDraftSentence
	}
	if !draftSentence.IsOwner(req.GetUserId()) {
		ctx.Logger().ErrorText("this user is not owner of this sentence")
		return nil, apperrors.Common.Forbidden
	}

	ctx.Logger().Text("analyze draft sentence")
	sentenceAnalysisResult, err := h.nlpRepository.AnalyzeSentence(ctx, draftSentence.Content.English)
	if err != nil {
		ctx.Logger().Error("failed to analyze draft sentence", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Info("check sentence's level, we don't accept sentences with level Beginner", appcontext.Fields{"level": sentenceAnalysisResult.Level.String()})
	if sentenceAnalysisResult.Level == domain.SentenceLevelBeginner {
		ctx.Logger().ErrorText("this sentence cannot be promoted")
		return nil, apperrors.Vocabulary.SentenceIsTooSimple
	}

	ctx.Logger().Text("create new community sentence model")
	sentence, err := domain.NewCommunitySentence(req.GetUserId(), draftSentence.VocabularyID)
	if err != nil {
		ctx.Logger().Error("failed to create new community sentence model", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("set sentence data")
	if err = h.setSentenceData(ctx, sentence, *draftSentence, *sentenceAnalysisResult); err != nil {
		ctx.Logger().Error("failed to set sentence data", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("persist sentence in db")
	if err = h.communitySentenceRepository.CreateCommunitySentence(ctx, *sentence); err != nil {
		ctx.Logger().Error("failed to persist sentence in db", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("done promote community sentence draft request")
	return &vocabularypb.PromoteCommunitySentenceDraftResponse{}, nil
}

func (PromoteCommunitySentenceDraftHandler) setSentenceData(ctx *appcontext.AppContext, sentence *domain.CommunitySentence, draft domain.CommunitySentenceDraft, analysisData domain.NlpSentenceAnalysisResult) error {
	if err := sentence.SetContent(draft.Content); err != nil {
		ctx.Logger().Error("failed to set content", err, appcontext.Fields{})
		return err
	}

	if err := sentence.SetRequiredTense(draft.RequiredTense.String()); err != nil {
		ctx.Logger().Error("failed to set required tense", err, appcontext.Fields{})
		return err
	}

	if err := sentence.SetRequiredVocabulary(draft.RequiredVocabulary); err != nil {
		ctx.Logger().Error("failed to set required vocabulary", err, appcontext.Fields{})
		return err
	}

	if err := sentence.SetClauses(draft.Clauses); err != nil {
		ctx.Logger().Error("failed to set clauses", err, appcontext.Fields{})
		return err
	}

	if err := sentence.SetSentiment(draft.Sentiment.Polarity, draft.Sentiment.Subjectivity); err != nil {
		ctx.Logger().Error("failed to set sentiment", err, appcontext.Fields{})
		return err
	}

	if err := sentence.SetPosTags(analysisData.PosTags); err != nil {
		ctx.Logger().Error("failed to set pos tags", err, appcontext.Fields{})
		return err
	}

	if err := sentence.SetDependencies(analysisData.Dependencies); err != nil {
		ctx.Logger().Error("failed to set dependencies", err, appcontext.Fields{})
		return err
	}

	if err := sentence.SetVerbs(analysisData.Verbs); err != nil {
		ctx.Logger().Error("failed to set verbs", err, appcontext.Fields{})
		return err
	}

	if err := sentence.SetLevel(analysisData.Level.String()); err != nil {
		ctx.Logger().Error("failed to set level", err, appcontext.Fields{})
		return err
	}

	return nil
}
