package domain

import (
	"time"

	"github.com/namhq1989/vocab-booster-english-hub/core/appcontext"
	apperrors "github.com/namhq1989/vocab-booster-english-hub/core/error"
	"github.com/namhq1989/vocab-booster-english-hub/core/language"
	"github.com/namhq1989/vocab-booster-english-hub/internal/database"
)

type CommunitySentenceDraftRepository interface {
	FindCommunitySentenceDraft(ctx *appcontext.AppContext, vocabularyID, userID string) (*CommunitySentenceDraft, error)
	CreateCommunitySentenceDraft(ctx *appcontext.AppContext, sentence CommunitySentenceDraft) error
	DeleteCommunitySentenceDrafts(ctx *appcontext.AppContext, vocabularyID, userID string) error
}

type CommunitySentenceDraft struct {
	ID                  string
	UserID              string
	VocabularyID        string
	Content             string
	RequiredVocabulary  []string
	RequiredTense       Tense
	IsCorrect           bool
	IsGrammarCorrect    bool
	GrammarErrors       []SentenceGrammarError
	IsEnglish           bool
	IsTenseCorrect      bool
	IsVocabularyCorrect bool
	Translated          language.TranslatedLanguages
	Sentiment           Sentiment
	Clauses             []SentenceClause
	CreatedAt           time.Time
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
	}, nil
}

func (v *CommunitySentenceDraft) SetContent(content string) error {
	if content == "" {
		return apperrors.Common.InvalidContent
	}

	v.Content = content
	return nil
}

func (v *CommunitySentenceDraft) SetRequiredVocabulary(required []string) error {
	if len(required) == 0 {
		return apperrors.Common.InvalidRequiredVocabulary
	}

	v.RequiredVocabulary = required
	return nil
}

func (v *CommunitySentenceDraft) SetRequiredTense(tense string) error {
	dTense := ToTense(tense)
	if !dTense.IsValid() {
		return apperrors.Common.InvalidTense
	}

	v.RequiredTense = dTense
	return nil
}

func (v *CommunitySentenceDraft) setIsCorrect() {
	v.IsCorrect = v.IsEnglish && v.IsGrammarCorrect && v.IsVocabularyCorrect && v.IsTenseCorrect
}

func (v *CommunitySentenceDraft) setIsGrammarCorrect(value bool) {
	v.IsGrammarCorrect = value
	v.setIsCorrect()
}

func (v *CommunitySentenceDraft) SetIsEnglish(value bool) {
	v.IsEnglish = value
	v.setIsCorrect()
}

func (v *CommunitySentenceDraft) SetIsVocabularyCorrect(value bool) {
	v.IsVocabularyCorrect = value
	v.setIsCorrect()
}

func (v *CommunitySentenceDraft) SetIsTenseCorrect(value bool) {
	v.IsTenseCorrect = value
	v.setIsCorrect()
}

func (v *CommunitySentenceDraft) SetClauses(clauses []SentenceClause) error {
	v.Clauses = clauses
	return nil
}

func (v *CommunitySentenceDraft) SetSentiment(polarity, subjectivity float64) error {
	v.Sentiment.Polarity = polarity
	v.Sentiment.Subjectivity = subjectivity
	return nil
}

func (v *CommunitySentenceDraft) SetTranslated(translated language.TranslatedLanguages) error {
	v.Translated = translated
	return nil
}

func (v *CommunitySentenceDraft) SetGrammarErrors(errors []SentenceGrammarError) error {
	v.GrammarErrors = errors
	v.setIsGrammarCorrect(len(v.GrammarErrors) == 0)
	return nil
}
