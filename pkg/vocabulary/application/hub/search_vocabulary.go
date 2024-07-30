package hub

import (
	"sync"

	"github.com/namhq1989/vocab-booster-english-hub/internal/genproto/vocabularypb"
	apperrors "github.com/namhq1989/vocab-booster-english-hub/internal/utils/error"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/domain"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/dto"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type SearchVocabularyHandler struct {
	vocabularyRepository               domain.VocabularyRepository
	vocabularyExampleRepository        domain.VocabularyExampleRepository
	userBookmarkedVocabularyRepository domain.UserBookmarkedVocabularyRepository
	aiRepository                       domain.AIRepository
	externalApiRepository              domain.ExternalApiRepository
	scraperRepository                  domain.ScraperRepository
	ttsRepository                      domain.TTSRepository
	nlpRepository                      domain.NlpRepository
	queueRepository                    domain.QueueRepository
	cachingRepository                  domain.CachingRepository
	service                            domain.Service
}

func NewSearchVocabularyHandler(
	vocabularyRepository domain.VocabularyRepository,
	vocabularyExampleRepository domain.VocabularyExampleRepository,
	userBookmarkedVocabularyRepository domain.UserBookmarkedVocabularyRepository,
	aiRepository domain.AIRepository,
	externalApiRepository domain.ExternalApiRepository,
	scraperRepository domain.ScraperRepository,
	ttsRepository domain.TTSRepository,
	nlpRepository domain.NlpRepository,
	queueRepository domain.QueueRepository,
	cachingRepository domain.CachingRepository,
	service domain.Service,
) SearchVocabularyHandler {
	return SearchVocabularyHandler{
		vocabularyRepository:               vocabularyRepository,
		vocabularyExampleRepository:        vocabularyExampleRepository,
		userBookmarkedVocabularyRepository: userBookmarkedVocabularyRepository,
		aiRepository:                       aiRepository,
		externalApiRepository:              externalApiRepository,
		scraperRepository:                  scraperRepository,
		ttsRepository:                      ttsRepository,
		nlpRepository:                      nlpRepository,
		queueRepository:                    queueRepository,
		cachingRepository:                  cachingRepository,
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
	if vocabulary != nil {
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

	ctx.Logger().Text("fetch vocabulary data with Datamuse")
	datamuseData, err := h.externalApiRepository.SearchTermWithDatamuse(ctx, req.GetTerm())
	if err != nil {
		ctx.Logger().Error("failed to fetch vocabulary data with Datamuse", err, appcontext.Fields{})
		return nil, err
	}
	if datamuseData == nil {
		ctx.Logger().ErrorText("datamuseData is null, respond")
		return nil, apperrors.Common.BadRequest
	}

	ctx.Logger().Text("generate vocabulary examples with GPT")
	aiExamplesData, err := h.aiRepository.VocabularyExamples(ctx, req.GetTerm(), datamuseData.PartsOfSpeech)
	if err != nil {
		ctx.Logger().Error("failed to generate vocabulary examples with GPT", err, appcontext.Fields{})
		return nil, err
	}
	if aiExamplesData == nil {
		ctx.Logger().ErrorText("aiExamplesData is null, respond")
		return nil, apperrors.Common.BadRequest
	}

	ctx.Logger().Text("translate and analyze the examples")
	examples, err := h.analyzeExamples(ctx, *vocabulary, aiExamplesData)
	if err != nil {
		ctx.Logger().Error("failed to analyze examples", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("generate sound for the vocabulary")
	soundGenerationResult, err := h.ttsRepository.GenerateVocabularySound(ctx, req.GetTerm())
	if err != nil {
		ctx.Logger().Error("failed to generate sound for the vocabulary", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("set vocabulary data")
	if err = h.setVocabularyData(ctx, vocabulary, datamuseData, soundGenerationResult); err != nil {
		ctx.Logger().Error("failed to set vocabulary data", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("persist vocabulary in db")
	if err = h.vocabularyRepository.CreateVocabulary(ctx, *vocabulary); err != nil {
		ctx.Logger().Error("failed to insert vocabulary", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("persist vocabulary examples to db")
	if err = h.vocabularyExampleRepository.CreateVocabularyExamples(ctx, examples); err != nil {
		ctx.Logger().Error("failed to insert vocabulary examples", err, appcontext.Fields{})
	}

	ctx.Logger().Text("enqueue tasks")
	if err = h.enqueueTasks(ctx, vocabulary, examples); err != nil {
		ctx.Logger().Error("failed to enqueue tasks", err, appcontext.Fields{})
	}

	ctx.Logger().Text("cache result")
	if err = h.cacheResult(ctx, *vocabulary); err != nil {
		ctx.Logger().Error("failed to cache result", err, appcontext.Fields{})
	}

	ctx.Logger().Text("check bookmarked")
	ubv, bookmarkedErr := h.userBookmarkedVocabularyRepository.FindBookmarkedVocabulary(ctx, req.GetPerformerId(), vocabulary.ID)
	if bookmarkedErr != nil {
		ctx.Logger().Error("failed to check bookmarked vocabulary", bookmarkedErr, appcontext.Fields{"vocabularyID": vocabulary.ID})
	}
	isBookmarked := ubv != nil

	ctx.Logger().Text("done search vocabulary request")
	result.Found = true
	result.Vocabulary = dto.ConvertVocabularyFromDomainToGrpc(*vocabulary, examples, isBookmarked)
	return result, nil
}

func (SearchVocabularyHandler) setVocabularyData(ctx *appcontext.AppContext, vocabulary *domain.Vocabulary, datamuseData *domain.DatamuseSearchTermResult, soundGenerationResult *domain.TTSGenerateSoundResult) error {
	if err := vocabulary.SetDefinitions(datamuseData.Definitions); err != nil {
		ctx.Logger().Error("failed to set vocabulary's definitions", err, appcontext.Fields{"definitions": datamuseData.Definitions})
		return err
	}

	if err := vocabulary.SetIPA(datamuseData.Ipa); err != nil {
		ctx.Logger().Error("failed to set vocabulary's ipa", err, appcontext.Fields{"ipa": datamuseData.Ipa})
		return err
	}

	if err := vocabulary.SetFrequency(datamuseData.Frequency); err != nil {
		ctx.Logger().Error("failed to set vocabulary's frequency", err, appcontext.Fields{"frequency": datamuseData.Frequency})
		return err
	}

	if err := vocabulary.SetLexicalRelations(datamuseData.Synonyms, datamuseData.Antonyms); err != nil {
		ctx.Logger().Error("failed to set vocabulary's lexical relations", err, appcontext.Fields{"synonyms": datamuseData.Synonyms, "antonyms": datamuseData.Antonyms})
		return err
	}

	if err := vocabulary.SetPartsOfSpeech(datamuseData.PartsOfSpeech); err != nil {
		ctx.Logger().Error("failed to set vocabulary's parts of speech", err, appcontext.Fields{"partsOfSpeech": datamuseData.PartsOfSpeech})
		return err
	}

	if soundGenerationResult != nil {
		if err := vocabulary.SetAudio(soundGenerationResult.FileName); err != nil {
			ctx.Logger().Error("failed to set vocabulary's audio", err, appcontext.Fields{"audio": soundGenerationResult.FileName})
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

			analysisResult, err := h.nlpRepository.AnalyzeSentence(ctx, aiExample.Example, vocabulary.Term)
			if err != nil {
				ctx.Logger().Error("failed to analyze sentence", err, appcontext.Fields{"sentence": e.Example})
				return
			}

			example, err := domain.NewVocabularyExample(vocabulary.ID)
			if err != nil {
				ctx.Logger().Error("failed to create vocabulary example", err, appcontext.Fields{"vocabularyID": vocabulary.ID})
				return
			}

			if err = example.SetContent(analysisResult.Translated); err != nil {
				ctx.Logger().Error("failed to set vocabulary example's content", err, appcontext.Fields{"content": e.Example, "translated": analysisResult.Translated})
				return
			}

			if err = example.SetPosTags(analysisResult.PosTags); err != nil {
				ctx.Logger().Error("failed to set vocabulary example's pos tags", err, appcontext.Fields{"posTags": analysisResult.PosTags})
				return
			}

			if err = example.SetDependencies(analysisResult.Dependencies); err != nil {
				ctx.Logger().Error("failed to set vocabulary example's dependencies", err, appcontext.Fields{"dependencies": analysisResult.Dependencies})
				return
			}

			if err = example.SetSentiment(analysisResult.Sentiment.Polarity, analysisResult.Sentiment.Subjectivity); err != nil {
				ctx.Logger().Error("failed to set vocabulary example's sentiment", err, appcontext.Fields{"polarity": analysisResult.Sentiment.Polarity, "subjectivity": analysisResult.Sentiment.Subjectivity})
				return
			}

			if err = example.SetVerbs(analysisResult.Verbs); err != nil {
				ctx.Logger().Error("failed to set vocabulary example's verbs", err, appcontext.Fields{"verbs": analysisResult.Verbs})
				return
			}

			if err = example.SetLevel(analysisResult.Level.String()); err != nil {
				ctx.Logger().Error("failed to set vocabulary example's level", err, appcontext.Fields{"level": analysisResult.Level.String()})
				return
			}

			if err = example.SetMainWordData(analysisResult.MainWord); err != nil {
				ctx.Logger().Error("failed to set vocabulary example's word data", err, appcontext.Fields{"word": e.Word})
				return
			}

			result[i] = *example
		}(i, aiExample)
	}

	wg.Wait()

	return result, nil
}

func (h SearchVocabularyHandler) enqueueTasks(ctx *appcontext.AppContext, vocabulary *domain.Vocabulary, examples []domain.VocabularyExample) error {
	var (
		wg         sync.WaitGroup
		totalTasks = 1 + len(examples)
	)
	wg.Add(totalTasks)

	go func() {
		defer wg.Done()

		ctx.Logger().Text("add task newVocabularyCreated")
		if err := h.queueRepository.NewVocabularyCreated(ctx, domain.QueueNewVocabularyCreatedPayload{
			Vocabulary: *vocabulary,
		}); err != nil {
			ctx.Logger().Error("failed to enqueue task", err, appcontext.Fields{})
		}
	}()

	for _, e := range examples {
		go func(example domain.VocabularyExample) {
			defer wg.Done()

			ctx.Logger().Info("add task newVocabularyExampleCreated", appcontext.Fields{"exampleID": example.ID})
			if err := h.queueRepository.NewVocabularyExampleCreated(ctx, domain.QueueNewVocabularyExampleCreatedPayload{
				Vocabulary: *vocabulary,
				Example:    example,
			}); err != nil {
				ctx.Logger().Error("failed to enqueue task", err, appcontext.Fields{})
			}
		}(e)
	}

	defer wg.Wait()

	return nil
}

func (h SearchVocabularyHandler) cacheResult(ctx *appcontext.AppContext, vocabulary domain.Vocabulary) error {
	ctx.Logger().Text("cache vocabulary")
	if err := h.cachingRepository.SetVocabularyByTerm(ctx, vocabulary.Term, &vocabulary); err != nil {
		ctx.Logger().Error("failed to cache vocabulary", err, appcontext.Fields{})
	}

	// ctx.Logger().Text("cache examples")
	// if err := h.cachingRepository.SetVocabularyExamplesByVocabularyID(ctx, vocabulary.ID, examples); err != nil {
	// 	ctx.Logger().Error("failed to cache examples", err, appcontext.Fields{})
	// }

	return nil
}
