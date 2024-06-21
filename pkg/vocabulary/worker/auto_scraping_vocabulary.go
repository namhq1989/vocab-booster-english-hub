package worker

import (
	"errors"
	"sync"

	"github.com/namhq1989/vocab-booster-english-hub/core/language"

	"github.com/namhq1989/vocab-booster-english-hub/core/appcontext"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/domain"
)

type AutoScrapingVocabularyHandler struct {
	vocabularyRepository           domain.VocabularyRepository
	vocabularyExampleRepository    domain.VocabularyExampleRepository
	vocabularyScrapeItemRepository domain.VocabularyScrapeItemRepository
	aiRepository                   domain.AIRepository
	scraperRepository              domain.ScraperRepository
	ttsRepository                  domain.TTSRepository
	nlpRepository                  domain.NlpRepository
	queueRepository                domain.QueueRepository
}

func NewAutoScrapingVocabularyHandler(
	vocabularyRepository domain.VocabularyRepository,
	vocabularyExampleRepository domain.VocabularyExampleRepository,
	vocabularyScrapeItemRepository domain.VocabularyScrapeItemRepository,
	aiRepository domain.AIRepository,
	scraperRepository domain.ScraperRepository,
	ttsRepository domain.TTSRepository,
	nlpRepository domain.NlpRepository,
	queueRepository domain.QueueRepository,
) AutoScrapingVocabularyHandler {
	return AutoScrapingVocabularyHandler{
		vocabularyRepository:           vocabularyRepository,
		vocabularyExampleRepository:    vocabularyExampleRepository,
		vocabularyScrapeItemRepository: vocabularyScrapeItemRepository,
		aiRepository:                   aiRepository,
		scraperRepository:              scraperRepository,
		ttsRepository:                  ttsRepository,
		nlpRepository:                  nlpRepository,
		queueRepository:                queueRepository,
	}
}

func (w AutoScrapingVocabularyHandler) AutoScrapingVocabulary(ctx *appcontext.AppContext, _ domain.QueueAutoScrapingVocabularyPayload) error {
	ctx.Logger().Text("picking random scraping item in db")
	scrapingItem, err := w.vocabularyScrapeItemRepository.RandomPickVocabularyScrapeItem(ctx)
	if err != nil {
		ctx.Logger().Error("failed to pick random vocabulary scrape item", err, appcontext.Fields{})
		return err
	}

	ctx.Logger().Info("item found, create new vocabulary model", appcontext.Fields{"term": scrapingItem.Term})
	vocabulary, err := domain.NewVocabulary("system", scrapingItem.Term)
	if err != nil {
		ctx.Logger().Error("failed to create vocabulary model", err, appcontext.Fields{"term": scrapingItem.Term})
		return err
	}

	ctx.Logger().Text("fetch vocabulary data with GPT")
	aiVocabularyData, err := w.aiRepository.GetVocabularyData(ctx, vocabulary.Term)
	if err != nil {
		ctx.Logger().Error("failed to fetch vocabulary data with GPT", err, appcontext.Fields{})
		return err
	}
	if aiVocabularyData == nil {
		ctx.Logger().ErrorText("aiVocabularyData is null, respond")
		return errors.New("invalid AI data")
	}

	ctx.Logger().Text("translate and analyze the examples")
	examples, err := w.analyzeExamples(ctx, *vocabulary, aiVocabularyData.Examples)
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
	if err = w.setVocabularyData(ctx, vocabulary, aiVocabularyData, soundGenerationResult); err != nil {
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
	if err = w.vocabularyScrapeItemRepository.DeleteVocabularyScrapeItemByTerm(ctx, scrapingItem.Term); err != nil {
		ctx.Logger().Error("failed to delete scraping item", err, appcontext.Fields{})
	}

	return nil
}

func (AutoScrapingVocabularyHandler) setVocabularyData(ctx *appcontext.AppContext, vocabulary *domain.Vocabulary, aiVocabularyData *domain.AIVocabularyData, soundGenerationResult *domain.TTSGenerateSoundResult) error {
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

func (w AutoScrapingVocabularyHandler) analyzeExamples(ctx *appcontext.AppContext, vocabulary domain.Vocabulary, aiExamples []domain.AIVocabularyExample) ([]domain.VocabularyExample, error) {
	var (
		wg     sync.WaitGroup
		result = make([]domain.VocabularyExample, len(aiExamples))
	)

	wg.Add(len(aiExamples))

	for i, aiExample := range aiExamples {
		go func(i int, e domain.AIVocabularyExample) {
			defer wg.Done()

			analysisResult, err := w.nlpRepository.AnalyzeSentence(ctx, e.Example)
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

			translatedDefinition := language.TranslatedLanguages{}
			mainWordDefinitionTranslated, err := w.nlpRepository.TranslateDefinition(ctx, e.Definition)
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
