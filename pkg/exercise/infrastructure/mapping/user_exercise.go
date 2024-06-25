package mapping

import (
	"encoding/json"

	"github.com/namhq1989/vocab-booster-english-hub/core/language"
	"github.com/namhq1989/vocab-booster-english-hub/internal/database/gen/vocab-booster/public/model"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/exercise/domain"
)

type UserExercise struct {
	Exercise model.Exercises            `alias:"e"`
	Status   model.UserExerciseStatuses `alias:"ues"`
}

type UserExerciseMapper struct{}

func (UserExerciseMapper) FromModelToDomain(ue UserExercise, lang string) (*domain.UserExercise, error) {
	var translated language.TranslatedLanguages
	if err := json.Unmarshal([]byte(ue.Exercise.Translated), &translated); err != nil {
		return nil, err
	}

	return &domain.UserExercise{
		ID:            ue.Exercise.ID,
		Level:         domain.ToExerciseLevel(ue.Exercise.Level),
		Audio:         ue.Exercise.Audio,
		Vocabulary:    ue.Exercise.Vocabulary,
		Content:       ue.Exercise.Content,
		Translated:    translated.GetLanguageValue(lang),
		CorrectAnswer: ue.Exercise.CorrectAnswer,
		Options:       ue.Exercise.Options,
		CorrectStreak: int(ue.Status.CorrectStreak),
		IsFavorite:    ue.Status.IsFavorite,
		IsMastered:    ue.Status.IsMastered,
		UpdatedAt:     ue.Status.UpdatedAt,
		NextReviewAt:  ue.Status.NextReviewAt,
	}, nil
}
