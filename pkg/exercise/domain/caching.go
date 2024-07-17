package domain

import "github.com/namhq1989/vocab-booster-utilities/appcontext"

type CachingRepository interface {
	GetExerciseCollections(ctx *appcontext.AppContext) (*[]ExerciseCollection, error)
	SetExerciseCollections(ctx *appcontext.AppContext, collections []ExerciseCollection) error
}
