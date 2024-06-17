package dbmodel

import (
	apperrors "github.com/namhq1989/vocab-booster-english-hub/core/error"
	"github.com/namhq1989/vocab-booster-english-hub/internal/database"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type VerbConjugation struct {
	ID           primitive.ObjectID `bson:"_id"`
	VocabularyID primitive.ObjectID `bson:"vocabularyId"`
	Value        string             `bson:"value"`
	Base         string             `bson:"base"`
	Form         string             `bson:"form"`
}

func (m VerbConjugation) ToDomain() domain.VerbConjugation {
	return domain.VerbConjugation{
		ID:           m.ID.Hex(),
		VocabularyID: m.VocabularyID.Hex(),
		Value:        m.Value,
		Base:         m.Base,
		Form:         domain.ToVerbForm(m.Form),
	}
}

func (VerbConjugation) FromDomain(verb domain.VerbConjugation) (*VerbConjugation, error) {
	id, err := database.ObjectIDFromString(verb.ID)
	if err != nil {
		return nil, apperrors.Common.InvalidID
	}

	vid, err := database.ObjectIDFromString(verb.VocabularyID)
	if err != nil {
		return nil, apperrors.Vocabulary.InvalidVocabularyID
	}

	return &VerbConjugation{
		ID:           id,
		VocabularyID: vid,
		Value:        verb.Value,
		Base:         verb.Base,
		Form:         verb.Form.String(),
	}, nil
}
