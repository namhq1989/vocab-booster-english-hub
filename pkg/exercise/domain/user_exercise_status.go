package domain

import (
	"time"

	apperrors "github.com/namhq1989/vocab-booster-english-hub/internal/utils/error"

	"github.com/namhq1989/vocab-booster-english-hub/internal/database"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type UserExerciseStatusRepository interface {
	CreateUserExerciseStatus(ctx *appcontext.AppContext, status UserExerciseStatus) error
	UpdateUserExerciseStatus(ctx *appcontext.AppContext, status UserExerciseStatus) error
	FindUserExerciseStatus(ctx *appcontext.AppContext, exerciseID, userID string) (*UserExerciseStatus, error)
	CountUserReadyToReviewExercises(ctx *appcontext.AppContext, userID string) (int64, error)
	FindUserReadyToReviewExercises(ctx *appcontext.AppContext, filter UserExerciseFilter) ([]UserExercise, error)
	FindUserFavoriteExercises(ctx *appcontext.AppContext, filter UserFavoriteExerciseFilter) ([]UserExercise, error)
	FindUserStats(ctx *appcontext.AppContext, userID string) (*UserStats, error)
}

type UserExerciseStatus struct {
	ID            string
	ExerciseID    string
	UserID        string
	CorrectStreak int
	AnswerCount   int
	IsFavorite    bool
	IsMastered    bool
	UpdatedAt     time.Time
	NextReviewAt  time.Time
}

const (
	maxCorrectStreak = 5
)

func NewUserExerciseStatus(exerciseID, userID string) (*UserExerciseStatus, error) {
	if exerciseID == "" {
		return nil, apperrors.Exercise.InvalidExerciseID
	}

	if userID == "" {
		return nil, apperrors.User.InvalidUserID
	}

	return &UserExerciseStatus{
		ID:            database.NewStringID(),
		ExerciseID:    exerciseID,
		UserID:        userID,
		CorrectStreak: 0,
		IsFavorite:    false,
		IsMastered:    false,
		UpdatedAt:     time.Now(),
		NextReviewAt:  time.Now(),
	}, nil
}

func (d *UserExerciseStatus) SetResult(isCorrect bool) error {
	if isCorrect {
		d.CorrectStreak++
	} else {
		d.CorrectStreak--
	}

	if d.CorrectStreak < 0 {
		d.CorrectStreak = 0
	}

	if d.CorrectStreak >= maxCorrectStreak {
		d.IsMastered = true
		d.CorrectStreak = maxCorrectStreak
	} else {
		d.IsMastered = false
	}

	_ = d.SetNextReviewAt()
	d.AnswerCount++
	d.UpdatedAt = time.Now()
	return nil
}

func (d *UserExerciseStatus) SetFavorite(value bool) error {
	d.IsFavorite = value
	d.UpdatedAt = time.Now()
	return nil
}

func (d *UserExerciseStatus) SetNextReviewAt() error {
	nextReviewDuration := 24 * time.Hour

	switch d.CorrectStreak {
	case 0:
		nextReviewDuration = 6 * time.Hour
	case 1:
		nextReviewDuration = 24 * time.Hour
	case 2:
		nextReviewDuration = 3 * 24 * time.Hour
	case 3:
		nextReviewDuration = 7 * 24 * time.Hour
	case 4:
		nextReviewDuration = 14 * 24 * time.Hour
	case 5:
		nextReviewDuration = 30 * 24 * time.Hour
	}

	d.NextReviewAt = time.Now().Add(nextReviewDuration)
	return nil
}

//
// AGGREGATED EXERCISES
//

type UserAggregatedExercise struct {
	Date     string
	Exercise int64
}
