package domain

import (
	"time"

	apperrors "github.com/namhq1989/vocab-booster-english-hub/core/error"
	"github.com/namhq1989/vocab-booster-english-hub/internal/database"
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
	NextReviewAt  time.Time
}

type UserExerciseFilter struct {
	UserID         string
	Level          ExerciseLevel
	Lang           string
	NumOfExercises int64
}

func NewUserExerciseFilter(userID, level, lang string) (*UserExerciseFilter, error) {
	if !database.IsValidID(userID) {
		return nil, apperrors.User.InvalidUserID
	}

	dLevel := ToExerciseLevel(level)
	if !dLevel.IsValid() {
		dLevel = ExerciseLevelUnknown
	}

	return &UserExerciseFilter{
		UserID:         userID,
		Level:          dLevel,
		Lang:           lang,
		NumOfExercises: 10,
	}, nil
}
