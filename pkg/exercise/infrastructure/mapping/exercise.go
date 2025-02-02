package mapping

import (
	"github.com/goccy/go-json"
	"github.com/namhq1989/vocab-booster-english-hub/internal/database/gen/vocab-booster/public/model"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/exercise/domain"
	"github.com/namhq1989/vocab-booster-utilities/language"
)

type ExerciseMapper struct{}

func (ExerciseMapper) FromModelToDomain(exercise model.Exercises) (*domain.Exercise, error) {
	var result = &domain.Exercise{
		ID:                  exercise.ID,
		VocabularyExampleID: exercise.VocabularyExampleID,
		Level:               domain.ToExerciseLevel(exercise.Level),
		Frequency:           exercise.Frequency,
		Audio:               exercise.Audio,
		Vocabulary:          exercise.Vocabulary,
		Content:             language.Multilingual{},
		CorrectAnswer:       exercise.CorrectAnswer,
		Options:             exercise.Options,
		CreatedAt:           exercise.CreatedAt,
	}

	if err := json.Unmarshal([]byte(exercise.Content), &result.Content); err != nil {
		return nil, err
	}

	return result, nil
}

func (ExerciseMapper) FromDomainToModel(exercise domain.Exercise) (*model.Exercises, error) {
	var result = &model.Exercises{
		ID:                  exercise.ID,
		VocabularyExampleID: exercise.VocabularyExampleID,
		Level:               exercise.Level.String(),
		Frequency:           exercise.Frequency,
		Audio:               exercise.Audio,
		Vocabulary:          exercise.Vocabulary,
		Content:             "",
		CorrectAnswer:       exercise.CorrectAnswer,
		Options:             exercise.Options,
		CreatedAt:           exercise.CreatedAt,
	}

	if data, err := json.Marshal(exercise.Content); err != nil {
		return nil, err
	} else {
		result.Content = string(data)
	}

	return result, nil
}
