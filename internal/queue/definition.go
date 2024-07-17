package queue

var TypeNames = struct {
	NewVocabularyCreated              string
	NewVocabularyExampleCreated       string
	CreateVocabularyExampleAudio      string
	CreateVerbConjugation             string
	AddOtherVocabularyToScrapingQueue string

	UpdateUserExerciseCollectionStats string

	// cronjob
	AutoScrapingVocabulary string

	UpdateExerciseCollectionStats string
}{
	NewVocabularyCreated:              "vocabulary.newVocabularyCreated",
	NewVocabularyExampleCreated:       "vocabulary.newVocabularyExampleCreated",
	CreateVocabularyExampleAudio:      "vocabulary.createVocabularyExampleAudio",
	CreateVerbConjugation:             "vocabulary.createVerbConjugation",
	AddOtherVocabularyToScrapingQueue: "vocabulary.addOtherVocabularyToScrapingQueue",

	UpdateUserExerciseCollectionStats: "exercise.updateUserExerciseCollectionStats",

	// cronjob
	AutoScrapingVocabulary: "vocabulary.autoScrapingVocabulary",

	UpdateExerciseCollectionStats: "exercise.updateExerciseCollectionStats",
}
