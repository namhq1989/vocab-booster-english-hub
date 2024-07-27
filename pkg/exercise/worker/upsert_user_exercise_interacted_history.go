package worker

import (
	"github.com/namhq1989/vocab-booster-english-hub/pkg/exercise/domain"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type UpsertUserExerciseInteractedHistoryHandler struct {
	userExerciseInteractedHistoryRepository domain.UserExerciseInteractedHistoryRepository
}

func NewUpsertUserExerciseInteractedHistoryHandler(userExerciseInteractedHistoryRepository domain.UserExerciseInteractedHistoryRepository) UpsertUserExerciseInteractedHistoryHandler {
	return UpsertUserExerciseInteractedHistoryHandler{
		userExerciseInteractedHistoryRepository: userExerciseInteractedHistoryRepository,
	}
}

func (w UpsertUserExerciseInteractedHistoryHandler) UpsertUserExerciseInteractedHistory(ctx *appcontext.AppContext, payload domain.QueueUpsertUserExerciseInteractedHistoryPayload) error {
	ctx.Logger().Text("create new model")
	history, err := domain.NewUserExerciseInteractedHistory(payload.UserID, payload.ExerciseID, payload.Timezone)
	if err != nil {
		ctx.Logger().Error("failed to create new model", err, appcontext.Fields{})
		return err
	}

	ctx.Logger().Text("persist in db")
	if err = w.userExerciseInteractedHistoryRepository.UpsertUserExerciseInteractedHistory(ctx, *history); err != nil {
		ctx.Logger().Error("failed to persist in db", err, appcontext.Fields{})
		return err
	}

	return nil
}
