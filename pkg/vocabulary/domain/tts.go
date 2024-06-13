package domain

import "github.com/namhq1989/vocab-booster-english-hub/core/appcontext"

type TTSRepository interface {
	GenerateVocabularyPronunciationSound(ctx *appcontext.AppContext, vocabulary string) (*TTSGenerateVocabularyPronunciationSoundResult, error)
}

type TTSGenerateVocabularyPronunciationSoundResult struct {
	FileName string
}
