package hub

import (
	"github.com/namhq1989/vocab-booster-english-hub/internal/genproto/exercisepb"
	apperrors "github.com/namhq1989/vocab-booster-english-hub/internal/utils/error"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/exercise/domain"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type UpdateExerciseAudioHandler struct {
	exerciseRepository domain.ExerciseRepository
}

func NewUpdateExerciseAudioHandler(exerciseRepository domain.ExerciseRepository) UpdateExerciseAudioHandler {
	return UpdateExerciseAudioHandler{
		exerciseRepository: exerciseRepository,
	}
}

func (h UpdateExerciseAudioHandler) UpdateExerciseAudio(ctx *appcontext.AppContext, req *exercisepb.UpdateExerciseAudioRequest) (*exercisepb.UpdateExerciseAudioResponse, error) {
	ctx.Logger().Info("[hub] new update exercise audio request", appcontext.Fields{"vocabularyExampleID": req.GetVocabularyExampleId(), "audio": req.GetAudio()})

	ctx.Logger().Text("find exercise in db by example id")
	exercise, err := h.exerciseRepository.FindExerciseByVocabularyExampleID(ctx, req.GetVocabularyExampleId())
	if err != nil {
		ctx.Logger().Error("failed to find exercise in db", err, appcontext.Fields{})
		return nil, err
	}
	if exercise == nil {
		ctx.Logger().Text("exercise not found in db")
		return nil, apperrors.Exercise.ExerciseNotFound
	}

	ctx.Logger().Text("set audio")
	if err = exercise.SetAudio(req.GetAudio()); err != nil {
		ctx.Logger().Error("failed to set audio", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("update exercise in db")
	if err = h.exerciseRepository.UpdateExercise(ctx, *exercise); err != nil {
		ctx.Logger().Error("failed to update exercise in db", err, appcontext.Fields{
			"exercise": exercise,
		})
		return nil, err
	}

	ctx.Logger().Text("done update exercise audio request")
	return &exercisepb.UpdateExerciseAudioResponse{}, nil
}
