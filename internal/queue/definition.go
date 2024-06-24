package queue

var TypeNames = struct {
	NewVocabularyCreated              string
	NewVocabularyExampleCreated       string
	CreateVocabularyExampleAudio      string
	CreateVerbConjugation             string
	AddOtherVocabularyToScrapingQueue string

	// cronjob
	AutoScrapingVocabulary string
}{
	NewVocabularyCreated:              "vocabulary.newVocabularyCreated",
	NewVocabularyExampleCreated:       "vocabulary.newVocabularyExampleCreated",
	CreateVocabularyExampleAudio:      "vocabulary.createVocabularyExampleAudio",
	CreateVerbConjugation:             "vocabulary.createVerbConjugation",
	AddOtherVocabularyToScrapingQueue: "vocabulary.addOtherVocabularyToScrapingQueue",

	// cronjob
	AutoScrapingVocabulary: "vocabulary.autoScrapingVocabulary",
}
