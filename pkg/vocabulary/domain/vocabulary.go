package domain

import (
	"time"

	"github.com/namhq1989/vocab-booster-english-hub/internal/database"
	apperrors "github.com/namhq1989/vocab-booster-english-hub/internal/utils/error"
	"github.com/namhq1989/vocab-booster-english-hub/internal/utils/manipulation"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
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
	Ipa           string
	Audio         string
	Synonyms      []string
	Antonyms      []string
	Frequency     float64
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
	d.SetUpdatedAt()

	return nil
}

func (d *Vocabulary) SetIPA(ipa string) error {
	if ipa == "" {
		return apperrors.Vocabulary.InvalidIPA
	}
	d.Ipa = ipa
	d.SetUpdatedAt()
	return nil
}

func (d *Vocabulary) SetAudio(audio string) error {
	if audio == "" {
		return apperrors.Vocabulary.InvalidAudioName
	}
	d.Audio = audio
	d.SetUpdatedAt()
	return nil
}

func (d *Vocabulary) SetLexicalRelations(synonyms, antonyms []string) error {
	d.Synonyms = synonyms
	d.Antonyms = antonyms
	d.SetUpdatedAt()
	return nil
}

func (d *Vocabulary) SetFrequency(frequency float64) error {
	d.Frequency = frequency
	d.SetUpdatedAt()
	return nil
}

func (d *Vocabulary) SetUpdatedAt() {
	d.UpdatedAt = manipulation.NowUTC()
}
