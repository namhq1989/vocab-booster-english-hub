package hub

import (
	"github.com/namhq1989/vocab-booster-english-hub/internal/genproto/vocabularypb"
	apperrors "github.com/namhq1989/vocab-booster-english-hub/internal/utils/error"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/domain"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
	"github.com/namhq1989/vocab-booster-utilities/language"
)

type UpdateCommunitySentenceDraftHandler struct {
	communitySentenceDraftRepository domain.CommunitySentenceDraftRepository
	nlpRepository                    domain.NlpRepository
}

func NewUpdateCommunitySentenceDraftHandler(
	communitySentenceDraftRepository domain.CommunitySentenceDraftRepository,
	nlpRepository domain.NlpRepository,
) UpdateCommunitySentenceDraftHandler {
	return UpdateCommunitySentenceDraftHandler{
		communitySentenceDraftRepository: communitySentenceDraftRepository,
		nlpRepository:                    nlpRepository,
	}
}

func (h UpdateCommunitySentenceDraftHandler) UpdateCommunitySentenceDraft(ctx *appcontext.AppContext, req *vocabularypb.UpdateCommunitySentenceDraftRequest) (*vocabularypb.UpdateCommunitySentenceDraftResponse, error) {
	ctx.Logger().Info("[hub] new update community sentence draft request", appcontext.Fields{
		"userID": req.GetUserId(), "sentenceID": req.GetSentenceId(),
		"sentence": req.GetSentence(), "lang": req.GetLang(),
	})

	ctx.Logger().Text("find sentence by id in db")
	sentence, err := h.communitySentenceDraftRepository.FindCommunitySentenceDraftByID(ctx, req.GetSentenceId())
	if err != nil {
		ctx.Logger().Error("failed to find sentence by id in db", err, appcontext.Fields{})
		return nil, err
	}
	if sentence == nil {
		ctx.Logger().ErrorText("sentence not found")
		return nil, apperrors.Vocabulary.InvalidSentence
	}

	if !sentence.IsOwner(req.GetUserId()) {
		ctx.Logger().ErrorText("this user is not owner of this sentence")
		return nil, apperrors.Common.Forbidden
	}
	if sentence.IsCorrect {
		ctx.Logger().ErrorText("this sentence is already correct and cannot be modified anymore")
		return nil, apperrors.Vocabulary.SentenceIsAlreadyCorrect
	}

	ctx.Logger().Text("call NLP to evaluate grammar")
	grammarErrors, err := h.nlpRepository.GrammarCheck(ctx, req.GetSentence())
	if err != nil {
		ctx.Logger().Error("failed to call NLP to evaluate grammar", err, appcontext.Fields{})
		return nil, err
	}

	if len(grammarErrors) > 0 {
		ctx.Logger().Text("has grammar errors")
		err = h.hasGrammarErrors(ctx, sentence, req.GetSentence(), grammarErrors)
	} else {
		ctx.Logger().Text("no grammar errors")
		err = h.noGrammarErrors(ctx, sentence, req.GetSentence())
	}
	if err != nil {
		ctx.Logger().Error("failed to set grammar data", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("update sentence draft in db")
	sentence.SetUpdatedAt()
	if err = h.communitySentenceDraftRepository.UpdateCommunitySentenceDraft(ctx, *sentence); err != nil {
		ctx.Logger().Error("failed to persist sentence draft in db", err, appcontext.Fields{})
		return nil, err
	}

	return nil, nil
}

func (UpdateCommunitySentenceDraftHandler) hasGrammarErrors(ctx *appcontext.AppContext, sentence *domain.CommunitySentenceDraft, content string, grammarErrors []domain.SentenceGrammarError) error {
	ctx.Logger().Text("set sentence data")
	if err := sentence.SetContent(language.Multilingual{
		English: content,
	}); err != nil {
		ctx.Logger().Error("failed to set content", err, appcontext.Fields{})
		return err
	}

	if err := sentence.SetGrammarErrors(grammarErrors); err != nil {
		ctx.Logger().Error("failed to set grammar errors", err, appcontext.Fields{})
		return err
	}

	return nil
}

func (h UpdateCommunitySentenceDraftHandler) noGrammarErrors(ctx *appcontext.AppContext, sentence *domain.CommunitySentenceDraft, content string) error {
	ctx.Logger().Text("call NLP to evaluate the sentence")
	sentenceEvaluationResult, err := h.nlpRepository.EvaluateSentence(ctx, content, sentence.RequiredTense.String(), sentence.RequiredVocabularies)
	if err != nil {
		ctx.Logger().Error("failed to call NLP to evaluate the sentence", err, appcontext.Fields{})
		return err
	}

	ctx.Logger().Text("set sentence data")
	if err = sentence.SetContent(sentenceEvaluationResult.Translated); err != nil {
		ctx.Logger().Error("failed to set content", err, appcontext.Fields{})
		return err
	}

	if err = sentence.SetGrammarErrors([]domain.SentenceGrammarError{}); err != nil {
		ctx.Logger().Error("failed to set grammar errors", err, appcontext.Fields{})
		return err
	}

	if err = sentence.SetSentiment(sentenceEvaluationResult.Sentiment.Polarity, sentenceEvaluationResult.Sentiment.Subjectivity); err != nil {
		ctx.Logger().Error("failed to set sentiment", err, appcontext.Fields{})
		return err
	}

	if err = sentence.SetClauses(sentenceEvaluationResult.Clauses); err != nil {
		ctx.Logger().Error("failed to set clauses", err, appcontext.Fields{})
		return err
	}

	// set error code
	if !sentenceEvaluationResult.IsEnglish {
		sentence.SetErrorCode(domain.SentenceErrorCodeIsNotEnglish)
	} else if !sentenceEvaluationResult.IsVocabularyCorrect {
		sentence.SetErrorCode(domain.SentenceErrorCodeInvalidVocabulary)
	} else if !sentenceEvaluationResult.IsTenseCorrect {
		sentence.SetErrorCode(domain.SentenceErrorCodeInvalidTense)
	} else {
		sentence.SetErrorCode(domain.SentenceErrorCodeEmpty)
	}
	if !sentence.ErrorCode.IsEmpty() {
		ctx.Logger().Error("sentence has error code", nil, appcontext.Fields{"error_code": sentence.ErrorCode.String()})
	} else {
		ctx.Logger().Text("sentence has no error code")
	}

	return nil
}
