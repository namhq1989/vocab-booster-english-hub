package infrastructure

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/namhq1989/vocab-booster-english-hub/internal/caching"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/exercise/domain"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type CachingRepository struct {
	caching caching.Operations
}

func NewCachingRepository(caching *caching.Caching) CachingRepository {
	return CachingRepository{
		caching: caching,
	}
}

func (r CachingRepository) GetExerciseCollections(ctx *appcontext.AppContext) (*[]domain.ExerciseCollection, error) {
	key := r.generateExerciseCollectionsKey()

	dataStr, err := r.caching.Get(ctx, key)
	if err != nil {
		return nil, err
	}

	var result []domain.ExerciseCollection
	if err = json.Unmarshal([]byte(dataStr), &result); err != nil {
		return nil, nil
	}

	return &result, nil
}

func (r CachingRepository) SetExerciseCollections(ctx *appcontext.AppContext, collections []domain.ExerciseCollection) error {
	key := r.generateExerciseCollectionsKey()
	r.caching.SetTTL(ctx, key, collections, r.getExerciseCollectionsTTL())
	return nil
}

func (r CachingRepository) generateExerciseCollectionsKey() string {
	return r.caching.GenerateKey("exercise", "exercise_collections")
}

func (CachingRepository) getExerciseCollectionsTTL() time.Duration {
	return 1 * time.Hour
}

func (r CachingRepository) GetUserExerciseCollections(ctx *appcontext.AppContext, userID string) (*[]domain.UserExerciseCollection, error) {
	key := r.generateUserExerciseCollectionsKey(userID)

	dataStr, err := r.caching.Get(ctx, key)
	if err != nil {
		return nil, err
	}

	var result []domain.UserExerciseCollection
	if err = json.Unmarshal([]byte(dataStr), &result); err != nil {
		return nil, nil
	}

	return &result, nil
}

func (r CachingRepository) SetUserExerciseCollections(ctx *appcontext.AppContext, userID string, collections []domain.UserExerciseCollection) error {
	key := r.generateUserExerciseCollectionsKey(userID)
	r.caching.SetTTL(ctx, key, collections, r.getUserExerciseCollectionsTTL())
	return nil
}

func (r CachingRepository) DeleteUserExerciseCollections(ctx *appcontext.AppContext, userID string) error {
	key := r.generateUserExerciseCollectionsKey(userID)
	_, err := r.caching.Del(ctx, key)
	return err
}

func (r CachingRepository) generateUserExerciseCollectionsKey(userID string) string {
	return r.caching.GenerateKey("exercise", fmt.Sprintf("%s_user_exercise_collections", userID))
}

func (CachingRepository) getUserExerciseCollectionsTTL() time.Duration {
	return 24 * time.Hour * 7
}
