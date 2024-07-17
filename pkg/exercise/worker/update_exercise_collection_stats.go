package worker

import (
	"github.com/namhq1989/vocab-booster-english-hub/pkg/exercise/domain"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type UpdateExerciseCollectionStatsHandler struct {
	exerciseRepository           domain.ExerciseRepository
	exerciseCollectionRepository domain.ExerciseCollectionRepository
	cachingRepository            domain.CachingRepository
	service                      domain.Service
}

func NewUpdateExerciseCollectionStatsHandler(
	exerciseRepository domain.ExerciseRepository,
	exerciseCollectionRepository domain.ExerciseCollectionRepository,
	cachingRepository domain.CachingRepository,
	service domain.Service,
) UpdateExerciseCollectionStatsHandler {
	return UpdateExerciseCollectionStatsHandler{
		exerciseRepository:           exerciseRepository,
		exerciseCollectionRepository: exerciseCollectionRepository,
		cachingRepository:            cachingRepository,
		service:                      service,
	}
}

func (w UpdateExerciseCollectionStatsHandler) UpdateExerciseCollectionStats(ctx *appcontext.AppContext, _ domain.QueueUpdateExerciseCollectionStatsPayload) error {
	ctx.Logger().Text("find exercise collections")
	collections, err := w.service.FindExerciseCollections(ctx)
	if err != nil {
		ctx.Logger().Error("failed to find exercise collections", err, appcontext.Fields{})
		return err
	}

	hasChanged := false

	for index, _ := range collections {
		collection := collections[index]
		if !collection.IsFromSystem {
			ctx.Logger().Info("skip this collection because it's not from system", appcontext.Fields{"name": collection.Name})
			continue
		}

		numOfNew, _ := w.exerciseRepository.CountExercisesByCriteria(ctx, collection.Criteria, collection.LastStatsUpdatedAt)
		if numOfNew > 0 {
			hasChanged = true

			_ = collection.IncreaseStatsExercises(int(numOfNew))

			ctx.Logger().Info("update collection stats", appcontext.Fields{"name": collection.Name, "numOfNew": numOfNew})
			err = w.exerciseCollectionRepository.UpdateExerciseCollection(ctx, collection)
			if err != nil {
				ctx.Logger().Error("failed to update collection stats", err, appcontext.Fields{})
				return err
			}
		} else {
			ctx.Logger().Info("no new exercise in this collection", appcontext.Fields{"name": collection.Name})
		}
	}

	if hasChanged {
		ctx.Logger().Text("has changed, set exercise collections to caching")
		if err = w.cachingRepository.SetExerciseCollections(ctx, collections); err != nil {
			ctx.Logger().Error("failed to set exercise collections to caching", err, appcontext.Fields{})
		}
	}

	return nil
}
