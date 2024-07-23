package domain

import (
	"time"

	"github.com/namhq1989/vocab-booster-english-hub/internal/database"
	apperrors "github.com/namhq1989/vocab-booster-english-hub/internal/utils/error"
	"github.com/namhq1989/vocab-booster-english-hub/internal/utils/pagetoken"
)

type UserExercise struct {
	ID            string
	Level         ExerciseLevel
	Audio         string
	Vocabulary    string
	Content       string
	Translated    string
	CorrectAnswer string
	Options       []string
	CorrectStreak int
	IsFavorite    bool
	IsMastered    bool
	UpdatedAt     time.Time
	NextReviewAt  time.Time
}

type UserExerciseFilter struct {
	UserID             string
	CollectionCriteria string
	Lang               string
	Timezone           string
	NumOfExercises     int64
}

func NewUserExerciseFilter(userID, collectionCriteria, lang, timezone string) (*UserExerciseFilter, error) {
	if !database.IsValidID(userID) {
		return nil, apperrors.User.InvalidUserID
	}

	return &UserExerciseFilter{
		UserID:             userID,
		CollectionCriteria: collectionCriteria,
		Lang:               lang,
		Timezone:           timezone,
		NumOfExercises:     10,
	}, nil
}

type UserFavoriteExerciseFilter struct {
	UserID         string
	Lang           string
	Timestamp      time.Time
	NumOfExercises int64
}

func NewUserFavoriteExerciseFilter(userID, lang, pageToken string) (*UserFavoriteExerciseFilter, error) {
	if !database.IsValidID(userID) {
		return nil, apperrors.User.InvalidUserID
	}

	pt := pagetoken.Decode(pageToken)
	return &UserFavoriteExerciseFilter{
		UserID:         userID,
		Lang:           lang,
		Timestamp:      pt.Timestamp,
		NumOfExercises: 10,
	}, nil
}
