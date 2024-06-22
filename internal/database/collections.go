package database

import "go.mongodb.org/mongo-driver/mongo"

var Collections = struct {
	Vocabulary           string
	VocabularyExample    string
	VocabularyScrapeItem string
	VerbConjugation      string

	Exercise           string
	UserExerciseStatus string
}{
	Vocabulary:           "englishHub.vocabulary",
	VocabularyExample:    "englishHub.vocabularyExamples",
	VocabularyScrapeItem: "englishHub.vocabularyScrapeItems",
	VerbConjugation:      "englishHub.verbConjugations",

	Exercise:           "englishHub.exercises",
	UserExerciseStatus: "englishHub.userExerciseStatuses",
}

func (db Database) GetCollection(table string) *mongo.Collection {
	return db.mongo.Collection(table)
}
