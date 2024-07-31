package domain

import (
	"time"

	"github.com/namhq1989/vocab-booster-english-hub/internal/database"
	apperrors "github.com/namhq1989/vocab-booster-english-hub/internal/utils/error"
	"github.com/namhq1989/vocab-booster-english-hub/internal/utils/manipulation"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
	"github.com/namhq1989/vocab-booster-utilities/language"
)

type WordOfTheDayRepository interface {
	FindWordOfTheDay(ctx *appcontext.AppContext) (*WordOfTheDay, error)
	CreateWordOfTheDay(ctx *appcontext.AppContext, word WordOfTheDay) error
}

type WordOfTheDay struct {
	ID           string
	VocabularyID string
	Country      string
	Information  language.Multilingual
	Date         time.Time
}

func NewWordOfTheDay(vocabularyID, country string, information language.Multilingual) (*WordOfTheDay, error) {
	if vocabularyID == "" {
		return nil, apperrors.Vocabulary.InvalidVocabularyID
	}

	return &WordOfTheDay{
		ID:           database.NewStringID(),
		VocabularyID: vocabularyID,
		Country:      country,
		Information:  information,
		Date:         manipulation.NowUTC(),
	}, nil
}
