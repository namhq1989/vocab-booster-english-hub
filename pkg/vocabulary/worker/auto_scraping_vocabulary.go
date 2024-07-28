package worker

import (
	"sync"

	apperrors "github.com/namhq1989/vocab-booster-english-hub/internal/utils/error"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/domain"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type AutoScrapingVocabularyHandler struct {
	vocabularyRepository             domain.VocabularyRepository
	vocabularyExampleRepository      domain.VocabularyExampleRepository
	vocabularyScrapingItemRepository domain.VocabularyScrapingItemRepository
	aiRepository                     domain.AIRepository
	externalApiRepository            domain.ExternalApiRepository
	scraperRepository                domain.ScraperRepository
	ttsRepository                    domain.TTSRepository
	nlpRepository                    domain.NlpRepository
	queueRepository                  domain.QueueRepository
}

func NewAutoScrapingVocabularyHandler(
	vocabularyRepository domain.VocabularyRepository,
	vocabularyExampleRepository domain.VocabularyExampleRepository,
	vocabularyScrapingItemRepository domain.VocabularyScrapingItemRepository,
	aiRepository domain.AIRepository,
	externalApiRepository domain.ExternalApiRepository,
	scraperRepository domain.ScraperRepository,
	ttsRepository domain.TTSRepository,
	nlpRepository domain.NlpRepository,
	queueRepository domain.QueueRepository,
) AutoScrapingVocabularyHandler {
	return AutoScrapingVocabularyHandler{
		vocabularyRepository:             vocabularyRepository,
		vocabularyExampleRepository:      vocabularyExampleRepository,
		vocabularyScrapingItemRepository: vocabularyScrapingItemRepository,
		aiRepository:                     aiRepository,
		externalApiRepository:            externalApiRepository,
		scraperRepository:                scraperRepository,
		ttsRepository:                    ttsRepository,
		nlpRepository:                    nlpRepository,
		queueRepository:                  queueRepository,
	}
}

func (w AutoScrapingVocabularyHandler) AutoScrapingVocabulary(ctx *appcontext.AppContext, _ domain.QueueAutoScrapingVocabularyPayload) error {
	ctx.Logger().Text("picking random scraping item in db")
	scrapingItem, err := w.vocabularyScrapingItemRepository.RandomPickVocabularyScrapingItem(ctx)
	if err != nil {
		ctx.Logger().Error("failed to pick random vocabulary scrape item", err, appcontext.Fields{})
		return err
	}
	if scrapingItem == nil {
		ctx.Logger().Text("no item for scraping, respond")
		return nil
	}

	ctx.Logger().Info("item found, create new vocabulary model", appcontext.Fields{"term": scrapingItem.Term})
	vocabulary, err := domain.NewVocabulary("system", scrapingItem.Term)
	if err != nil {
		ctx.Logger().Error("failed to create vocabulary model", err, appcontext.Fields{"term": scrapingItem.Term})
		return err
	}

	ctx.Logger().Text("fetch vocabulary data with Datamuse")
	datamuseData, err := w.externalApiRepository.SearchTermWithDatamuse(ctx, scrapingItem.Term)
	if err != nil {
		ctx.Logger().Error("failed to fetch vocabulary data with Datamuse", err, appcontext.Fields{})
		return err
	}
	if datamuseData == nil {
		ctx.Logger().ErrorText("datamuseData is null, return")
		return apperrors.Common.BadRequest
	}

	ctx.Logger().Text("generate vocabulary examples with GPT")
	aiExamplesData, err := w.aiRepository.VocabularyExamples(ctx, scrapingItem.Term, datamuseData.PartsOfSpeech)
	if err != nil {
		ctx.Logger().Error("failed to generate vocabulary examples with GPT", err, appcontext.Fields{})
		return err
	}
	if aiExamplesData == nil {
		ctx.Logger().ErrorText("aiExamplesData is null, return")
		return apperrors.Common.BadRequest
	}

	ctx.Logger().Text("translate and analyze the examples")
	examples, err := w.analyzeExamples(ctx, *vocabulary, aiExamplesData)
	if err != nil {
		ctx.Logger().Error("failed to analyze examples", err, appcontext.Fields{})
		return err
	}

	ctx.Logger().Text("generate sound for the vocabulary")
	soundGenerationResult, err := w.ttsRepository.GenerateVocabularySound(ctx, vocabulary.Term)
	if err != nil {
		ctx.Logger().Error("failed to generate sound for the vocabulary", err, appcontext.Fields{})
		return err
	}

	ctx.Logger().Text("set vocabulary data")
	if err = w.setVocabularyData(ctx, vocabulary, datamuseData, soundGenerationResult); err != nil {
		ctx.Logger().Error("failed to set vocabulary data", err, appcontext.Fields{})
		return err
	}

	ctx.Logger().Text("persist vocabulary data to db")
	if err = w.vocabularyRepository.CreateVocabulary(ctx, *vocabulary); err != nil {
		ctx.Logger().Error("failed to insert vocabulary", err, appcontext.Fields{})
		return err
	}

	ctx.Logger().Text("persist vocabulary examples to db")
	if err = w.vocabularyExampleRepository.CreateVocabularyExamples(ctx, examples); err != nil {
		ctx.Logger().Error("failed to insert vocabulary examples", err, appcontext.Fields{})
	}

	ctx.Logger().Text("enqueue tasks")
	if err = w.enqueueTasks(ctx, vocabulary, examples); err != nil {
		ctx.Logger().Error("failed to enqueue tasks", err, appcontext.Fields{})
	}

	ctx.Logger().Text("delete scraping item")
	if err = w.vocabularyScrapingItemRepository.DeleteVocabularyScrapingItemByTerm(ctx, scrapingItem.Term); err != nil {
		ctx.Logger().Error("failed to delete scraping item", err, appcontext.Fields{})
	}

	return nil
}

