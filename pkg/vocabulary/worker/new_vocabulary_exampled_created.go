package worker

import (
	"sync"

	"github.com/namhq1989/vocab-booster-english-hub/core/appcontext"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/domain"
)

type NewVocabularyExampleCreatedHandler struct {
	vocabularyExampleRepository domain.VocabularyExampleRepository
	queueRepository             domain.QueueRepository
}

func NewNewVocabularyExampleCreatedHandler(
	vocabularyExampleRepository domain.VocabularyExampleRepository,
	queueRepository domain.QueueRepository,
) NewVocabularyExampleCreatedHandler {
	return NewVocabularyExampleCreatedHandler{
		vocabularyExampleRepository: vocabularyExampleRepository,
		queueRepository:             queueRepository,
	}
}

func (w NewVocabularyExampleCreatedHandler) NewVocabularyExampleCreated(ctx *appcontext.AppContext, payload domain.QueueNewVocabularyExampleCreatedPayload) error {
	ctx.Logger().Info("post process for vocabulary example", appcontext.Fields{"exampleID": payload.Example.ID, "content": payload.Example.FromLang})

	var (
		wg sync.WaitGroup
	)

	wg.Add(3)

	go func() {
		ctx.Logger().Text("create audio task")

		defer wg.Done()

		if err := w.queueRepository.CreateVocabularyExampleAudio(ctx, domain.QueueCreateVocabularyExampleAudioPayload{
			Example: payload.Example,
		}); err != nil {
			ctx.Logger().Error("failed to create audio", err, appcontext.Fields{})
		}
	}()

	go func() {
		ctx.Logger().Text("create verb conjugation")

		defer wg.Done()

		if err := w.queueRepository.CreateVerbConjugation(ctx, domain.QueueCreateVerbConjugationPayload{
			Example: payload.Example,
		}); err != nil {
			ctx.Logger().Error("failed to create audio", err, appcontext.Fields{})
		}
	}()

	go func() {
		ctx.Logger().Text("add other vocabulary to scrape queue")

		defer wg.Done()

		if err := w.queueRepository.AddOtherVocabularyToScrapeQueue(ctx, domain.QueueAddOtherVocabularyToScrapeQueuePayload{
			Example: payload.Example,
		}); err != nil {
			ctx.Logger().Error("failed to create audio", err, appcontext.Fields{})
		}
	}()

	wg.Wait()

	return nil
}
