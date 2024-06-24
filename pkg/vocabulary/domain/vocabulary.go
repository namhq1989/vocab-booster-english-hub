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
	RandomPickVocabularyForExercise(ctx *appcontext.AppContext, numOfVocabulary int64) ([]Vocabulary, error)
}

type Vocabulary struct {
	ID            string
	AuthorID      string
	Term          string
	PartsOfSpeech []PartOfSpeech
	IPA           string
	Audio         string
	Synonyms      []string
	Antonyms      []string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type VocabularyData struct {
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
	d.PartsOfSpeech = dPartsOfSpeech

	return nil
}

func (d *Vocabulary) SetIPA(ipa string) error {
	if ipa == "" {
		return apperrors.Vocabulary.InvalidIPA
	}
	d.IPA = ipa
	return nil
}

func (d *Vocabulary) SetAudio(audio string) error {
	if audio == "" {
		return apperrors.Vocabulary.InvalidAudioName
	}
	d.Audio = audio
	return nil
}

func (d *Vocabulary) SetLexicalRelations(synonyms, antonyms []string) error {
	d.Synonyms = synonyms
	d.Antonyms = antonyms
	return nil
}

func (d *Vocabulary) SetUpdatedAt() {
	d.UpdatedAt = time.Now()
}
