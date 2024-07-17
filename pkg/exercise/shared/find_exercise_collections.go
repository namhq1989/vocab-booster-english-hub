package shared

import (
	"github.com/namhq1989/vocab-booster-english-hub/pkg/exercise/domain"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

func (s Service) FindExerciseCollections(ctx *appcontext.AppContext) ([]domain.ExerciseCollection, error) {
	ctx.Logger().Text("find in caching first")
	cachingCollections, err := s.cachingRepository.GetExerciseCollections(ctx)
	if cachingCollections != nil {
		ctx.Logger().Info("got data in caching, respond", appcontext.Fields{"numOfCollections": len(*cachingCollections)})
		return *cachingCollections, err
	} else if err != nil {
		ctx.Logger().Error("failed to find in caching", err, appcontext.Fields{})
	}

	ctx.Logger().Text("find collections in db")
	collections, err := s.exerciseCollectionRepository.FindExerciseCollections(ctx)
	if err != nil {
		ctx.Logger().Error("failed to find collections in db", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("cache collections")
	err = s.cachingRepository.SetExerciseCollections(ctx, collections)
	if err != nil {
		ctx.Logger().Error("failed to cache collections", err, appcontext.Fields{})
	}

	ctx.Logger().Text("done find exercise collections")
	return collections, nil
}
