package hub

import (
	"github.com/namhq1989/vocab-booster-english-hub/internal/genproto/exercisepb"
	apperrors "github.com/namhq1989/vocab-booster-english-hub/internal/utils/error"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/exercise/domain"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type ChangeExerciseFavoriteHandler struct {
	userExerciseStatusRepository domain.UserExerciseStatusRepository
}

func NewChangeExerciseFavoriteHandler(
	userExerciseStatusRepository domain.UserExerciseStatusRepository,
) ChangeExerciseFavoriteHandler {
	return ChangeExerciseFavoriteHandler{
		userExerciseStatusRepository: userExerciseStatusRepository,
	}
}

func (h ChangeExerciseFavoriteHandler) ChangeExerciseFavorite(ctx *appcontext.AppContext, req *exercisepb.ChangeExerciseFavoriteRequest) (*exercisepb.ChangeExerciseFavoriteResponse, error) {
	ctx.Logger().Info("[hub] new change exercise favorite request", appcontext.Fields{"exerciseID": req.GetExerciseId(), "userID": req.GetUserId(), "isFavorite": req.GetIsFavorite()})

	ctx.Logger().Text("find status in db")
	ues, err := h.userExerciseStatusRepository.FindUserExerciseStatus(ctx, req.GetExerciseId(), req.GetUserId())
	if err != nil {
		ctx.Logger().Error("failed to find status in db", err, appcontext.Fields{})
		return nil, err
	}
	if ues == nil {
		ctx.Logger().ErrorText("user exercise status not found in db")
		return nil, apperrors.Exercise.ExerciseNotFound
	}

	ctx.Logger().Info("check current status and new status", appcontext.Fields{"current": ues.IsFavorite, "new": req.GetIsFavorite()})
	if ues.IsFavorite == req.GetIsFavorite() {
		ctx.Logger().Text("isFavorite flag is the same with current status, respond")

		ctx.Logger().Text("done change exercise favorite request")
		return &exercisepb.ChangeExerciseFavoriteResponse{IsFavorite: ues.IsFavorite}, nil
	}

	ctx.Logger().Text("different status, set status data")
	if err = ues.SetFavorite(req.GetIsFavorite()); err != nil {
		ctx.Logger().Error("failed to set status favorite", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("update status in db")
	if err = h.userExerciseStatusRepository.UpdateUserExerciseStatus(ctx, *ues); err != nil {
		ctx.Logger().Error("failed to update status in db", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("done change exercise favorite request")
	return &exercisepb.ChangeExerciseFavoriteResponse{IsFavorite: ues.IsFavorite}, nil
}
