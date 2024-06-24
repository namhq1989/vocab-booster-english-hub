package domain

import (
	"time"

	"github.com/namhq1989/vocab-booster-english-hub/core/appcontext"
	apperrors "github.com/namhq1989/vocab-booster-english-hub/core/error"
	"github.com/namhq1989/vocab-booster-english-hub/internal/database"
)

type VocabularyScrapingItemRepository interface {
	FindVocabularyScrapingItemByTerm(ctx *appcontext.AppContext, term string) (*VocabularyScrapingItem, error)
	CreateVocabularyScrapingItems(ctx *appcontext.AppContext, items []VocabularyScrapingItem) error
	DeleteVocabularyScrapingItemByTerm(ctx *appcontext.AppContext, term string) error
	RandomPickVocabularyScrapingItem(ctx *appcontext.AppContext) (*VocabularyScrapingItem, error)
}

type VocabularyScrapingItem struct {
	ID        string
	Term      string
	CreatedAt time.Time
}

func NewVocabularyScrapingItem(term string) (*VocabularyScrapingItem, error) {
	if term == "" {
		return nil, apperrors.Vocabulary.InvalidTerm
	}

	return &VocabularyScrapingItem{
		ID:        database.NewStringID(),
		Term:      term,
		CreatedAt: time.Now(),
	}, nil
}

var ScrapingPosTagList = []PartOfSpeech{
	PartOfSpeechNoun,
	PartOfSpeechAdjective,
	PartOfSpeechAdverb,
}
