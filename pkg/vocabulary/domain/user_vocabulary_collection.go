package domain

import (
	"time"

	"github.com/namhq1989/vocab-booster-english-hub/core/appcontext"
	apperrors "github.com/namhq1989/vocab-booster-english-hub/core/error"
	"github.com/namhq1989/vocab-booster-english-hub/internal/database"
)

const (
	maxCollectionsPerUser = 20
)

type UserVocabularyCollectionRepository interface {
	FindUserVocabularyCollectionsByUserID(ctx *appcontext.AppContext, userID string) ([]UserVocabularyCollection, error)
	CountTotalUserVocabularyCollectionsByUserID(ctx *appcontext.AppContext, userID string) (int64, error)
	FindUserVocabularyCollectionByID(ctx *appcontext.AppContext, id string) (*UserVocabularyCollection, error)
	CreateUserVocabularyCollection(ctx *appcontext.AppContext, collection UserVocabularyCollection) error
	UpdateUserVocabularyCollection(ctx *appcontext.AppContext, collection UserVocabularyCollection) error
}

type UserVocabularyCollection struct {
	ID              string
	UserID          string
	Name            string
	NumOfVocabulary int
	CreatedAt       time.Time
}

func NewUserVocabularyCollection(userID, name string) (*UserVocabularyCollection, error) {
	if !database.IsValidID(userID) {
		return nil, apperrors.User.InvalidUserID
	}

	if name == "" {
		return nil, apperrors.Common.InvalidName
	}

	return &UserVocabularyCollection{
		ID:              database.NewStringID(),
		UserID:          userID,
		Name:            name,
		NumOfVocabulary: 0,
		CreatedAt:       time.Now(),
	}, nil
}

func (d *UserVocabularyCollection) SetName(name string) error {
	if name == "" {
		return apperrors.Common.InvalidName
	}

	d.Name = name
	return nil
}

func (d *UserVocabularyCollection) IncreaseNumOfVocabulary() {
	d.NumOfVocabulary++
}

func (d *UserVocabularyCollection) DecreaseNumOfVocabulary() {
	d.NumOfVocabulary--
}

func (d *UserVocabularyCollection) CanCreate(totalCollections int) bool {
	return totalCollections <= maxCollectionsPerUser
}
