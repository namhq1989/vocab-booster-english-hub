package domain

import (
	"time"

	"github.com/namhq1989/vocab-booster-english-hub/internal/database"
	apperrors "github.com/namhq1989/vocab-booster-english-hub/internal/utils/error"
	"github.com/namhq1989/vocab-booster-english-hub/internal/utils/manipulation"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type UserExerciseInteractedHistoryRepository interface {
	UpsertUserExerciseInteractedHistory(ctx *appcontext.AppContext, history UserExerciseInteractedHistory) error
	AggregateUserExercisesInTimeRange(ctx *appcontext.AppContext, userID, timezone string, from, to time.Time) ([]UserAggregatedExercise, error)
}

type UserExerciseInteractedHistory struct {
	ID         string
	ExerciseID string
	UserID     string
	Date       time.Time
}

func NewUserExerciseInteractedHistory(userID, exerciseID, tz string) (*UserExerciseInteractedHistory, error) {
	if !database.IsValidID(userID) {
		return nil, apperrors.User.InvalidUserID
	}

	if !database.IsValidID(exerciseID) {
		return nil, apperrors.Exercise.InvalidExerciseID
	}

	return &UserExerciseInteractedHistory{
		ID:         database.NewStringID(),
		ExerciseID: exerciseID,
		UserID:     userID,
		Date:       manipulation.StartOfToday(tz),
	}, nil
}
