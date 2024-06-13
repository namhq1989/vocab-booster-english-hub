package hub

import (
	"sync"

	"github.com/namhq1989/vocab-booster-english-hub/core/appcontext"
	apperrors "github.com/namhq1989/vocab-booster-english-hub/core/error"
	"github.com/namhq1989/vocab-booster-english-hub/internal/genproto/vocabularypb"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/domain"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/dto"
)

type SearchVocabularyHandler struct {
	vocabularyRepository        domain.VocabularyRepository
	vocabularyExampleRepository domain.VocabularyExampleRepository
	aiRepository                domain.AIRepository
	scraperRepository           domain.ScraperRepository
	ttsRepository               domain.TTSRepository
}

func NewSearchVocabularyHandler(
	vocabularyRepository domain.VocabularyRepository,
	vocabularyExampleRepository domain.VocabularyExampleRepository,
	aiRepository domain.AIRepository,
	scraperRepository domain.ScraperRepository,
	ttsRepository domain.TTSRepository,
) SearchVocabularyHandler {
	return SearchVocabularyHandler{
		vocabularyRepository:        vocabularyRepository,
		vocabularyExampleRepository: vocabularyExampleRepository,
		aiRepository:                aiRepository,
		scraperRepository:           scraperRepository,
		ttsRepository:               ttsRepository,
	}
}

func (h SearchVocabularyHandler) SearchVocabulary(ctx *appcontext.AppContext, req *vocabularypb.SearchVocabularyRequest) (*vocabularypb.SearchVocabularyResponse, error) {
	ctx.Logger().Info("new search vocabulary request", appcontext.Fields{"term": req.GetTerm()})
	var result = &vocabularypb.SearchVocabularyResponse{
		Found:       false,
		Suggestions: make([]string, 0),
		Vocabulary:  nil,
	}

	ctx.Logger().Text("find vocabulary in db with term")
	vocabulary, err := h.vocabularyRepository.FindVocabularyByTerm(ctx, req.GetTerm())
	if err != nil {
		ctx.Logger().Error("failed to find vocabulary with term", err, appcontext.Fields{})
		return nil, err
	}
	if vocabulary != nil {
		ctx.Logger().Text("vocabulary found in db, find examples")
		var examples = make([]domain.VocabularyExample, 0)
		examples, err = h.vocabularyExampleRepository.FindVocabularyExamplesByVocabularyID(ctx, vocabulary.ID)
		if err != nil {
			ctx.Logger().Error("failed to find vocabulary examples", err, appcontext.Fields{})
		}

		ctx.Logger().Text("respond data")
		result.Found = true
		result.Vocabulary = dto.ConvertVocabularyFromDomainToGrpc(*vocabulary, examples)
		return result, nil
	}

	ctx.Logger().Text("vocabulary not found, determine the term is valid or not")
	isValidTerm, suggestions, err := h.scraperRepository.IsTermValid(ctx, req.Term)
	if err != nil {
		ctx.Logger().Error("failed to determine the term is valid or not", err, appcontext.Fields{})
		return nil, err
	}
	if !isValidTerm {
		ctx.Logger().ErrorText("the term is not valid, respond")
		result.Suggestions = suggestions
		return result, nil
	}

	ctx.Logger().Text("the term is valid, create new vocabulary model")
	vocabulary, err = domain.NewVocabulary(req.PerformerId, req.Term)
	if err != nil {
		ctx.Logger().Error("failed to create new vocabulary", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("fetch data from internet")
	var (
		vocabularyData        *domain.VocabularyData
		vocabularyExamples    []domain.VocabularyExample
		soundGenerationResult *domain.TTSGenerateVocabularyPronunciationSoundResult
		wg                    sync.WaitGroup
	)
	wg.Add(3)

	go func() {
		defer wg.Done()

		ctx.Logger().Text("call OpenAI's api to fetch vocabulary data")
		vocabularyData, err = h.aiRepository.GetVocabularyData(ctx, req.GetTerm())
		if err != nil {
			ctx.Logger().Error("failed to get vocabulary data", err, appcontext.Fields{})
		}
	}()

	go func() {
		defer wg.Done()

		ctx.Logger().Text("call OpenAI's api to fetch vocabulary examples")
		vocabularyExamples, err = h.aiRepository.GetVocabularyExamples(ctx, vocabulary.ID, req.GetTerm())
		if err != nil {
			ctx.Logger().Error("failed to get vocabulary examples", err, appcontext.Fields{})
		}
	}()

	go func() {
		defer wg.Done()

		ctx.Logger().Text("call AWS Polly to generate pronunciation audio")
		soundGenerationResult, err = h.ttsRepository.GenerateVocabularyPronunciationSound(ctx, req.GetTerm())
	}()

	wg.Wait()

	if vocabularyData == nil {
		ctx.Logger().ErrorText("vocabulary data is null, respond")
		return nil, apperrors.Common.BadRequest
	}

	ctx.Logger().Text("set vocabulary data")
	if err = vocabulary.SetIPA(vocabularyData.IPA); err != nil {
		ctx.Logger().Error("failed to set vocabulary's ipa", err, appcontext.Fields{})
		return nil, err
	}
	if err = vocabulary.SetLexicalRelations(vocabularyData.Synonyms, vocabularyData.Antonyms); err != nil {
		ctx.Logger().Error("failed to set vocabulary's lexical relations", err, appcontext.Fields{})
		return nil, err
	}
	partsOfSpeech := make([]string, 0)
	for _, example := range vocabularyExamples {
		partsOfSpeech = append(partsOfSpeech, example.POS.String())
	}
	if err = vocabulary.SetPartsOfSpeech(partsOfSpeech); err != nil {
		ctx.Logger().Error("failed to set vocabulary's parts of speech", err, appcontext.Fields{})
		return nil, err
	}
	if soundGenerationResult != nil {
		if err = vocabulary.SetAudioName(soundGenerationResult.FileName); err != nil {
			ctx.Logger().Error("failed to set vocabulary's audio name", err, appcontext.Fields{})
			return nil, err
		}
	}

	ctx.Logger().Text("persist vocabulary data to db")
	if err = h.vocabularyRepository.CreateVocabulary(ctx, *vocabulary); err != nil {
		ctx.Logger().Error("failed to insert vocabulary", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("persist vocabulary examples to db")
	if err = h.vocabularyExampleRepository.CreateVocabularyExamples(ctx, vocabularyExamples); err != nil {
		ctx.Logger().Error("failed to insert vocabulary examples", err, appcontext.Fields{})
	}

	ctx.Logger().Text("done search vocabulary request")
	result.Found = true
	result.Vocabulary = dto.ConvertVocabularyFromDomainToGrpc(*vocabulary, vocabularyExamples)
	return result, nil
}
