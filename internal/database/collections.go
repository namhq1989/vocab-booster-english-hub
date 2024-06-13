package database

import "go.mongodb.org/mongo-driver/mongo"

var Collections = struct {
	Vocabulary        string
	VocabularyExample string
	Sentence          string
}{
	Vocabulary:        "englishHub.vocabulary",
	VocabularyExample: "englishHub.vocabularyExamples",
	Sentence:          "englishHub.sentences",
}

func (db Database) GetCollection(table string) *mongo.Collection {
	return db.mongo.Collection(table)
}
