package domain

import (
	"time"

	"github.com/namhq1989/vocab-booster-english-hub/internal/utils/pagetoken"

	apperrors "github.com/namhq1989/vocab-booster-english-hub/internal/utils/error"

	"github.com/namhq1989/vocab-booster-english-hub/internal/database"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
	"github.com/namhq1989/vocab-booster-utilities/language"
)

type CommunitySentenceDraftRepository interface {
	FindCommunitySentenceDraftByID(ctx *appcontext.AppContext, sentenceID string) (*CommunitySentenceDraft, error)
	CreateCommunitySentenceDraft(ctx *appcontext.AppContext, sentence CommunitySentenceDraft) error
	UpdateCommunitySentenceDraft(ctx *appcontext.AppContext, sentence CommunitySentenceDraft) error
	DeleteCommunitySentenceDraft(ctx *appcontext.AppContext, vocabularyID, userID string) error
	FindUserCommunitySentenceDrafts(ctx *appcontext.AppContext, filter CommunitySentenceDraftFilter) ([]CommunitySentenceDraft, error)
}

type CommunitySentenceDraft struct {
	ID                   string
	UserID               string
	VocabularyID         string
	Content              language.Multilingual
	RequiredVocabularies []string
	RequiredTense        Tense
	IsCorrect            bool
	ErrorCode            SentenceErrorCode
	GrammarErrors        []SentenceGrammarError
	Sentiment            Sentiment
	Clauses              []SentenceClause
	Level                SentenceLevel
	CreatedAt            time.Time
	UpdatedAt            time.Time
}

func NewCommunitySentenceDraft(userID, vocabularyID string) (*CommunitySentenceDraft, error) {
	if !database.IsValidID(userID) {
		return nil, apperrors.User.InvalidUserID
	}

	if !database.IsValidID(vocabularyID) {
		return nil, apperrors.Vocabulary.InvalidVocabularyID
	}

	return &CommunitySentenceDraft{
		ID:           database.NewStringID(),
		UserID:       userID,
		VocabularyID: vocabularyID,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}, nil
}

func (s *CommunitySentenceDraft) SetContent(content language.Multilingual) error {
	if content.IsEmpty() {
		return apperrors.Common.InvalidContent
	}

	s.Content = content
	return nil
}

func (s *CommunitySentenceDraft) SetRequiredVocabularies(required []string) error {
	if len(required) == 0 {
		return apperrors.Common.InvalidRequiredVocabularies
	}

	s.RequiredVocabularies = required
	return nil
}

func (s *CommunitySentenceDraft) SetRequiredTense(tense string) error {
	dTense := ToTense(tense)
	if !dTense.IsValid() {
		return apperrors.Common.InvalidTense
	}

	s.RequiredTense = dTense
	return nil
}

func (s *CommunitySentenceDraft) setIsCorrect() {
	s.IsCorrect = s.ErrorCode.IsEmpty()
}

func (s *CommunitySentenceDraft) SetErrorCode(code SentenceErrorCode) {
	s.ErrorCode = code
	s.setIsCorrect()
}

func (s *CommunitySentenceDraft) SetClauses(clauses []SentenceClause) error {
	s.Clauses = clauses
	return nil
}

func (s *CommunitySentenceDraft) SetSentiment(polarity, subjectivity float64) error {
	s.Sentiment.Polarity = polarity
	s.Sentiment.Subjectivity = subjectivity
	return nil
}

func (s *CommunitySentenceDraft) SetGrammarErrors(errors []SentenceGrammarError) error {
	s.GrammarErrors = errors
	if len(s.GrammarErrors) > 0 {
		s.SetErrorCode(SentenceErrorCodeInvalidGrammar)
	} else {
		s.SetErrorCode(SentenceErrorCodeEmpty)
	}
	return nil
}

func (s *CommunitySentenceDraft) SetLevel(level string) error {
	dLevel := ToSentenceLevel(level)
	if !dLevel.IsValid() {
		return apperrors.Vocabulary.InvalidExampleLevel
	}

	s.Level = dLevel
	return nil
}

func (s *CommunitySentenceDraft) SetUpdatedAt() {
	s.UpdatedAt = time.Now()
}

func (s *CommunitySentenceDraft) IsOwner(userID string) bool {
	return s.UserID == userID
}

//
// FILTER
//

type CommunitySentenceDraftFilter struct {
	UserID       string
	VocabularyID string
	Timestamp    time.Time
	Limit        int64
}

func NewCommunitySentenceDraftFilter(userID, vocabularyID, pageToken string) (*CommunitySentenceDraftFilter, error) {
	if !database.IsValidID(userID) {
		return nil, apperrors.User.InvalidUserID
	}

	if !database.IsValidID(vocabularyID) {
		return nil, apperrors.Vocabulary.InvalidVocabularyID
	}

	pt := pagetoken.Decode(pageToken)
	return &CommunitySentenceDraftFilter{
		UserID:       userID,
		VocabularyID: vocabularyID,
		Timestamp:    pt.Timestamp,
		Limit:        10,
	}, nil
}
