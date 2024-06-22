package hub

import (
	"sync"

	"github.com/namhq1989/vocab-booster-english-hub/core/language"

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
	cachingRepository           domain.CachingRepository
}

func NewSearchVocabularyHandler(
	vocabularyRepository domain.VocabularyRepository,
	vocabularyExampleRepository domain.VocabularyExampleRepository,
	aiRepository domain.AIRepository,
	scraperRepository domain.ScraperRepository,
	ttsRepository domain.TTSRepository,
	nlpRepository domain.NlpRepository,
	queueRepository domain.QueueRepository,
	cachingRepository domain.CachingRepository,
) SearchVocabularyHandler {
	return SearchVocabularyHandler{
		vocabularyRepository:        vocabularyRepository,
		vocabularyExampleRepository: vocabularyExampleRepository,
		aiRepository:                aiRepository,
		scraperRepository:           scraperRepository,
		ttsRepository:               ttsRepository,
		nlpRepository:               nlpRepository,
		queueRepository:             queueRepository,
		cachingRepository:           cachingRepository,
	}
}

func (h SearchVocabularyHandler) SearchVocabulary(ctx *appcontext.AppContext, req *vocabularypb.SearchVocabularyRequest) (*vocabularypb.SearchVocabularyResponse, error) {
	ctx.Logger().Info("new search vocabulary request", appcontext.Fields{"term": req.GetTerm()})
	var result = &vocabularypb.SearchVocabularyResponse{
		Found:       false,
		Suggestions: make([]string, 0),
		Vocabulary:  nil,
	}

	ctx.Logger().Text("find in caching first")
	vocabulary, _ := h.cachingRepository.GetVocabularyByTerm(ctx, req.GetTerm())
	if vocabulary != nil {
		ctx.Logger().Text("vocabulary found in caching layer, find related data and response")
		var examples = make([]domain.VocabularyExample, 0)
		examples, _ = h.cachingRepository.GetVocabularyExamplesByVocabularyID(ctx, vocabulary.ID)
		result.Found = true
		result.Vocabulary = dto.ConvertVocabularyFromDomainToGrpc(*vocabulary, examples)
		return result, nil
	}

	ctx.Logger().Text("find vocabulary in db with term")
	vocabulary, err := h.vocabularyRepository.FindVocabularyByTerm(ctx, req.GetTerm())
	if err != nil {
		ctx.Logger().Error("failed to find vocabulary with term", err, appcontext.Fields{})
		return nil, err
	}
	if vocabulary != nil {
		ctx.Logger().Text("vocabulary found in db, find related data and response")
		var examples = make([]domain.VocabularyExample, 0)
		examples, err = h.vocabularyExampleRepository.FindVocabularyExamplesByVocabularyID(ctx, vocabulary.ID)
		if err != nil {
			ctx.Logger().Error("failed to find vocabulary examples", err, appcontext.Fields{})
		}

		ctx.Logger().Text("cache result")
		if err = h.cacheResult(ctx, *vocabulary, examples); err != nil {
			ctx.Logger().Error("failed to cache result", err, appcontext.Fields{})
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
	if err != nil {
		ctx.Logger().Error("failed to generate sound for the vocabulary", err, appcontext.Fields{})
		return nil, err
	}

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

	ctx.Logger().Text("enqueue tasks")
	if err = h.enqueueTasks(ctx, vocabulary, examples); err != nil {
		ctx.Logger().Error("failed to enqueue tasks", err, appcontext.Fields{})
	}

	ctx.Logger().Text("cache result")
	if err = h.cacheResult(ctx, *vocabulary, examples); err != nil {
		ctx.Logger().Error("failed to cache result", err, appcontext.Fields{})
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

			analysisResult, err := h.nlpRepository.AnalyzeSentence(ctx, e.Example)
			if err != nil {
				ctx.Logger().Error("failed to analyze sentence", err, appcontext.Fields{"sentence": e.Example})
				return
			}

			example, err := domain.NewVocabularyExample(vocabulary.ID)
			if err != nil {
				ctx.Logger().Error("failed to create vocabulary example", err, appcontext.Fields{"vocabularyID": vocabulary.ID})
				return
			}

			if err = example.SetContent(e.Example, analysisResult.Translated); err != nil {
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

			translatedDefinition := language.TranslatedLanguages{}
			mainWordDefinitionTranslated, err := h.nlpRepository.TranslateDefinition(ctx, e.Definition)
			if err != nil {
				ctx.Logger().Error("failed to translate main word definition", err, appcontext.Fields{"definition": e.Definition})
			} else {
				translatedDefinition = *mainWordDefinitionTranslated
			}

			if err = example.SetMainWordData(e.Word, vocabulary.Term, e.Definition, e.Pos, translatedDefinition); err != nil {
				ctx.Logger().Error("failed to set vocabulary example's word data", err, appcontext.Fields{"word": e.Word, "definition": e.Definition, "pos": e.Pos})
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

func (h SearchVocabularyHandler) cacheResult(ctx *appcontext.AppContext, vocabulary domain.Vocabulary, examples []domain.VocabularyExample) error {
	ctx.Logger().Text("cache vocabulary")
	if err := h.cachingRepository.SetVocabularyByTerm(ctx, vocabulary.Term, &vocabulary); err != nil {
		ctx.Logger().Error("failed to cache vocabulary", err, appcontext.Fields{})
	}

	ctx.Logger().Text("cache vocabulary examples")
	if err := h.cachingRepository.SetVocabularyExamplesByVocabularyID(ctx, vocabulary.ID, examples); err != nil {
		ctx.Logger().Error("failed to cache vocabulary examples", err, appcontext.Fields{})
	}
	return nil
}
