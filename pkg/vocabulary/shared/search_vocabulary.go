package shared

import (
	"sync"

	apperrors "github.com/namhq1989/vocab-booster-english-hub/internal/utils/error"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/domain"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

func (s Service) SearchVocabulary(ctx *appcontext.AppContext, performerID, term string) (*domain.Vocabulary, []string, error) {
	ctx.Logger().Text("vocabulary not found, determine the term is valid or not")
	isValidTerm, suggestions, err := s.scraperRepository.IsTermValid(ctx, term)
	if err != nil {
		ctx.Logger().Error("failed to determine the term is valid or not", err, appcontext.Fields{})
		return nil, suggestions, err
	}
	if !isValidTerm {
		ctx.Logger().ErrorText("the term is not valid, respond")
		return nil, suggestions, nil
	}

	ctx.Logger().Text("the term is valid, create new vocabulary model")
	vocabulary, err := domain.NewVocabulary(performerID, term)
	if err != nil {
		ctx.Logger().Error("failed to create new vocabulary", err, appcontext.Fields{})
		return nil, suggestions, err
	}

	ctx.Logger().Text("fetch vocabulary data with Datamuse")
	datamuseData, err := s.externalApiRepository.SearchTermWithDatamuse(ctx, term)
	if err != nil {
		ctx.Logger().Error("failed to fetch vocabulary data with Datamuse", err, appcontext.Fields{})
		return nil, suggestions, err
	}
	if datamuseData == nil {
		ctx.Logger().ErrorText("datamuseData is null, respond")
		return nil, suggestions, apperrors.Common.BadRequest
	}

	ctx.Logger().Text("generate vocabulary examples with GPT")
	aiExamplesData, err := s.aiRepository.VocabularyExamples(ctx, term, datamuseData.PartsOfSpeech)
	if err != nil {
		ctx.Logger().Error("failed to generate vocabulary examples with GPT", err, appcontext.Fields{})
		return nil, suggestions, err
	}
	if aiExamplesData == nil {
		ctx.Logger().ErrorText("aiExamplesData is null, respond")
		return nil, suggestions, apperrors.Common.BadRequest
	}

	ctx.Logger().Text("translate and analyze the examples")
	examples, err := s.analyzeVocabularyExamples(ctx, *vocabulary, aiExamplesData)
	if err != nil {
		ctx.Logger().Error("failed to analyze examples", err, appcontext.Fields{})
		return nil, suggestions, err
	}

	ctx.Logger().Text("generate sound for the vocabulary")
	soundGenerationResult, err := s.ttsRepository.GenerateVocabularySound(ctx, term)
	if err != nil {
		ctx.Logger().Error("failed to generate sound for the vocabulary", err, appcontext.Fields{})
		return nil, suggestions, err
	}

	ctx.Logger().Text("set vocabulary data")
	if err = s.setVocabularyData(ctx, vocabulary, datamuseData, soundGenerationResult); err != nil {
		ctx.Logger().Error("failed to set vocabulary data", err, appcontext.Fields{})
		return nil, suggestions, err
	}

	ctx.Logger().Text("persist vocabulary in db")
	if err = s.vocabularyRepository.CreateVocabulary(ctx, *vocabulary); err != nil {
		ctx.Logger().Error("failed to insert vocabulary", err, appcontext.Fields{})
		return nil, suggestions, err
	}

	ctx.Logger().Text("persist vocabulary examples to db")
	if err = s.vocabularyExampleRepository.CreateVocabularyExamples(ctx, examples); err != nil {
		ctx.Logger().Error("failed to insert vocabulary examples", err, appcontext.Fields{})
	}

	ctx.Logger().Text("enqueue tasks")
	if err = s.enqueueSearchVocabularyTasks(ctx, vocabulary, examples); err != nil {
		ctx.Logger().Error("failed to enqueue tasks", err, appcontext.Fields{})
	}

	return vocabulary, suggestions, err
}

func (s Service) analyzeVocabularyExamples(ctx *appcontext.AppContext, vocabulary domain.Vocabulary, aiExamples []domain.AIVocabularyExample) ([]domain.VocabularyExample, error) {
	var (
		wg     sync.WaitGroup
		result = make([]domain.VocabularyExample, len(aiExamples))
	)

	wg.Add(len(aiExamples))

	for i, aiExample := range aiExamples {
		go func(i int, e domain.AIVocabularyExample) {
			defer wg.Done()

			analysisResult, err := s.nlpRepository.AnalyzeSentence(ctx, aiExample.Example, vocabulary.Term)
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

func (Service) setVocabularyData(ctx *appcontext.AppContext, vocabulary *domain.Vocabulary, datamuseData *domain.DatamuseSearchTermResult, soundGenerationResult *domain.TTSGenerateSoundResult) error {
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

func (s Service) enqueueSearchVocabularyTasks(ctx *appcontext.AppContext, vocabulary *domain.Vocabulary, examples []domain.VocabularyExample) error {
	var (
		wg         sync.WaitGroup
		totalTasks = 1 + len(examples)
	)
	wg.Add(totalTasks)

	go func() {
		defer wg.Done()

		ctx.Logger().Text("add task newVocabularyCreated")
		if err := s.queueRepository.NewVocabularyCreated(ctx, domain.QueueNewVocabularyCreatedPayload{
			Vocabulary: *vocabulary,
		}); err != nil {
			ctx.Logger().Error("failed to enqueue task", err, appcontext.Fields{})
		}
	}()

	for _, e := range examples {
		go func(example domain.VocabularyExample) {
			defer wg.Done()

			ctx.Logger().Info("add task newVocabularyExampleCreated", appcontext.Fields{"exampleID": example.ID})
			if err := s.queueRepository.NewVocabularyExampleCreated(ctx, domain.QueueNewVocabularyExampleCreatedPayload{
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
