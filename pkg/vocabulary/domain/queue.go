package domain

import "github.com/namhq1989/vocab-booster-english-hub/core/appcontext"

type QueueRepository interface {
	NewVocabularyCreated(ctx *appcontext.AppContext, payload QueueNewVocabularyCreatedPayload) error
	NewVocabularyExampleCreated(ctx *appcontext.AppContext, payload QueueNewVocabularyExampleCreatedPayload) error
	CreateVocabularyExampleAudio(ctx *appcontext.AppContext, payload QueueCreateVocabularyExampleAudioPayload) error
	CreateVerbConjugation(ctx *appcontext.AppContext, payload QueueCreateVerbConjugationPayload) error
	AddOtherVocabularyToScrapeQueue(ctx *appcontext.AppContext, payload QueueAddOtherVocabularyToScrapeQueuePayload) error
}

type QueueNewVocabularyCreatedPayload struct {
	Vocabulary Vocabulary
}

type QueueNewVocabularyExampleCreatedPayload struct {
	Example VocabularyExample
}

type QueueCreateVocabularyExampleAudioPayload struct {
	Example VocabularyExample
}

type QueueCreateVerbConjugationPayload struct {
	Example VocabularyExample
}

type QueueAddOtherVocabularyToScrapeQueuePayload struct {
	Example VocabularyExample
}
