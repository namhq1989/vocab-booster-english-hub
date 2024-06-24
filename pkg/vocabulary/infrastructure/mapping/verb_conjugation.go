package mapping

import (
	apperrors "github.com/namhq1989/vocab-booster-english-hub/core/error"
	"github.com/namhq1989/vocab-booster-english-hub/internal/database"
	"github.com/namhq1989/vocab-booster-english-hub/internal/database/gen/vocab-booster/public/model"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/domain"
)

type VerbConjugationMapper struct{}

func (VerbConjugationMapper) FromModelToDomain(verb model.VerbConjugations) (*domain.VerbConjugation, error) {
	return &domain.VerbConjugation{
		ID:           verb.ID,
		VocabularyID: verb.VocabularyID,
		Value:        verb.Value,
		Base:         verb.Base,
		Form:         domain.ToVerbForm(verb.Form),
	}, nil
}

func (VerbConjugationMapper) FromDomainToModel(verb domain.VerbConjugation) (*model.VerbConjugations, error) {
	if !database.IsValidID(verb.ID) {
		return nil, apperrors.Common.InvalidID
	}

	if !database.IsValidID(verb.VocabularyID) {
		return nil, apperrors.Vocabulary.InvalidVocabularyID
	}

	return &model.VerbConjugations{
		ID:           verb.ID,
		VocabularyID: verb.VocabularyID,
		Value:        verb.Value,
		Base:         verb.Base,
		Form:         verb.Form.String(),
	}, nil
}
