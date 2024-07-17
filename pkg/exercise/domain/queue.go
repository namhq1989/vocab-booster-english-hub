package domain

import "github.com/namhq1989/vocab-booster-utilities/appcontext"

type QueueRepository interface {
	UpdateUserExerciseCollectionStats(ctx *appcontext.AppContext, payload QueueUpdateUserExerciseCollectionStatsPayload) error
}

type QueueUpdateUserExerciseCollectionStatsPayload struct {
	UserExerciseStatus UserExerciseStatus
	Exercise           Exercise
}

type QueueUpdateExerciseCollectionStatsPayload struct{}
