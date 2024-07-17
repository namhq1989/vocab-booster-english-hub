package domain

import "github.com/namhq1989/vocab-booster-utilities/appcontext"

type CachingRepository interface {
	GetExerciseCollections(ctx *appcontext.AppContext) (*[]ExerciseCollection, error)
	SetExerciseCollections(ctx *appcontext.AppContext, collections []ExerciseCollection) error

	GetUserExerciseCollections(ctx *appcontext.AppContext, userID string) (*[]UserExerciseCollection, error)
	SetUserExerciseCollections(ctx *appcontext.AppContext, userID string, collections []UserExerciseCollection) error
	DeleteUserExerciseCollections(ctx *appcontext.AppContext, userID string) error
}
