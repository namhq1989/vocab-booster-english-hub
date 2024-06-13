package dbmodel

import (
	"time"

	apperrors "github.com/namhq1989/vocab-booster-english-hub/core/error"
	"github.com/namhq1989/vocab-booster-english-hub/internal/database"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type VocabularyExample struct {
	ID           primitive.ObjectID `bson:"_id"`
	VocabularyID primitive.ObjectID `bson:"vocabularyId"`
	English      string             `bson:"english"`
	Vietnamese   string             `bson:"vietnamese"`
	POS          string             `bson:"pos"`
	Definition   string             `bson:"definition"`
	Word         string             `bson:"word"`
	CreatedAt    time.Time          `bson:"createdAt"`
}

func (m VocabularyExample) ToDomain() domain.VocabularyExample {
	return domain.VocabularyExample{
		ID:           m.ID.Hex(),
		VocabularyID: m.VocabularyID.Hex(),
		English:      m.English,
		Vietnamese:   m.Vietnamese,
		POS:          domain.ToPartOfSpeech(m.POS),
		Definition:   m.Definition,
		Word:         m.Word,
		CreatedAt:    m.CreatedAt,
	}
}

func (VocabularyExample) FromDomain(example domain.VocabularyExample) (*VocabularyExample, error) {
	id, err := database.ObjectIDFromString(example.ID)
	if err != nil {
		return nil, apperrors.Common.InvalidID
	}

	vid, err := database.ObjectIDFromString(example.VocabularyID)
	if err != nil {
		return nil, apperrors.Vocabulary.InvalidVocabularyID
	}

	return &VocabularyExample{
		ID:           id,
		VocabularyID: vid,
		English:      example.English,
		Vietnamese:   example.Vietnamese,
		POS:          example.POS.String(),
		Definition:   example.Definition,
		Word:         example.Word,
		CreatedAt:    example.CreatedAt,
	}, nil
}
