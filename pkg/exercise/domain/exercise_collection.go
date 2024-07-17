package domain

import (
	"time"

	"github.com/namhq1989/vocab-booster-english-hub/internal/database"
	apperrors "github.com/namhq1989/vocab-booster-english-hub/internal/utils/error"
	"github.com/namhq1989/vocab-booster-english-hub/internal/utils/manipulation"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
	"github.com/namhq1989/vocab-booster-utilities/language"
)

type ExerciseCollectionRepository interface {
	CreateExerciseCollection(ctx *appcontext.AppContext, collection ExerciseCollection) error
	UpdateExerciseCollection(ctx *appcontext.AppContext, collection ExerciseCollection) error
	CountExerciseCollections(ctx *appcontext.AppContext) (int64, error)
	FindExerciseCollections(ctx *appcontext.AppContext) ([]ExerciseCollection, error)
	FindUserExerciseCollections(ctx *appcontext.AppContext, userID string) ([]UserExerciseCollection, error)
}

type ExerciseCollection struct {
	ID                 string
	Name               string
	Slug               string
	Translated         language.TranslatedLanguages
	Criteria           string
	IsFromSystem       bool
	Order              int
	Image              string
	StatsExercises     int
	LastStatsUpdatedAt time.Time
}

func NewExerciseCollection(name string, translated language.TranslatedLanguages, criteria string, isFromSystem bool, order int, image string) (*ExerciseCollection, error) {
	if name == "" {
		return nil, apperrors.Common.InvalidName
	}

	if translated.IsEmpty() {
		return nil, apperrors.Common.InvalidLanguage
	}

	return &ExerciseCollection{
		ID:                 database.NewStringID(),
		Name:               name,
		Slug:               manipulation.Slugify(name),
		Translated:         translated,
		Criteria:           criteria,
		IsFromSystem:       isFromSystem,
		Order:              order,
		Image:              image,
		StatsExercises:     0,
		LastStatsUpdatedAt: time.Now(),
	}, nil
}

func (d *ExerciseCollection) IncreaseStatsExercises(num int) error {
	d.StatsExercises += num
	d.LastStatsUpdatedAt = time.Now()
	return nil
}
