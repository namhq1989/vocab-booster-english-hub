package domain

import (
	"time"

	"github.com/namhq1989/vocab-booster-english-hub/internal/database"
	apperrors "github.com/namhq1989/vocab-booster-english-hub/internal/utils/error"
	"github.com/namhq1989/vocab-booster-english-hub/internal/utils/pagetoken"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
	"github.com/namhq1989/vocab-booster-utilities/language"
)

type CommunitySentenceRepository interface {
	FindCommunitySentenceByID(ctx *appcontext.AppContext, id string) (*CommunitySentence, error)
	FindCommunitySentences(ctx *appcontext.AppContext, filter CommunitySentenceFilter) ([]ExtendedCommunitySentence, error)
	CreateCommunitySentence(ctx *appcontext.AppContext, sentence CommunitySentence) error
	UpdateCommunitySentence(ctx *appcontext.AppContext, sentence CommunitySentence) error
	FindCommunitySentenceWithUserID(ctx *appcontext.AppContext, sentenceID, userID string) (*ExtendedCommunitySentence, error)
}

type CommunitySentence struct {
	ID                   string
	UserID               string
	VocabularyID         string
	Content              language.Multilingual
	MainWord             VocabularyMainWord
	RequiredVocabularies []string
	RequiredTense        Tense
	Clauses              []SentenceClause
	PosTags              []PosTag
	Sentiment            Sentiment
	Dependencies         []Dependency
	Verbs                []Verb
	Level                SentenceLevel
	StatsLike            int
	CreatedAt            time.Time
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

func (s *CommunitySentence) SetContent(content language.Multilingual) error {
	if content.IsEmpty() {
		return apperrors.Common.InvalidContent
	}

	s.Content = content
	return nil
}

func (s *CommunitySentence) SetMainWordData(mainWord VocabularyMainWord) error {
	if mainWord.Base == "" {
		return apperrors.Vocabulary.InvalidTerm
	}

	if mainWord.Word == "" {
		return apperrors.Vocabulary.InvalidTerm
	}

	s.MainWord = mainWord
	return nil
}

func (s *CommunitySentence) SetRequiredVocabularies(required []string) error {
	if len(required) == 0 {
		return apperrors.Common.InvalidRequiredVocabularies
	}

	s.RequiredVocabularies = required
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

type CommunitySentenceFilter struct {
	UserID       string
	VocabularyID string
	Lang         string
	Timestamp    time.Time
	Limit        int64
}

func NewCommunitySentenceFilter(userID, vocabularyID, lang, pageToken string) (*CommunitySentenceFilter, error) {
	if !database.IsValidID(userID) {
		return nil, apperrors.User.InvalidUserID
	}

	if !database.IsValidID(vocabularyID) {
		return nil, apperrors.Vocabulary.InvalidVocabularyID
	}

	pt := pagetoken.Decode(pageToken)
	return &CommunitySentenceFilter{
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
	IsLiked bool
}
