package queue

var TypeNames = struct {
	NewVocabularyCreated            string
	NewVocabularyExampleCreated     string
	CreateVocabularyExampleAudio    string
	CreateVerbConjugation           string
	AddOtherVocabularyToScrapeQueue string
}{
	NewVocabularyCreated:            "vocabulary.newVocabularyCreated",
	NewVocabularyExampleCreated:     "vocabulary.newVocabularyExampleCreated",
	CreateVocabularyExampleAudio:    "vocabulary.createVocabularyExampleAudio",
	CreateVerbConjugation:           "vocabulary.createVerbConjugation",
	AddOtherVocabularyToScrapeQueue: "vocabulary.addOtherVocabularyToScrapeQueue",
}
