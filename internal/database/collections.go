package database

import "go.mongodb.org/mongo-driver/mongo"

var Collections = struct {
	Vocabulary           string
	VocabularyExample    string
	VocabularyScrapeItem string
	VerbConjugation      string

	Exercise string
}{
	Vocabulary:           "englishHub.vocabulary",
	VocabularyExample:    "englishHub.vocabularyExamples",
	VocabularyScrapeItem: "englishHub.vocabularyScrapeItems",
	VerbConjugation:      "englishHub.verbConjugations",

	Exercise: "englishHub.exercises",
}

func (db Database) GetCollection(table string) *mongo.Collection {
	return db.mongo.Collection(table)
}
