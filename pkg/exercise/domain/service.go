package domain

import "github.com/namhq1989/vocab-booster-utilities/appcontext"

type Service interface {
	FindExerciseCollections(ctx *appcontext.AppContext) ([]ExerciseCollection, error)
}
