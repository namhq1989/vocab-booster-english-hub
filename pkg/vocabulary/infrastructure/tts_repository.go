package infrastructure

import (
	"github.com/namhq1989/vocab-booster-english-hub/core/appcontext"
	"github.com/namhq1989/vocab-booster-english-hub/internal/tts"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/domain"
)

type TTSRepository struct {
	tts tts.Operations
}

func NewTTSRepository(tts tts.Operations) TTSRepository {
	return TTSRepository{
		tts: tts,
	}
}

func (r TTSRepository) GenerateVocabularyPronunciationSound(ctx *appcontext.AppContext, vocabulary string) (*domain.TTSGenerateVocabularyPronunciationSoundResult, error) {
	fileName, err := r.tts.GenerateVocabularyPronunciationSound(ctx, vocabulary)
	if err != nil {
		return nil, err
	}

	return &domain.TTSGenerateVocabularyPronunciationSoundResult{
		FileName: fileName,
	}, nil
}
