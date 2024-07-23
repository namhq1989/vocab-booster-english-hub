package hub

import (
	"github.com/namhq1989/vocab-booster-english-hub/internal/genproto/exercisepb"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/exercise/domain"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/exercise/dto"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type NewExerciseHandler struct {
	exerciseRepository domain.ExerciseRepository
}

func NewNewExerciseHandler(exerciseRepository domain.ExerciseRepository) NewExerciseHandler {
	return NewExerciseHandler{
		exerciseRepository: exerciseRepository,
	}
}

func (h NewExerciseHandler) NewExercise(ctx *appcontext.AppContext, req *exercisepb.NewExerciseRequest) (*exercisepb.NewExerciseResponse, error) {
	ctx.Logger().Info("[hub] new new exercise request", appcontext.Fields{
		"exampleID": req.GetVocabularyExampleId(),
		"level":     req.GetLevel(), "content": req.GetContent(),
		"correctAnswer": req.GetCorrectAnswer(), "options": req.GetOptions(),
	})

	ctx.Logger().Text("create new exercise model")
	exercise, err := domain.NewExercise(req.GetVocabularyExampleId(), req.GetLevel(), dto.ConvertGrpcDataToMultilingual(req.GetContent()), req.GetVocabulary(), req.GetCorrectAnswer(), req.GetOptions())
	if err != nil {
		ctx.Logger().Error("failed to create new exercise model", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("persist exercise in db")
	if err = h.exerciseRepository.CreateExercise(ctx, *exercise); err != nil {
		ctx.Logger().Error("failed to persist exercise in db", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("done new exercise request")
	return &exercisepb.NewExerciseResponse{
		Id: exercise.ID,
	}, nil
}
