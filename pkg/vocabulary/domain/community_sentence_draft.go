package domain

import (
	"time"

	"github.com/namhq1989/vocab-booster-english-hub/core/appcontext"
	apperrors "github.com/namhq1989/vocab-booster-english-hub/core/error"
	"github.com/namhq1989/vocab-booster-english-hub/core/language"
	"github.com/namhq1989/vocab-booster-english-hub/internal/database"
)

type CommunitySentenceDraftRepository interface {
	FindCommunitySentenceDraftByID(ctx *appcontext.AppContext, sentenceID string) (*CommunitySentenceDraft, error)
	CreateCommunitySentenceDraft(ctx *appcontext.AppContext, sentence CommunitySentenceDraft) error
	UpdateCommunitySentenceDraft(ctx *appcontext.AppContext, sentence CommunitySentenceDraft) error
	DeleteCommunitySentenceDrafts(ctx *appcontext.AppContext, vocabularyID, userID string) error
}

type CommunitySentenceDraft struct {
	ID                 string
	UserID             string
	VocabularyID       string
	Content            string
	RequiredVocabulary []string
	RequiredTense      Tense
	IsCorrect          bool
	ErrorCode          SentenceErrorCode
	GrammarErrors      []SentenceGrammarError
	Translated         language.TranslatedLanguages
	Sentiment          Sentiment
	Clauses            []SentenceClause
	CreatedAt          time.Time
	UpdatedAt          time.Time
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

func (s *CommunitySentenceDraft) SetContent(content string) error {
	if content == "" {
		return apperrors.Common.InvalidContent
	}

	s.Content = content
	return nil
}

func (s *CommunitySentenceDraft) SetRequiredVocabulary(required []string) error {
	if len(required) == 0 {
		return apperrors.Common.InvalidRequiredVocabulary
	}

	s.RequiredVocabulary = required
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

func (s *CommunitySentenceDraft) SetTranslated(translated language.TranslatedLanguages) error {
	s.Translated = translated
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

func (s *CommunitySentenceDraft) SetUpdatedAt() {
	s.UpdatedAt = time.Now()
}

func (s *CommunitySentenceDraft) IsOwner(userID string) bool {
	return s.UserID == userID
}
