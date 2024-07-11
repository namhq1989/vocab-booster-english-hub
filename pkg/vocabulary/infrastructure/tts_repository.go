package infrastructure

import (
	"github.com/namhq1989/vocab-booster-english-hub/internal/tts"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/domain"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type TTSRepository struct {
	tts tts.Operations
}

func NewTTSRepository(tts tts.Operations) TTSRepository {
	return TTSRepository{
		tts: tts,
	}
}

func (r TTSRepository) GenerateVocabularySound(ctx *appcontext.AppContext, vocabulary string) (*domain.TTSGenerateSoundResult, error) {
	fileName, err := r.tts.GenerateVocabularySound(ctx, vocabulary)
	if err != nil {
		return nil, err
	}

	return &domain.TTSGenerateSoundResult{
		FileName: fileName,
	}, nil
}

func (r TTSRepository) GenerateVocabularyExampleSound(ctx *appcontext.AppContext, exampleID, exampleContent string) (*domain.TTSGenerateSoundResult, error) {
	fileName, err := r.tts.GenerateVocabularyExampleSound(ctx, exampleID, exampleContent)
	if err != nil {
		return nil, err
	}

	return &domain.TTSGenerateSoundResult{
		FileName: fileName,
	}, nil
}
