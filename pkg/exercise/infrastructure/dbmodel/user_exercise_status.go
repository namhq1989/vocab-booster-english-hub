package dbmodel

import (
	"time"

	apperrors "github.com/namhq1989/vocab-booster-english-hub/core/error"
	"github.com/namhq1989/vocab-booster-english-hub/internal/database"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/exercise/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserExerciseStatus struct {
	ID            primitive.ObjectID `bson:"_id"`
	ExerciseID    primitive.ObjectID `bson:"exerciseId"`
	UserID        primitive.ObjectID `bson:"userId"`
	CorrectStreak int                `bson:"correctStreak"`
	IsFavorite    bool               `bson:"isFavorite"`
	IsMastered    bool               `bson:"isMastered"`
	UpdatedAt     time.Time          `bson:"updatedAt"`
	NextReviewAt  time.Time          `bson:"nextReviewAt"`
}

func (m UserExerciseStatus) ToDomain() domain.UserExerciseStatus {
	return domain.UserExerciseStatus{
		ID:            m.ID.Hex(),
		ExerciseID:    m.ExerciseID.Hex(),
		UserID:        m.UserID.Hex(),
		CorrectStreak: m.CorrectStreak,
		IsFavorite:    m.IsFavorite,
		IsMastered:    m.IsMastered,
		UpdatedAt:     m.UpdatedAt,
		NextReviewAt:  m.NextReviewAt,
	}
}

func (UserExerciseStatus) FromDomain(status domain.UserExerciseStatus) (*UserExerciseStatus, error) {
	id, err := database.ObjectIDFromString(status.ID)
	if err != nil {
		return nil, apperrors.Exercise.InvalidExerciseID
	}

	eid, err := database.ObjectIDFromString(status.ExerciseID)
	if err != nil {
		return nil, apperrors.Exercise.InvalidExerciseID
	}

	uid, err := database.ObjectIDFromString(status.UserID)
	if err != nil {
		return nil, apperrors.User.InvalidUserID
	}

	return &UserExerciseStatus{
		ID:            id,
		ExerciseID:    eid,
		UserID:        uid,
		CorrectStreak: status.CorrectStreak,
		IsFavorite:    status.IsFavorite,
		IsMastered:    status.IsMastered,
		UpdatedAt:     status.UpdatedAt,
		NextReviewAt:  status.NextReviewAt,
	}, nil
}
