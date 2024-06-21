package dbmodel

import (
	"time"

	apperrors "github.com/namhq1989/vocab-booster-english-hub/core/error"
	"github.com/namhq1989/vocab-booster-english-hub/core/language"
	"github.com/namhq1989/vocab-booster-english-hub/internal/database"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/exercise/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Exercise struct {
	ID                  primitive.ObjectID           `bson:"_id"`
	VocabularyExampleID primitive.ObjectID           `bson:"vocabularyExampleId"`
	Content             string                       `bson:"content"`
	Translated          language.TranslatedLanguages `bson:"translated"`
	Vocabulary          string                       `bson:"vocabulary"`
	CorrectAnswer       string                       `bson:"correctAnswer"`
	Options             []string                     `bson:"options"`
	CreatedAt           time.Time                    `bson:"createdAt"`
}

func (m Exercise) ToDomain() domain.Exercise {
	return domain.Exercise{
		ID:                  m.ID.Hex(),
		VocabularyExampleID: m.VocabularyExampleID.Hex(),
		Content:             m.Content,
		Translated:          m.Translated,
		Vocabulary:          m.Vocabulary,
		CorrectAnswer:       m.CorrectAnswer,
		Options:             m.Options,
		CreatedAt:           m.CreatedAt,
	}
}

func (Exercise) FromDomain(exercise domain.Exercise) (*Exercise, error) {
	id, err := database.ObjectIDFromString(exercise.ID)
	if err != nil {
		return nil, apperrors.Exercise.InvalidExerciseID
	}

	vid, err := database.ObjectIDFromString(exercise.VocabularyExampleID)
	if err != nil {
		return nil, apperrors.Exercise.InvalidVocabularyExampleID
	}

	return &Exercise{
		ID:                  id,
		VocabularyExampleID: vid,
		Content:             exercise.Content,
		Translated:          exercise.Translated,
		Vocabulary:          exercise.Vocabulary,
		CorrectAnswer:       exercise.CorrectAnswer,
		Options:             exercise.Options,
		CreatedAt:           exercise.CreatedAt,
	}, nil
}
