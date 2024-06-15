package worker

import (
	"github.com/namhq1989/vocab-booster-english-hub/core/appcontext"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/domain"
)

type NewVocabularyCreatedHandler struct {
	vocabularyRepository domain.VocabularyRepository
}

func NewNewVocabularyCreatedHandler(
	vocabularyRepository domain.VocabularyRepository,
) NewVocabularyCreatedHandler {
	return NewVocabularyCreatedHandler{
		vocabularyRepository: vocabularyRepository,
	}
}

func (NewVocabularyCreatedHandler) NewVocabularyCreated(ctx *appcontext.AppContext, _ domain.QueueNewVocabularyCreatedPayload) error {
	ctx.Logger().Text("** DO NOTHING **")
	return nil
}