func (AutoScrapingVocabularyHandler) setVocabularyData(ctx *appcontext.AppContext, vocabulary *domain.Vocabulary, datamuseData *domain.DatamuseSearchTermResult, soundGenerationResult *domain.TTSGenerateSoundResult) error {
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

func (w AutoScrapingVocabularyHandler) analyzeExamples(ctx *appcontext.AppContext, vocabulary domain.Vocabulary, aiExamples []domain.AIVocabularyExample) ([]domain.VocabularyExample, error) {
	var (
		wg     sync.WaitGroup
		result = make([]domain.VocabularyExample, len(aiExamples))
	)

	wg.Add(len(aiExamples))

	for i, aiExample := range aiExamples {
		go func(i int, e domain.AIVocabularyExample) {
			defer wg.Done()

			analysisResult, err := w.nlpRepository.AnalyzeSentence(ctx, e.Example, vocabulary.Term)
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

func (w AutoScrapingVocabularyHandler) enqueueTasks(ctx *appcontext.AppContext, vocabulary *domain.Vocabulary, examples []domain.VocabularyExample) error {
	var (
		wg         sync.WaitGroup
		totalTasks = 1 + len(examples)
	)
	wg.Add(totalTasks)

	go func() {
		defer wg.Done()

		ctx.Logger().Text("add task newVocabularyCreated")
		if err := w.queueRepository.NewVocabularyCreated(ctx, domain.QueueNewVocabularyCreatedPayload{
			Vocabulary: *vocabulary,
		}); err != nil {
			ctx.Logger().Error("failed to enqueue task", err, appcontext.Fields{})
		}
	}()

	for _, e := range examples {
		go func(example domain.VocabularyExample) {
			defer wg.Done()

			ctx.Logger().Info("add task newVocabularyExampleCreated", appcontext.Fields{"exampleID": example.ID})
			if err := w.queueRepository.NewVocabularyExampleCreated(ctx, domain.QueueNewVocabularyExampleCreatedPayload{
				Example: example,
			}); err != nil {
				ctx.Logger().Error("failed to enqueue task", err, appcontext.Fields{})
			}
		}(e)
	}

	defer wg.Wait()

	return nil
}
