package domain

import (
	"time"

	apperrors "github.com/namhq1989/vocab-booster-english-hub/internal/utils/error"

	"github.com/namhq1989/vocab-booster-english-hub/internal/database"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

const (
	maxCollectionsPerUser = 20
)

type CollectionRepository interface {
	FindCollectionsByUserID(ctx *appcontext.AppContext, userID string) ([]Collection, error)
	CountTotalCollectionsByUserID(ctx *appcontext.AppContext, userID string) (int64, error)
	FindCollectionByID(ctx *appcontext.AppContext, id string) (*Collection, error)
	CreateCollection(ctx *appcontext.AppContext, collection Collection) error
	UpdateCollection(ctx *appcontext.AppContext, collection Collection) error
}

type Collection struct {
	ID              string
	UserID          string
	Name            string
	Description     string
	NumOfVocabulary int
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

func NewCollection(userID, name, description string) (*Collection, error) {
	if !database.IsValidID(userID) {
		return nil, apperrors.User.InvalidUserID
	}

	collection := &Collection{
		ID:              database.NewStringID(),
		UserID:          userID,
		NumOfVocabulary: 0,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	if err := collection.SetName(name); err != nil {
		return nil, err
	}

	if err := collection.SetDescription(description); err != nil {
		return nil, err
	}

	return collection, nil
}

func (d *Collection) SetName(name string) error {
	if name == "" {
		return apperrors.Common.InvalidName
	}

	d.Name = name
	d.UpdatedAt = time.Now()
	return nil
}

func (d *Collection) SetDescription(description string) error {
	if len(description) > 200 {
		return apperrors.Collection.DescriptionTooLong
	}

	d.Description = description
	d.UpdatedAt = time.Now()
	return nil
}

func (d *Collection) IncreaseNumOfVocabulary() {
	d.NumOfVocabulary++
}

func (d *Collection) DecreaseNumOfVocabulary() {
	d.NumOfVocabulary--
	if d.NumOfVocabulary < 0 {
		d.NumOfVocabulary = 0
	}
}

func (*Collection) CanCreate(totalCollections int) bool {
	return totalCollections <= maxCollectionsPerUser
}

func (d *Collection) IsOwner(userID string) bool {
	return d.UserID == userID
}
