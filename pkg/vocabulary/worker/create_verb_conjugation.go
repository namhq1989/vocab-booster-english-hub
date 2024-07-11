package worker

import (
	"errors"

	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/domain"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type CreateVerbConjugationHandler struct {
	verbConjugationRepository domain.VerbConjugationRepository
}

func NewCreateVerbConjugationHandler(
	verbConjugationRepository domain.VerbConjugationRepository,
) CreateVerbConjugationHandler {
	return CreateVerbConjugationHandler{
		verbConjugationRepository: verbConjugationRepository,
	}
}

func (w CreateVerbConjugationHandler) CreateVerbConjugation(ctx *appcontext.AppContext, payload domain.QueueCreateVerbConjugationPayload) error {
	ctx.Logger().Info("create verbs conjugation", appcontext.Fields{"exampleID": payload.Example.ID})
	var (
		verbsMap     = map[string]bool{}
		vocabularyID = payload.Example.VocabularyID
		bulkDocs     = make([]domain.VerbConjugation, 0)
	)

	ctx.Logger().Text("add all verbs to bulk documents")
	for _, verb := range payload.Example.Verbs {
		// skip if verb already exists
		if _, exists := verbsMap[verb.Base]; exists {
			continue
		}

		w.addVerbToBulk(ctx, &bulkDocs, vocabularyID, verb.Base, verb.Base, domain.VerbFormInfinitive.String())
		w.addVerbToBulk(ctx, &bulkDocs, vocabularyID, verb.Past, verb.Base, domain.VerbFormPast.String())
		w.addVerbToBulk(ctx, &bulkDocs, vocabularyID, verb.PastParticiple, verb.Base, domain.VerbFormPastParticiple.String())
		w.addVerbToBulk(ctx, &bulkDocs, vocabularyID, verb.Gerund, verb.Base, domain.VerbFormGerund.String())
		w.addVerbToBulk(ctx, &bulkDocs, vocabularyID, verb.ThirdPersonSingular, verb.Base, domain.VerbFormThirdPersonSingular.String())

		verbsMap[verb.Base] = true
	}

	if len(bulkDocs) == 0 {
		ctx.Logger().Text("no verbs to be inserted")
		return errors.New("no verbs to be inserted")
	}

	ctx.Logger().Text("insert verbs conjugation bulk docs to db")
	if err := w.verbConjugationRepository.CreateVerbConjugations(ctx, bulkDocs); err != nil {
		ctx.Logger().Error("failed to insert verbs conjugation bulk docs to db", err, appcontext.Fields{"bulkDocs": bulkDocs})
		return err
	}

	return nil
}

func (CreateVerbConjugationHandler) addVerbToBulk(ctx *appcontext.AppContext, bulkDocs *[]domain.VerbConjugation, vocabularyID, value, base, form string) {
	doc, err := domain.NewVerbConjugation(vocabularyID, value, base, form)
	if err != nil {
		ctx.Logger().Error("failed to create verb conjugation", err, appcontext.Fields{"vocabularyID": vocabularyID, "value": value, "base": base, "form": form})
	} else {
		*bulkDocs = append(*bulkDocs, *doc)
	}
}
