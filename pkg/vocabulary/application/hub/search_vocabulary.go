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
	nlpRepository               domain.NlpRepository
	queueRepository             domain.QueueRepository
}

func NewSearchVocabularyHandler(
	vocabularyRepository domain.VocabularyRepository,
	vocabularyExampleRepository domain.VocabularyExampleRepository,
	aiRepository domain.AIRepository,
	scraperRepository domain.ScraperRepository,
	ttsRepository domain.TTSRepository,
	nlpRepository domain.NlpRepository,
	queueRepository domain.QueueRepository,
) SearchVocabularyHandler {
	return SearchVocabularyHandler{
		vocabularyRepository:        vocabularyRepository,
		vocabularyExampleRepository: vocabularyExampleRepository,
		aiRepository:                aiRepository,
		scraperRepository:           scraperRepository,
		ttsRepository:               ttsRepository,
		nlpRepository:               nlpRepository,
		queueRepository:             queueRepository,
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

	ctx.Logger().Text("fetch vocabulary data with GPT")
	aiVocabularyData, err := h.aiRepository.GetVocabularyData(ctx, req.GetTerm())
	if err != nil {
		ctx.Logger().Error("failed to fetch vocabulary data with GPT", err, appcontext.Fields{})
		return nil, err
	}
	if aiVocabularyData == nil {
		ctx.Logger().ErrorText("aiVocabularyData is null, respond")
		return nil, apperrors.Common.BadRequest
	}

	ctx.Logger().Text("translate and analyze the examples")
	examples, err := h.analyzeExamples(ctx, *vocabulary, aiVocabularyData.Examples)
	if err != nil {
		ctx.Logger().Error("failed to analyze examples", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("generate sound for the vocabulary")
	soundGenerationResult, err := h.ttsRepository.GenerateVocabularySound(ctx, req.GetTerm())

	ctx.Logger().Text("set vocabulary data")
	if err = h.setVocabularyData(ctx, vocabulary, aiVocabularyData, soundGenerationResult); err != nil {
		ctx.Logger().Error("failed to set vocabulary data", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("persist vocabulary data to db")
	if err = h.vocabularyRepository.CreateVocabulary(ctx, *vocabulary); err != nil {
		ctx.Logger().Error("failed to insert vocabulary", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("persist vocabulary examples to db")
	if err = h.vocabularyExampleRepository.CreateVocabularyExamples(ctx, examples); err != nil {
		ctx.Logger().Error("failed to insert vocabulary examples", err, appcontext.Fields{})
	}

	ctx.Logger().Text("done search vocabulary request")
	result.Found = true
	result.Vocabulary = dto.ConvertVocabularyFromDomainToGrpc(*vocabulary, examples)
	return result, nil
}

func (SearchVocabularyHandler) setVocabularyData(ctx *appcontext.AppContext, vocabulary *domain.Vocabulary, aiVocabularyData *domain.AIVocabularyData, soundGenerationResult *domain.TTSGenerateSoundResult) error {
	if err := vocabulary.SetIPA(aiVocabularyData.IPA); err != nil {
		ctx.Logger().Error("failed to set vocabulary's ipa", err, appcontext.Fields{"ipa": aiVocabularyData.IPA})
		return err
	}

	if err := vocabulary.SetLexicalRelations(aiVocabularyData.Synonyms, aiVocabularyData.Antonyms); err != nil {
		ctx.Logger().Error("failed to set vocabulary's lexical relations", err, appcontext.Fields{"synonyms": aiVocabularyData.Synonyms, "antonyms": aiVocabularyData.Antonyms})
		return err
	}

	if err := vocabulary.SetPartsOfSpeech(aiVocabularyData.PosTags); err != nil {
		ctx.Logger().Error("failed to set vocabulary's parts of speech", err, appcontext.Fields{"posTags": aiVocabularyData.PosTags})
		return err
	}

	if soundGenerationResult != nil {
		if err := vocabulary.SetAudio(soundGenerationResult.FileName); err != nil {
			ctx.Logger().Error("failed to set vocabulary's audio name", err, appcontext.Fields{"audio": soundGenerationResult.FileName})
			return err
		}
	}

	return nil
}

func (h SearchVocabularyHandler) analyzeExamples(ctx *appcontext.AppContext, vocabulary domain.Vocabulary, aiExamples []domain.AIVocabularyExample) ([]domain.VocabularyExample, error) {
	var (
		wg     sync.WaitGroup
		result = make([]domain.VocabularyExample, len(aiExamples))
	)

	wg.Add(len(aiExamples))

	for i, aiExample := range aiExamples {
		go func(i int, e domain.AIVocabularyExample) {
			defer wg.Done()

			analysisResult, err := h.nlpRepository.AnalyzeSentence(ctx, e.Example)
			if err != nil {
				ctx.Logger().Error("failed to analyze sentence", err, appcontext.Fields{"sentence": e.Example})
				return
			}

			example, err := domain.NewVocabularyExample(vocabulary.ID)
			if err != nil {
				ctx.Logger().Error("failed to create vocabulary example", err, appcontext.Fields{})
				return
			}

			if err = example.SetContent(e.Example, analysisResult.Translated); err != nil {
				ctx.Logger().Error("failed to set vocabulary example's content", err, appcontext.Fields{})
				return
			}

			if err = example.SetPosTags(analysisResult.PosTags); err != nil {
				ctx.Logger().Error("failed to set vocabulary example's pos tags", err, appcontext.Fields{})
				return
			}

			if err = example.SetDependencies(analysisResult.Dependencies); err != nil {
				ctx.Logger().Error("failed to set vocabulary example's dependencies", err, appcontext.Fields{})
				return
			}

			if err = example.SetSentiment(analysisResult.Sentiment.Polarity, analysisResult.Sentiment.Subjectivity); err != nil {
				ctx.Logger().Error("failed to set vocabulary example's sentiment", err, appcontext.Fields{})
				return
			}

			if err = example.SetVerbs(analysisResult.Verbs); err != nil {
				ctx.Logger().Error("failed to set vocabulary example's verbs", err, appcontext.Fields{})
				return
			}

			if err = example.SetWordData(e.Word, e.Definition, e.Pos); err != nil {
				ctx.Logger().Error("failed to set vocabulary example's word data", err, appcontext.Fields{})
				return
			}

			result[i] = *example
		}(i, aiExample)
	}

	wg.Wait()

	return result, nil
}

func (h SearchVocabularyHandler) enqueueTasks(ctx *appcontext.AppContext, vocabulary *domain.Vocabulary) error {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()

		ctx.Logger().Text("add task newVocabularyCreated")
		if err := h.queueRepository.NewVocabularyCreated(ctx, domain.QueueNewVocabularyCreatedPayload{
			Vocabulary: *vocabulary,
		}); err != nil {
			ctx.Logger().Error("failed to enqueue task", err, appcontext.Fields{})
		}
	}()

	return nil
}
