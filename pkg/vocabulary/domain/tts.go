package domain

import "github.com/namhq1989/vocab-booster-utilities/appcontext"

type TTSRepository interface {
	GenerateVocabularySound(ctx *appcontext.AppContext, vocabulary string) (*TTSGenerateSoundResult, error)
	GenerateVocabularyExampleSound(ctx *appcontext.AppContext, exampleID, exampleContent string) (*TTSGenerateSoundResult, error)
}

type TTSGenerateSoundResult struct {
	FileName string
}
