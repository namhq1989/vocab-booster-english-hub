package queue

var TypeNames = struct {
	NewVocabularyCreated              string
	NewVocabularyExampleCreated       string
	CreateVocabularyExampleAudio      string
	CreateVerbConjugation             string
	AddOtherVocabularyToScrapingQueue string

	UpdateUserExerciseCollectionStats   string
	UpsertUserExerciseInteractedHistory string

	// cronjob
	AutoScrapingVocabulary string
	FetchWordOfTheDay      string

	UpdateExerciseCollectionStats string
}{
	NewVocabularyCreated:              "vocabulary.newVocabularyCreated",
	NewVocabularyExampleCreated:       "vocabulary.newVocabularyExampleCreated",
	CreateVocabularyExampleAudio:      "vocabulary.createVocabularyExampleAudio",
	CreateVerbConjugation:             "vocabulary.createVerbConjugation",
	AddOtherVocabularyToScrapingQueue: "vocabulary.addOtherVocabularyToScrapingQueue",

	UpdateUserExerciseCollectionStats:   "exercise.updateUserExerciseCollectionStats",
	UpsertUserExerciseInteractedHistory: "exercise.upsertUserExerciseInteractedHistory",

	// cronjob
	AutoScrapingVocabulary: "vocabulary.autoScrapingVocabulary",
	FetchWordOfTheDay:      "vocabulary.fetchWordOfTheDay",

	UpdateExerciseCollectionStats: "exercise.updateExerciseCollectionStats",
}
