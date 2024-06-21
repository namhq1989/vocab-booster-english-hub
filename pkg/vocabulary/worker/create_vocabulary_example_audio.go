package worker

import (
	"github.com/namhq1989/vocab-booster-english-hub/core/appcontext"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/domain"
)

type CreateVocabularyExampleAudioHandler struct {
	vocabularyExampleRepository domain.VocabularyExampleRepository
	ttsRepository               domain.TTSRepository
}

func NewCreateVocabularyExampleAudioHandler(
	vocabularyExampleRepository domain.VocabularyExampleRepository,
	ttsRepository domain.TTSRepository,
) CreateVocabularyExampleAudioHandler {
	return CreateVocabularyExampleAudioHandler{
		vocabularyExampleRepository: vocabularyExampleRepository,
		ttsRepository:               ttsRepository,
	}
}

func (w CreateVocabularyExampleAudioHandler) CreateVocabularyExampleAudio(ctx *appcontext.AppContext, payload domain.QueueCreateVocabularyExampleAudioPayload) error {
	ctx.Logger().Info("generate sound for vocabulary example", appcontext.Fields{"exampleID": payload.Example.ID, "content": payload.Example.Content})
	example := payload.Example

	result, err := w.ttsRepository.GenerateVocabularyExampleSound(ctx, example.ID, example.Content)
	if err != nil {
		ctx.Logger().Error("failed to generate sound for vocabulary example", err, appcontext.Fields{})
		return err
	}

	ctx.Logger().Info("set audio data", appcontext.Fields{"fileName": result.FileName})
	if err = example.SetAudio(result.FileName); err != nil {
		ctx.Logger().Error("failed to update audio to db document", err, appcontext.Fields{})
		return err
	}

	ctx.Logger().Text("update to db")
	if err = w.vocabularyExampleRepository.UpdateVocabularyExample(ctx, example); err != nil {
		ctx.Logger().Error("failed to update audio to db document", err, appcontext.Fields{})
		return err
	}

	return nil
}
