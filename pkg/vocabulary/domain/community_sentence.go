package domain

import (
	"time"

	apperrors "github.com/namhq1989/vocab-booster-english-hub/internal/utils/error"

	"github.com/namhq1989/vocab-booster-english-hub/internal/database"
	"github.com/namhq1989/vocab-booster-english-hub/internal/utils/pagetoken"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
	"github.com/namhq1989/vocab-booster-utilities/language"
)

type CommunitySentenceRepository interface {
	FindCommunitySentenceByID(ctx *appcontext.AppContext, id string) (*CommunitySentence, error)
	FindVocabularyCommunitySentences(ctx *appcontext.AppContext, filter VocabularyCommunitySentenceFilter) ([]ExtendedCommunitySentence, error)
	CreateCommunitySentence(ctx *appcontext.AppContext, sentence CommunitySentence) error
	UpdateCommunitySentence(ctx *appcontext.AppContext, sentence CommunitySentence) error
}

type CommunitySentence struct {
	ID                 string
	UserID             string
	VocabularyID       string
	Content            string
	RequiredVocabulary []string
	RequiredTense      Tense
	Translated         language.TranslatedLanguages
	Clauses            []SentenceClause
	PosTags            []PosTag
	Sentiment          Sentiment
	Dependencies       []Dependency
	Verbs              []Verb
	Level              SentenceLevel
	StatsLike          int
	CreatedAt          time.Time
}

func NewCommunitySentence(userID, vocabularyID string) (*CommunitySentence, error) {
	if !database.IsValidID(userID) {
		return nil, apperrors.User.InvalidUserID
	}

	if !database.IsValidID(vocabularyID) {
		return nil, apperrors.Vocabulary.InvalidVocabularyID
	}

	return &CommunitySentence{
		ID:           database.NewStringID(),
		UserID:       userID,
		VocabularyID: vocabularyID,
		CreatedAt:    time.Now(),
	}, nil
}

func (s *CommunitySentence) SetContent(content string) error {
	if content == "" {
		return apperrors.Common.InvalidContent
	}

	s.Content = content
	return nil
}

func (s *CommunitySentence) SetRequiredVocabulary(required []string) error {
	if len(required) == 0 {
		return apperrors.Common.InvalidRequiredVocabulary
	}

	s.RequiredVocabulary = required
	return nil
}

func (s *CommunitySentence) SetRequiredTense(tense string) error {
	dTense := ToTense(tense)
	if !dTense.IsValid() {
		return apperrors.Common.InvalidTense
	}

	s.RequiredTense = dTense
	return nil
}

func (s *CommunitySentence) SetClauses(clauses []SentenceClause) error {
	s.Clauses = clauses
	return nil
}

func (s *CommunitySentence) SetTranslated(translated language.TranslatedLanguages) error {
	s.Translated = translated
	return nil
}

func (s *CommunitySentence) SetPosTags(posTags []PosTag) error {
	if len(posTags) == 0 {
		return apperrors.Vocabulary.InvalidExamplePosTags
	}

	s.PosTags = posTags
	return nil
}

func (s *CommunitySentence) SetSentiment(polarity, subjectivity float64) error {
	s.Sentiment.Polarity = polarity
	s.Sentiment.Subjectivity = subjectivity
	return nil
}

func (s *CommunitySentence) SetDependencies(deps []Dependency) error {
	s.Dependencies = deps
	return nil
}

func (s *CommunitySentence) SetVerbs(verbs []Verb) error {
	s.Verbs = verbs
	return nil
}

func (s *CommunitySentence) SetLevel(level string) error {
	dLevel := ToSentenceLevel(level)
	if !dLevel.IsValid() {
		return apperrors.Vocabulary.InvalidExampleLevel
	}

	s.Level = dLevel
	return nil
}

func (s *CommunitySentence) IncreaseStatsLike() {
	s.StatsLike++
}

func (s *CommunitySentence) DecreaseStatsLike() {
	s.StatsLike--
	if s.StatsLike < 0 {
		s.StatsLike = 0
	}
}

//
// FILTER
//

type VocabularyCommunitySentenceFilter struct {
	UserID       string
	VocabularyID string
	Lang         string
	Timestamp    time.Time
	Limit        int64
}

func NewVocabularyCommunitySentenceFilter(userID, vocabularyID, lang, pageToken string) (*VocabularyCommunitySentenceFilter, error) {
	if !database.IsValidID(userID) {
		return nil, apperrors.User.InvalidUserID
	}

	if !database.IsValidID(vocabularyID) {
		return nil, apperrors.Vocabulary.InvalidVocabularyID
	}

	pt := pagetoken.Decode(pageToken)
	return &VocabularyCommunitySentenceFilter{
		UserID:       userID,
		VocabularyID: vocabularyID,
		Timestamp:    pt.Timestamp,
		Lang:         lang,
		Limit:        10,
	}, nil
}

//
// EXTENDED SENTENCE
//

type ExtendedCommunitySentence struct {
	CommunitySentence
	Translated string
	IsLiked    bool
}
