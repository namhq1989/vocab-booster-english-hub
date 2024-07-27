package domain

import "github.com/namhq1989/vocab-booster-utilities/appcontext"

type QueueRepository interface {
	UpdateUserExerciseCollectionStats(ctx *appcontext.AppContext, payload QueueUpdateUserExerciseCollectionStatsPayload) error
	UpsertUserExerciseInteractedHistory(ctx *appcontext.AppContext, payload QueueUpsertUserExerciseInteractedHistoryPayload) error
}

type QueueUpdateUserExerciseCollectionStatsPayload struct {
	UserExerciseStatus UserExerciseStatus
	Exercise           Exercise
}

type QueueUpsertUserExerciseInteractedHistoryPayload struct {
	UserID     string
	ExerciseID string
	Timezone   string
}

type QueueUpdateExerciseCollectionStatsPayload struct{}
