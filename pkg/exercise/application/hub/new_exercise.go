package hub

import (
	"github.com/namhq1989/vocab-booster-english-hub/core/appcontext"
	"github.com/namhq1989/vocab-booster-english-hub/core/language"
	"github.com/namhq1989/vocab-booster-english-hub/internal/genproto/exercisepb"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/exercise/domain"
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
	exercise, err := domain.NewExercise(req.GetVocabularyExampleId(), req.GetLevel(), req.GetContent(), req.GetVocabulary(), req.GetCorrectAnswer(), language.TranslatedLanguages{
		Vi: req.GetTranslated().GetVi(),
	}, req.GetOptions())
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
