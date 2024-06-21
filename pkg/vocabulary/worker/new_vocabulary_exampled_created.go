package worker

import (
	"sync"

	"github.com/namhq1989/vocab-booster-english-hub/internal/utils/manipulation"

	"github.com/namhq1989/vocab-booster-english-hub/core/appcontext"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/domain"
)

type NewVocabularyExampleCreatedHandler struct {
	vocabularyRepository        domain.VocabularyRepository
	vocabularyExampleRepository domain.VocabularyExampleRepository
	queueRepository             domain.QueueRepository
	exerciseHub                 domain.ExerciseHub
	numOfExerciseOptions        int
	vocabularyBank              []string
}

func NewNewVocabularyExampleCreatedHandler(
	vocabularyRepository domain.VocabularyRepository,
	vocabularyExampleRepository domain.VocabularyExampleRepository,
	queueRepository domain.QueueRepository,
	exerciseHub domain.ExerciseHub,
) NewVocabularyExampleCreatedHandler {
	return NewVocabularyExampleCreatedHandler{
		vocabularyRepository:        vocabularyRepository,
		vocabularyExampleRepository: vocabularyExampleRepository,
		queueRepository:             queueRepository,
		exerciseHub:                 exerciseHub,
		numOfExerciseOptions:        4,
		vocabularyBank: []string{
			"accomplish",
			"accurate",
			"astonishing",
			"crucial",
			"demand",
			"emphasize",
			"generate",
			"implement",
			"justify",
			"overcome",
			"comprehensive",
			"deteriorate",
			"indispensable",
			"mitigate",
			"ambiguous",
			"noteworthy",
			"prevalent",
			"profound",
			"sustainable",
			"tentative",
		},
	}
}

func (w NewVocabularyExampleCreatedHandler) NewVocabularyExampleCreated(ctx *appcontext.AppContext, payload domain.QueueNewVocabularyExampleCreatedPayload) error {
	ctx.Logger().Info("post process for vocabulary example", appcontext.Fields{"exampleID": payload.Example.ID, "content": payload.Example.Content})

	var (
		example = payload.Example
		wg      sync.WaitGroup
	)

	wg.Add(4)

	go func() {
		ctx.Logger().Text("create audio task")

		defer wg.Done()

		if err := w.queueRepository.CreateVocabularyExampleAudio(ctx, domain.QueueCreateVocabularyExampleAudioPayload{
			Example: example,
		}); err != nil {
			ctx.Logger().Error("failed to create audio", err, appcontext.Fields{})
		}
	}()

	go func() {
		ctx.Logger().Text("create verb conjugation")

		defer wg.Done()

		if err := w.queueRepository.CreateVerbConjugation(ctx, domain.QueueCreateVerbConjugationPayload{
			Example: example,
		}); err != nil {
			ctx.Logger().Error("failed to create audio", err, appcontext.Fields{})
		}
	}()

	go func() {
		ctx.Logger().Text("add other vocabulary to scrape queue")

		defer wg.Done()

		if err := w.queueRepository.AddOtherVocabularyToScrapeQueue(ctx, domain.QueueAddOtherVocabularyToScrapeQueuePayload{
			Example: example,
		}); err != nil {
			ctx.Logger().Error("failed to create audio", err, appcontext.Fields{})
		}
	}()

	go func() {
		ctx.Logger().Text("create exercise")

		defer wg.Done()

		ctx.Logger().Text("find random vocabulary for options")
		var options = w.randomVocabularyFromBank(ctx)
		randomVocabulary, err := w.vocabularyRepository.RandomPickVocabularyForExercise(ctx, w.numOfExerciseOptions)
		if err != nil {
			ctx.Logger().Error("failed to find random vocabulary for options", err, appcontext.Fields{})
		} else {
			for index, vocabulary := range randomVocabulary {
				options[index] = vocabulary.Term
			}
		}

		// set answer into options
		options[0] = example.MainWord.Word

		ctx.Logger().Text("create new exercise")
		if err = w.exerciseHub.CreateExercise(ctx, example.ID, example.Content, example.MainWord.Base, example.MainWord.Word, example.Translated, options); err != nil {
			ctx.Logger().Error("failed to create new exercise", err, appcontext.Fields{})
		}
	}()

	wg.Wait()

	return nil
}

func (w NewVocabularyExampleCreatedHandler) randomVocabularyFromBank(ctx *appcontext.AppContext) []string {
	var (
		randomVocabulary = make([]string, 0)
		total            = len(w.vocabularyBank)
	)

	for i := 0; i < w.numOfExerciseOptions; i++ {
		rand := manipulation.RandomIntInRange(0, total-1)
		randomVocabulary = append(randomVocabulary, w.vocabularyBank[rand])
	}
	return randomVocabulary
}
