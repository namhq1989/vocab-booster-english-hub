package worker

import (
	"fmt"
	"strconv"
	"strings"

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

	level := fmt.Sprintf("level=%s", payload.Exercise.Level)
	frequency := fmt.Sprintf("frequency=%f", payload.Exercise.Frequency)
	ctx.Logger().Info("prepared criteria with exercise data", appcontext.Fields{"level": level, "frequency": frequency})

	for _, collection := range collections {
		if collection.Criteria == "" {
			if err = w.updateStats(ctx, payload.UserExerciseStatus.UserID, collection); err != nil {
				return err
			}
			continue
		}

		for _, c := range []string{level, frequency} {
			if !w.checkCollectionCriteria(ctx, collection.Criteria, c) {
				continue
			}
			if err = w.updateStats(ctx, payload.UserExerciseStatus.UserID, collection); err != nil {
				return err
			}
		}
	}

	// delete caching data
	err = w.cachingRepository.DeleteUserExerciseCollections(ctx, payload.UserExerciseStatus.UserID)
	if err != nil {
		ctx.Logger().Error("failed to delete caching data", err, appcontext.Fields{})
	}

	return nil
}

func (UpdateUserExerciseCollectionStatsHandler) checkCollectionCriteria(ctx *appcontext.AppContext, collectionCriteria, exerciseCriteria string) bool {
	if collectionCriteria == "" {
		return false
	}

	if strings.Contains(collectionCriteria, "level") || strings.Contains(exerciseCriteria, "level") {
		return collectionCriteria == exerciseCriteria
	}

	collectionFrequency, err := strconv.ParseFloat(strings.Split(collectionCriteria, "=")[1], 64)
	if err != nil {
		ctx.Logger().Error("failed to parse collection frequency", err, appcontext.Fields{"collectionCriteria": collectionCriteria})
		return false
	}

	exerciseFrequency, err := strconv.ParseFloat(strings.Split(exerciseCriteria, "=")[1], 64)
	if err != nil {
		ctx.Logger().Error("failed to parse exercise frequency", err, appcontext.Fields{"exerciseCriteria": exerciseCriteria})
		return false
	}

	return collectionFrequency >= exerciseFrequency
}

func (w UpdateUserExerciseCollectionStatsHandler) updateStats(ctx *appcontext.AppContext, userID string, collection domain.ExerciseCollection) error {
	ctx.Logger().Info("create new user exercise collection status model", appcontext.Fields{"id": collection.ID, "criteria": collection.Criteria})
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
