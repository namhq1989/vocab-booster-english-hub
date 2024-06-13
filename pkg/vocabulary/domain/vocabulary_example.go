package domain

import (
	"time"

	"github.com/namhq1989/vocab-booster-english-hub/core/appcontext"
	apperrors "github.com/namhq1989/vocab-booster-english-hub/core/error"
	"github.com/namhq1989/vocab-booster-english-hub/internal/database"
)

type VocabularyExampleRepository interface {
	FindVocabularyExamplesByVocabularyID(ctx *appcontext.AppContext, vocabularyID string) ([]VocabularyExample, error)
	CreateVocabularyExamples(ctx *appcontext.AppContext, examples []VocabularyExample) error
}

type VocabularyExample struct {
	ID           string
	VocabularyID string
	English      string
	Vietnamese   string
	POS          PartOfSpeech
	Definition   string
	Word         string
	CreatedAt    time.Time
}

func NewVocabularyExample(vocabularyID, english, vietnamese, pos, definition, word string) (*VocabularyExample, error) {
	if vocabularyID == "" {
		return nil, apperrors.Vocabulary.InvalidVocabularyID
	}

	if english == "" || vietnamese == "" {
		return nil, apperrors.Vocabulary.InvalidExampleLanguage
	}

	dPOS := ToPartOfSpeech(pos)
	if !dPOS.IsValid() {
		return nil, apperrors.Vocabulary.InvalidPartOfSpeech
	}

	return &VocabularyExample{
		ID:           database.NewStringID(),
		VocabularyID: vocabularyID,
		English:      english,
		Vietnamese:   vietnamese,
		POS:          dPOS,
		Definition:   definition,
		Word:         word,
		CreatedAt:    time.Now(),
	}, nil
}
