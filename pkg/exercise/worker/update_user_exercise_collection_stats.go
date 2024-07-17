package worker

import (
	"fmt"
	"slices"

	apperrors "github.com/namhq1989/vocab-booster-english-hub/internal/utils/error"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/exercise/domain"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type UpdateUserExerciseCollectionStatsHandler struct {
	userExerciseCollectionStatsRepository domain.UserExerciseCollectionStatusRepository
	cachingRepository                     domain.CachingRepository
	service                               domain.Service
}

func NewUpdateUserExerciseCollectionStatsHandler(
	userExerciseCollectionStatsRepository domain.UserExerciseCollectionStatusRepository,
	cachingRepository domain.CachingRepository,
	service domain.Service,
) UpdateUserExerciseCollectionStatsHandler {
	return UpdateUserExerciseCollectionStatsHandler{
		userExerciseCollectionStatsRepository: userExerciseCollectionStatsRepository,
		cachingRepository:                     cachingRepository,
		service:                               service,
	}
}

func (w UpdateUserExerciseCollectionStatsHandler) UpdateUserExerciseCollectionStats(ctx *appcontext.AppContext, payload domain.QueueUpdateUserExerciseCollectionStatsPayload) error {
	//
	// We only have the collections from system, then we just check the collection by getting from level
	// In the future, when we have more collections from other sources, we will add the logic into exercise documents
	//

	ctx.Logger().Text("find collections with service")
	collections, err := w.service.FindExerciseCollections(ctx)
	if err != nil {
		ctx.Logger().Error("failed to find collections with service", err, appcontext.Fields{})
		return err
	}

	criteria := fmt.Sprintf("level=%s", payload.Exercise.Level)

	// always add Random collection into the criteria list
	for _, c := range []string{"", criteria} {
		if err = w.updateStats(ctx, payload.UserExerciseStatus.UserID, collections, c); err != nil {
			return err
		}
	}

	// delete caching data
	err = w.cachingRepository.DeleteUserExerciseCollections(ctx, payload.UserExerciseStatus.UserID)
	if err != nil {
		ctx.Logger().Error("failed to delete caching data", err, appcontext.Fields{})
	}

	return nil
}

func (w UpdateUserExerciseCollectionStatsHandler) updateStats(ctx *appcontext.AppContext, userID string, collections []domain.ExerciseCollection, criteria string) error {
	ctx.Logger().Info("find collection with criteria", appcontext.Fields{"criteria": criteria})
	collectionIndex := slices.IndexFunc(collections, func(c domain.ExerciseCollection) bool {
		return c.Criteria == criteria
	})
	if collectionIndex == -1 {
		ctx.Logger().ErrorText("collection not found")
		return apperrors.Collection.CollectionNotFound
	}

	collection := collections[collectionIndex]

	ctx.Logger().Text("create new user exercise collection status model")
	uecs, err := domain.NewUserExerciseCollectionStatus(userID, collection.ID)
	if err != nil {
		ctx.Logger().Error("failed to create new user exercise collection status model", err, appcontext.Fields{})
		return err
	}

	ctx.Logger().Text("update in db")
	if err = w.userExerciseCollectionStatsRepository.IncreaseUserExerciseCollectionStatusStats(ctx, *uecs, 1); err != nil {
		ctx.Logger().Error("failed to update user exercise collection stats", err, appcontext.Fields{})
		return err
	}

	return nil
}
