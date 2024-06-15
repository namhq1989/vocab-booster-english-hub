package domain

import "github.com/namhq1989/vocab-booster-english-hub/core/appcontext"

type QueueRepository interface {
	NewVocabularyCreated(ctx *appcontext.AppContext, payload QueueNewVocabularyCreatedPayload) error
}

type QueueNewVocabularyCreatedPayload struct {
	Vocabulary Vocabulary
}

type QueueNewVocabularyExampleCreatedPayload struct {
	Example VocabularyExample
}
