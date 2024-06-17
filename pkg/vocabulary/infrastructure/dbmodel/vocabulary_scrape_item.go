package dbmodel

import (
	"time"

	apperrors "github.com/namhq1989/vocab-booster-english-hub/core/error"
	"github.com/namhq1989/vocab-booster-english-hub/internal/database"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type VocabularyScrapeItem struct {
	ID        primitive.ObjectID `bson:"_id"`
	Term      string             `bson:"term"`
	CreatedAt time.Time          `bson:"createdAt"`
}

func (m VocabularyScrapeItem) ToDomain() domain.VocabularyScrapeItem {
	return domain.VocabularyScrapeItem{
		ID:        m.ID.Hex(),
		Term:      m.Term,
		CreatedAt: m.CreatedAt,
	}
}

func (VocabularyScrapeItem) FromDomain(item domain.VocabularyScrapeItem) (*VocabularyScrapeItem, error) {
	id, err := database.ObjectIDFromString(item.ID)
	if err != nil {
		return nil, apperrors.Common.InvalidID
	}

	return &VocabularyScrapeItem{
		ID:        id,
		Term:      item.Term,
		CreatedAt: item.CreatedAt,
	}, nil
}
