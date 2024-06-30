package domain

import (
	"time"

	"github.com/namhq1989/vocab-booster-english-hub/core/appcontext"
	apperrors "github.com/namhq1989/vocab-booster-english-hub/core/error"
	"github.com/namhq1989/vocab-booster-english-hub/internal/database"
)

type CommunitySentenceLikeRepository interface {
	FindCommunitySentenceLike(ctx *appcontext.AppContext, userID, sentenceID string) (*CommunitySentenceLike, error)
	CreateCommunitySentenceLike(ctx *appcontext.AppContext, like CommunitySentenceLike) error
	DeleteCommunitySentenceLike(ctx *appcontext.AppContext, like CommunitySentenceLike) error
}

type CommunitySentenceLike struct {
	ID         string
	UserID     string
	SentenceID string
	CreatedAt  time.Time
}

func NewCommunitySentenceLike(userID, sentenceID string) (*CommunitySentenceLike, error) {
	if !database.IsValidID(userID) {
		return nil, apperrors.User.InvalidUserID
	}

	if !database.IsValidID(sentenceID) {
		return nil, apperrors.Vocabulary.InvalidSentence
	}

	return &CommunitySentenceLike{
		ID:         database.NewStringID(),
		UserID:     userID,
		SentenceID: sentenceID,
		CreatedAt:  time.Now(),
	}, nil
}
