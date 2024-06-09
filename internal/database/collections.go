package database

import "go.mongodb.org/mongo-driver/mongo"

var Collections = struct {
	Vocabulary string
	Sentence   string
}{
	Vocabulary: "englishHub.vocabulary",
	Sentence:   "englishHub.sentences",
}

func (db Database) GetCollection(table string) *mongo.Collection {
	return db.mongo.Collection(table)
}
