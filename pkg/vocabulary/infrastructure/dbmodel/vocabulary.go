package dbmodel

import (
	"time"

	apperrors "github.com/namhq1989/vocab-booster-english-hub/core/error"
	"github.com/namhq1989/vocab-booster-english-hub/internal/database"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Vocabulary struct {
	ID        primitive.ObjectID `bson:"_id"`
	AuthorID  string             `bson:"authorId"`
	Term      string             `bson:"term"`
	Data      VocabularyData     `bson:"data"`
	CreatedAt time.Time          `bson:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt"`
}

type VocabularyData struct {
	PartsOfSpeech []string `bson:"partsOfSpeech"`
	IPA           string   `bson:"ipa"`
	AudioName     string   `bson:"audioName"`
	Synonyms      []string `bson:"synonyms"`
	Antonyms      []string `bson:"antonyms"`
}

type VocabularyDefinition struct {
	POS        string
	English    string
	Vietnamese string
}

func (m Vocabulary) ToDomain() domain.Vocabulary {
	partsOfSpeech := make([]domain.PartOfSpeech, 0)
	for _, pos := range m.Data.PartsOfSpeech {
		partsOfSpeech = append(partsOfSpeech, domain.ToPartOfSpeech(pos))
	}

	return domain.Vocabulary{
		ID:       m.ID.Hex(),
		AuthorID: m.AuthorID,
		Term:     m.Term,
		Data: domain.VocabularyData{
			PartsOfSpeech: partsOfSpeech,
			IPA:           m.Data.IPA,
			AudioName:     m.Data.AudioName,
			Synonyms:      m.Data.Synonyms,
			Antonyms:      m.Data.Antonyms,
		},
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

func (m Vocabulary) FromDomain(vocabulary domain.Vocabulary) (*Vocabulary, error) {
	id, err := database.ObjectIDFromString(vocabulary.ID)
	if err != nil {
		return nil, apperrors.Common.InvalidID
	}

	partsOfSpeech := make([]string, 0)
	for _, pos := range vocabulary.Data.PartsOfSpeech {
		partsOfSpeech = append(partsOfSpeech, pos.String())
	}

	return &Vocabulary{
		ID:       id,
		AuthorID: vocabulary.AuthorID,
		Term:     vocabulary.Term,
		Data: VocabularyData{
			PartsOfSpeech: partsOfSpeech,
			IPA:           vocabulary.Data.IPA,
			AudioName:     vocabulary.Data.AudioName,
			Synonyms:      vocabulary.Data.Synonyms,
			Antonyms:      vocabulary.Data.Antonyms,
		},
		CreatedAt: vocabulary.CreatedAt,
		UpdatedAt: vocabulary.UpdatedAt,
	}, nil
}
