package hub

import (
	"slices"

	"github.com/namhq1989/vocab-booster-english-hub/core/appcontext"
	apperrors "github.com/namhq1989/vocab-booster-english-hub/core/error"
	"github.com/namhq1989/vocab-booster-english-hub/internal/genproto/vocabularypb"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/domain"
)

type CreateCommunitySentenceDraftHandler struct {
	vocabularyRepository             domain.VocabularyRepository
	communitySentenceDraftRepository domain.CommunitySentenceDraftRepository
	nlpRepository                    domain.NlpRepository
}

func NewCreateCommunitySentenceDraftHandler(
	vocabularyRepository domain.VocabularyRepository,
	communitySentenceDraftRepository domain.CommunitySentenceDraftRepository,
	nlpRepository domain.NlpRepository,
) CreateCommunitySentenceDraftHandler {
	return CreateCommunitySentenceDraftHandler{
		vocabularyRepository:             vocabularyRepository,
		communitySentenceDraftRepository: communitySentenceDraftRepository,
		nlpRepository:                    nlpRepository,
	}
}

func (h CreateCommunitySentenceDraftHandler) CreateCommunitySentenceDraft(ctx *appcontext.AppContext, req *vocabularypb.CreateCommunitySentenceDraftRequest) (*vocabularypb.CreateCommunitySentenceDraftResponse, error) {
	ctx.Logger().Info("[hub] new create community sentence draft request", appcontext.Fields{
		"userID": req.GetUserId(), "vocabularyID": req.GetVocabularyId(),
		"sentence": req.GetSentence(), "vocabulary": req.GetVocabulary(),
		"tense": req.GetTense(), "lang": req.GetLang(),
	})

	ctx.Logger().Text("find vocabulary by id")
	vocabulary, err := h.vocabularyRepository.FindVocabularyByID(ctx, req.GetVocabularyId())
	if err != nil {
		ctx.Logger().Error("failed to find vocabulary by id", err, appcontext.Fields{})
		return nil, err
	}
	if vocabulary == nil {
		ctx.Logger().ErrorText("vocabulary not found")
		return nil, apperrors.Vocabulary.VocabularyNotFound
	}

	ctx.Logger().Info("check if vocabulary is in list sentence's required vocabulary or not", appcontext.Fields{"term": vocabulary.Term})
	if !slices.Contains(req.GetVocabulary(), vocabulary.Term) {
		ctx.Logger().ErrorText("vocabulary is not in list sentence's required vocabulary")
		return nil, apperrors.Vocabulary.InvalidVocabularyID
	} else {
		ctx.Logger().Text("valid vocabulary")
	}

	ctx.Logger().Text("create community sentence draft model")
	sentence, err := domain.NewCommunitySentenceDraft(req.GetUserId(), req.GetVocabularyId())
	if err != nil {
		ctx.Logger().Error("failed to create community sentence draft model", err, appcontext.Fields{})
		return nil, err
	}
	if sentence == nil {
		ctx.Logger().ErrorText("failed to create community sentence draft model")
		return nil, apperrors.Vocabulary.InvalidSentence
	}

	ctx.Logger().Text("call NLP to evaluate grammar")
	grammarErrors, err := h.nlpRepository.GrammarCheck(ctx, req.GetSentence())
	if err != nil {
		ctx.Logger().Error("failed to call NLP to evaluate grammar", err, appcontext.Fields{})
		return nil, err
	}

	if len(grammarErrors) > 0 {
		ctx.Logger().Text("has grammar errors")
		sentence, err = h.hasGrammarErrors(ctx, req, grammarErrors)
	} else {
		ctx.Logger().Text("no grammar errors")
		sentence, err = h.noGrammarErrors(ctx, req)
	}

	if err != nil {
		ctx.Logger().Error("failed to create community sentence draft", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("done create community sentence draft request")
	return &vocabularypb.CreateCommunitySentenceDraftResponse{
		Id: sentence.ID,
	}, nil
}

func (h CreateCommunitySentenceDraftHandler) hasGrammarErrors(ctx *appcontext.AppContext, req *vocabularypb.CreateCommunitySentenceDraftRequest, grammarErrors []domain.SentenceGrammarError) (*domain.CommunitySentenceDraft, error) {
	ctx.Logger().Text("create new sentence draft model")
	sentence, err := domain.NewCommunitySentenceDraft(req.GetUserId(), req.GetVocabularyId())
	if err != nil {
		ctx.Logger().Error("failed to create new sentence draft model", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("set sentence data")
	if err = sentence.SetContent(req.GetSentence()); err != nil {
		ctx.Logger().Error("failed to set content", err, appcontext.Fields{})
		return nil, err
	}

	if err = sentence.SetGrammarErrors(grammarErrors); err != nil {
		ctx.Logger().Error("failed to set grammar errors", err, appcontext.Fields{})
		return nil, err
	}

	if err = sentence.SetRequiredVocabulary(req.GetVocabulary()); err != nil {
		ctx.Logger().Error("failed to set required vocabulary", err, appcontext.Fields{})
		return nil, err
	}

	if err = sentence.SetRequiredTense(req.GetTense()); err != nil {
		ctx.Logger().Error("failed to set required tense", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("persist sentence draft in db")
	if err = h.communitySentenceDraftRepository.CreateCommunitySentenceDraft(ctx, *sentence); err != nil {
		ctx.Logger().Error("failed to persist sentence draft in db", err, appcontext.Fields{})
		return nil, err
	}

	return sentence, nil
}

func (h CreateCommunitySentenceDraftHandler) noGrammarErrors(ctx *appcontext.AppContext, req *vocabularypb.CreateCommunitySentenceDraftRequest) (*domain.CommunitySentenceDraft, error) {
	ctx.Logger().Text("create new sentence draft model")
	sentence, err := domain.NewCommunitySentenceDraft(req.GetUserId(), req.GetVocabularyId())
	if err != nil {
		ctx.Logger().Error("failed to create new sentence draft model", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("call NLP to evaluate the sentence")
	sentenceEvaluationResult, err := h.nlpRepository.EvaluateSentence(ctx, req.GetSentence(), req.GetTense(), req.GetVocabulary())
	if err != nil {
		ctx.Logger().Error("failed to call NLP to evaluate the sentence", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("set sentence data")
	if err = sentence.SetContent(req.GetSentence()); err != nil {
		ctx.Logger().Error("failed to set content", err, appcontext.Fields{})
		return nil, err
	}

	if err = sentence.SetGrammarErrors([]domain.SentenceGrammarError{}); err != nil {
		ctx.Logger().Error("failed to set grammar errors", err, appcontext.Fields{})
		return nil, err
	}

	if err = sentence.SetRequiredVocabulary(req.GetVocabulary()); err != nil {
		ctx.Logger().Error("failed to set required vocabulary", err, appcontext.Fields{})
		return nil, err
	}

	if err = sentence.SetRequiredTense(req.GetTense()); err != nil {
		ctx.Logger().Error("failed to set required tense", err, appcontext.Fields{})
		return nil, err
	}

	if err = sentence.SetSentiment(sentenceEvaluationResult.Sentiment.Polarity, sentenceEvaluationResult.Sentiment.Subjectivity); err != nil {
		ctx.Logger().Error("failed to set sentiment", err, appcontext.Fields{})
		return nil, err
	}

	if err = sentence.SetClauses(sentenceEvaluationResult.Clauses); err != nil {
		ctx.Logger().Error("failed to set clauses", err, appcontext.Fields{})
		return nil, err
	}

	if err = sentence.SetTranslated(sentenceEvaluationResult.Translated); err != nil {
		ctx.Logger().Error("failed to set translated", err, appcontext.Fields{})
		return nil, err
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

	ctx.Logger().Text("persist sentence draft in db")
	if err = h.communitySentenceDraftRepository.CreateCommunitySentenceDraft(ctx, *sentence); err != nil {
		ctx.Logger().Error("failed to persist sentence draft in db", err, appcontext.Fields{})
		return nil, err
	}

	return sentence, nil
}
