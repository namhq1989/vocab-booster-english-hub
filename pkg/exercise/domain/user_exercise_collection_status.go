package domain

import (
	"github.com/namhq1989/vocab-booster-english-hub/internal/database"
	apperrors "github.com/namhq1989/vocab-booster-english-hub/internal/utils/error"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type UserExerciseCollectionStatusRepository interface {
	IncreaseUserExerciseCollectionStatusStats(ctx *appcontext.AppContext, uecs UserExerciseCollectionStatus, numOfExercises int64) error
}

type UserExerciseCollectionStatus struct {
	ID                  string
	UserID              string
	CollectionID        string
	InteractedExercises int
}

func NewUserExerciseCollectionStatus(userID, collectionID string) (*UserExerciseCollectionStatus, error) {
	if !database.IsValidID(userID) {
		return nil, apperrors.User.InvalidUserID
	}

	if !database.IsValidID(collectionID) {
		return nil, apperrors.Collection.InvalidCollectionID
	}

	return &UserExerciseCollectionStatus{
		ID:                  database.NewStringID(),
		UserID:              userID,
		CollectionID:        collectionID,
		InteractedExercises: 1,
	}, nil
}
