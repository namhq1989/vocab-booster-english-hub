package domain

import (
	"time"

	"github.com/namhq1989/vocab-booster-english-hub/core/appcontext"
	apperrors "github.com/namhq1989/vocab-booster-english-hub/core/error"
	"github.com/namhq1989/vocab-booster-english-hub/internal/database"
)

type VocabularyScrapeItemRepository interface {
	FindVocabularyScrapeItemByTerm(ctx *appcontext.AppContext, term string) (*VocabularyScrapeItem, error)
	CreateVocabularyScrapeItems(ctx *appcontext.AppContext, items []VocabularyScrapeItem) error
	DeleteVocabularyScrapeItemByTerm(ctx *appcontext.AppContext, term string) error
	RandomPickVocabularyScrapeItem(ctx *appcontext.AppContext) (*VocabularyScrapeItem, error)
}

type VocabularyScrapeItem struct {
	ID        string
	Term      string
	CreatedAt time.Time
}

func NewVocabularyScrapeItem(term string) (*VocabularyScrapeItem, error) {
	if term == "" {
		return nil, apperrors.Vocabulary.InvalidTerm
	}

	return &VocabularyScrapeItem{
		ID:        database.NewStringID(),
		Term:      term,
		CreatedAt: time.Now(),
	}, nil
}

var ScrapePosTagList = []PartOfSpeech{
	PartOfSpeechNoun,
	PartOfSpeechAdjective,
	PartOfSpeechAdverb,
}
