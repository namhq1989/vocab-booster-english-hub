package domain

import (
	"time"

	"github.com/namhq1989/vocab-booster-english-hub/core/appcontext"
	apperrors "github.com/namhq1989/vocab-booster-english-hub/core/error"
	"github.com/namhq1989/vocab-booster-english-hub/internal/database"
)

type VocabularyRepository interface {
	FindVocabularyByID(ctx *appcontext.AppContext, vocabularyID string) (*Vocabulary, error)
	FindVocabularyByTerm(ctx *appcontext.AppContext, term string) (*Vocabulary, error)
	CreateVocabulary(ctx *appcontext.AppContext, vocabulary Vocabulary) error
	UpdateVocabulary(ctx *appcontext.AppContext, vocabulary Vocabulary) error
}

type Vocabulary struct {
	ID        string
	AuthorID  string
	Term      string
	Data      VocabularyData
	CreatedAt time.Time
	UpdatedAt time.Time
}

type VocabularyData struct {
	PartsOfSpeech []PartOfSpeech
	IPA           string
	AudioName     string
	Synonyms      []string
	Antonyms      []string
}

type VocabularyDefinition struct {
	POS        PartOfSpeech
	English    string
	Vietnamese string
}

func NewVocabulary(authorID, term string) (*Vocabulary, error) {
	if authorID == "" {
		return nil, apperrors.Vocabulary.InvalidAuthor
	}

	if term == "" {
		return nil, apperrors.Vocabulary.InvalidTerm
	}

	return &Vocabulary{
		ID:        database.NewStringID(),
		AuthorID:  authorID,
		Term:      term,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

func (d *Vocabulary) SetPartsOfSpeech(partsOfSpeech []string) error {
	dPartsOfSpeech := make([]PartOfSpeech, 0)
	for _, pos := range partsOfSpeech {
		partOfSpeech := ToPartOfSpeech(pos)
		if partOfSpeech.IsValid() {
			dPartsOfSpeech = append(dPartsOfSpeech, partOfSpeech)
		}
	}
	if len(dPartsOfSpeech) == 0 {
		return apperrors.Vocabulary.InvalidPartOfSpeech
	}
	d.Data.PartsOfSpeech = dPartsOfSpeech

	return nil
}

func (d *Vocabulary) SetIPA(ipa string) error {
	if ipa == "" {
		return apperrors.Vocabulary.InvalidIPA
	}
	d.Data.IPA = ipa
	return nil
}

func (d *Vocabulary) SetAudioName(audioName string) error {
	if audioName == "" {
		return apperrors.Vocabulary.InvalidAudioName
	}
	d.Data.AudioName = audioName
	return nil
}

func (d *Vocabulary) SetLexicalRelations(synonyms, antonyms []string) error {
	d.Data.Synonyms = synonyms
	d.Data.Antonyms = antonyms
	return nil
}

func (d *Vocabulary) SetUpdatedAt() {
	d.UpdatedAt = time.Now()
}
