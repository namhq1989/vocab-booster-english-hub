package domain

import (
	"github.com/namhq1989/vocab-booster-english-hub/internal/database"
	apperrors "github.com/namhq1989/vocab-booster-english-hub/internal/utils/error"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type VerbConjugationRepository interface {
	FindVerbConjugationByValue(ctx *appcontext.AppContext, value string) (*VerbConjugation, error)
	CreateVerbConjugations(ctx *appcontext.AppContext, verbs []VerbConjugation) error
	FindVerbConjugationByVocabularyID(ctx *appcontext.AppContext, vocabularyID string) ([]VerbConjugation, error)
}

type VerbConjugation struct {
	ID           string
	VocabularyID string
	Value        string
	Base         string
	Form         VerbForm
}

func NewVerbConjugation(vocabularyID, value, base, form string) (*VerbConjugation, error) {
	if vocabularyID == "" {
		return nil, apperrors.Vocabulary.InvalidVocabularyID
	}

	if value == "" || base == "" {
		return nil, apperrors.Vocabulary.InvalidVerbConjugation
	}

	dForm := ToVerbForm(form)
	if !dForm.IsValid() {
		return nil, apperrors.Vocabulary.InvalidVerbForm
	}

	return &VerbConjugation{
		ID:           database.NewStringID(),
		VocabularyID: vocabularyID,
		Value:        value,
		Base:         base,
		Form:         dForm,
	}, nil
}
