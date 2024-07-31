package domain

import "github.com/namhq1989/vocab-booster-utilities/appcontext"

type QueueRepository interface {
	NewVocabularyCreated(ctx *appcontext.AppContext, payload QueueNewVocabularyCreatedPayload) error
	NewVocabularyExampleCreated(ctx *appcontext.AppContext, payload QueueNewVocabularyExampleCreatedPayload) error
	CreateVocabularyExampleAudio(ctx *appcontext.AppContext, payload QueueCreateVocabularyExampleAudioPayload) error
	CreateVerbConjugation(ctx *appcontext.AppContext, payload QueueCreateVerbConjugationPayload) error
	AddOtherVocabularyToScrapingQueue(ctx *appcontext.AppContext, payload QueueAddOtherVocabularyToScrapingQueuePayload) error
}

type QueueNewVocabularyCreatedPayload struct {
	Vocabulary Vocabulary
}

type QueueNewVocabularyExampleCreatedPayload struct {
	Vocabulary Vocabulary
	Example    VocabularyExample
}

type QueueCreateVocabularyExampleAudioPayload struct {
	Example VocabularyExample
}

type QueueCreateVerbConjugationPayload struct {
	Example VocabularyExample
}

type QueueAddOtherVocabularyToScrapingQueuePayload struct {
	Example VocabularyExample
}

type QueueAutoScrapingVocabularyPayload struct{}
type QueueFetchWordOfTheDayPayload struct{}
