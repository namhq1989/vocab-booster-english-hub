package worker

import (
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/domain"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type CreateVocabularyExampleAudioHandler struct {
	vocabularyExampleRepository domain.VocabularyExampleRepository
	ttsRepository               domain.TTSRepository
	exerciseHub                 domain.ExerciseHub
}

func NewCreateVocabularyExampleAudioHandler(
	vocabularyExampleRepository domain.VocabularyExampleRepository,
	ttsRepository domain.TTSRepository,
	exerciseHub domain.ExerciseHub,
) CreateVocabularyExampleAudioHandler {
	return CreateVocabularyExampleAudioHandler{
		vocabularyExampleRepository: vocabularyExampleRepository,
		ttsRepository:               ttsRepository,
		exerciseHub:                 exerciseHub,
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

	ctx.Logger().Text("update in db")
	if err = w.vocabularyExampleRepository.UpdateVocabularyExample(ctx, example); err != nil {
		ctx.Logger().Error("failed to update audio ib db document", err, appcontext.Fields{})
		return err
	}

	ctx.Logger().Text("update exercise audio via grpc")
	if err = w.exerciseHub.UpdateExerciseAudio(ctx, example.ID, result.FileName); err != nil {
		ctx.Logger().Error("failed to update exercise audio via grpc", err, appcontext.Fields{})
	}

	return nil
}
