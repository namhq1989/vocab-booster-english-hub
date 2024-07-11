package hub

import (
	"github.com/namhq1989/vocab-booster-english-hub/internal/genproto/exercisepb"
	apperrors "github.com/namhq1989/vocab-booster-english-hub/internal/utils/error"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/exercise/domain"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type AnswerExerciseHandler struct {
	exerciseRepository           domain.ExerciseRepository
	userExerciseStatusRepository domain.UserExerciseStatusRepository
}

func NewAnswerExerciseHandler(
	exerciseRepository domain.ExerciseRepository,
	userExerciseStatusRepository domain.UserExerciseStatusRepository,
) AnswerExerciseHandler {
	return AnswerExerciseHandler{
		exerciseRepository:           exerciseRepository,
		userExerciseStatusRepository: userExerciseStatusRepository,
	}
}

func (h AnswerExerciseHandler) AnswerExercise(ctx *appcontext.AppContext, req *exercisepb.AnswerExerciseRequest) (*exercisepb.AnswerExerciseResponse, error) {
	ctx.Logger().Info("[hub] new answer exercise request", appcontext.Fields{"exerciseID": req.GetExerciseId(), "userID": req.GetUserId(), "isCorrect": req.GetIsCorrect()})

	ctx.Logger().Text("find exercise in db")
	exercise, err := h.exerciseRepository.FindExerciseByID(ctx, req.GetExerciseId())
	if err != nil {
		ctx.Logger().Error("failed to find exercise in db", err, appcontext.Fields{})
		return nil, err
	}
	if exercise == nil {
		ctx.Logger().Text("exercise not found in db")
		return nil, apperrors.Exercise.ExerciseNotFound
	}

	ctx.Logger().Text("find user exercise status in db")
	ues, err := h.userExerciseStatusRepository.FindUserExerciseStatus(ctx, req.GetExerciseId(), req.GetUserId())
	if err != nil {
		ctx.Logger().Error("failed to find user exercise status in db", err, appcontext.Fields{})
		return nil, err
	}

	if ues == nil {
		ctx.Logger().Text("this is the first time user complete this exercise, create new")
		ues, err = domain.NewUserExerciseStatus(req.GetExerciseId(), req.GetUserId())
		if err != nil {
			ctx.Logger().Error("failed to create new user exercise status", err, appcontext.Fields{})
			return nil, err
		}

		ctx.Logger().Text("persist user exercise status in db")
		if err = h.userExerciseStatusRepository.CreateUserExerciseStatus(ctx, *ues); err != nil {
			ctx.Logger().Error("failed to persist user exercise status in db", err, appcontext.Fields{})
			return nil, err
		}
	} else {
		ctx.Logger().Text("user already complete this exercise before, just update")
	}

	ctx.Logger().Text("update status data")
	_ = ues.SetResult(req.GetIsCorrect())

	ctx.Logger().Text("update user exercise status in db")
	if err = h.userExerciseStatusRepository.UpdateUserExerciseStatus(ctx, *ues); err != nil {
		ctx.Logger().Error("failed to update user exercise status in db", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("done answer exercise request")
	return &exercisepb.AnswerExerciseResponse{
		NextReviewAt: timestamppb.New(ues.NextReviewAt),
	}, nil
}
