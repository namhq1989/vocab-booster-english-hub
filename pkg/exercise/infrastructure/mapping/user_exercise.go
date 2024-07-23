package mapping

import (
	"encoding/json"

	"github.com/namhq1989/vocab-booster-english-hub/internal/database/gen/vocab-booster/public/model"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/exercise/domain"
	"github.com/namhq1989/vocab-booster-utilities/language"
)

type UserExercise struct {
	Exercise model.Exercises            `alias:"e"`
	Status   model.UserExerciseStatuses `alias:"ues"`
}

type UserExerciseMapper struct{}

func (UserExerciseMapper) FromModelToDomain(ue UserExercise, lang string) (*domain.UserExercise, error) {
	var multilingual language.Multilingual
	if err := json.Unmarshal([]byte(ue.Exercise.Content), &multilingual); err != nil {
		return nil, err
	}

	return &domain.UserExercise{
		ID:            ue.Exercise.ID,
		Level:         domain.ToExerciseLevel(ue.Exercise.Level),
		Audio:         ue.Exercise.Audio,
		Vocabulary:    ue.Exercise.Vocabulary,
		Content:       multilingual.GetLocalized(lang),
		CorrectAnswer: ue.Exercise.CorrectAnswer,
		Options:       ue.Exercise.Options,
		CorrectStreak: int(ue.Status.CorrectStreak),
		IsFavorite:    ue.Status.IsFavorite,
		IsMastered:    ue.Status.IsMastered,
		UpdatedAt:     ue.Status.UpdatedAt,
		NextReviewAt:  ue.Status.NextReviewAt,
	}, nil
}
