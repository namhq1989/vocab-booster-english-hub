package shared

import (
	"slices"

	"github.com/namhq1989/vocab-booster-english-hub/pkg/exercise/domain"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

func (s Service) FindExerciseCollectionByID(ctx *appcontext.AppContext, collectionID string) (*domain.ExerciseCollection, error) {
	ctx.Logger().Text("find collections in caching first")
	cachingCollections, err := s.cachingRepository.GetExerciseCollections(ctx)
	if cachingCollections != nil {
		ctx.Logger().Text("got collections in caching, find index by id")

		collectionIndex := slices.IndexFunc(*cachingCollections, func(c domain.ExerciseCollection) bool {
			return c.ID == collectionID
		})
		if collectionIndex != -1 {
			ctx.Logger().ErrorText("collection found, respond")
			collection := (*cachingCollections)[collectionIndex]
			return &collection, nil
		} else {
			ctx.Logger().Text("collection not found in caching")
		}
	} else if err != nil {
		ctx.Logger().Error("failed to find in caching", err, appcontext.Fields{})
	}

	ctx.Logger().Text("find collection in db")
	collection, err := s.exerciseCollectionRepository.FindExerciseCollectionByID(ctx, collectionID)
	if err != nil {
		ctx.Logger().Error("failed to find collection in db", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("done find exercise collection by id")
	return collection, nil
}
