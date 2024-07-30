package domain

import (
	"time"

	"github.com/namhq1989/vocab-booster-english-hub/internal/database"
	apperrors "github.com/namhq1989/vocab-booster-english-hub/internal/utils/error"
	"github.com/namhq1989/vocab-booster-english-hub/internal/utils/manipulation"
	"github.com/namhq1989/vocab-booster-english-hub/internal/utils/pagetoken"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type UserBookmarkedVocabularyRepository interface {
	FindBookmarkedVocabulary(ctx *appcontext.AppContext, userID, vocabularyID string) (*UserBookmarkedVocabulary, error)
	FindBookmarkedVocabulariesByUserID(ctx *appcontext.AppContext, userID string, filter UserBookmarkedVocabularyFilter) ([]ExtendedUserBookmarkedVocabulary, error)
	CreateUserBookmarkedVocabulary(ctx *appcontext.AppContext, ubv UserBookmarkedVocabulary) error
	DeleteUserBookmarkedVocabulary(ctx *appcontext.AppContext, ubv UserBookmarkedVocabulary) error
}

type UserBookmarkedVocabulary struct {
	ID           string
	UserID       string
	VocabularyID string
	BookmarkedAt time.Time
}

func NewUserBookmarkedVocabulary(userID, vocabularyID string) (*UserBookmarkedVocabulary, error) {
	if userID == "" {
		return nil, apperrors.User.InvalidUserID
	}

	if vocabularyID == "" {
		return nil, apperrors.Vocabulary.InvalidVocabularyID
	}

	return &UserBookmarkedVocabulary{
		ID:           database.NewStringID(),
		UserID:       userID,
		VocabularyID: vocabularyID,
		BookmarkedAt: manipulation.NowUTC(),
	}, nil
}

type UserBookmarkedVocabularyFilter struct {
	Timestamp time.Time
	Limit     int64
}

func NewUserBookmarkedVocabularyFilter(pageToken string) (*UserBookmarkedVocabularyFilter, error) {
	pt := pagetoken.Decode(pageToken)
	return &UserBookmarkedVocabularyFilter{
		Timestamp: pt.Timestamp,
		Limit:     10,
	}, nil
}

type ExtendedUserBookmarkedVocabulary struct {
	Vocabulary   Vocabulary
	BookmarkedAt time.Time
}
