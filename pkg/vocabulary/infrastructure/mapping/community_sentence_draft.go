package mapping

import (
	"github.com/goccy/go-json"
	"github.com/namhq1989/vocab-booster-english-hub/core/language"
	"github.com/namhq1989/vocab-booster-english-hub/internal/database/gen/vocab-booster/public/model"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/domain"
)

type CommunitySentenceDraftMapper struct{}

func (CommunitySentenceDraftMapper) FromModelToDomain(sentence model.CommunitySentenceDrafts) (*domain.CommunitySentenceDraft, error) {
	var result = &domain.CommunitySentenceDraft{
		ID:                 sentence.ID,
		UserID:             sentence.UserID,
		VocabularyID:       sentence.VocabularyID,
		Content:            sentence.Content,
		RequiredVocabulary: sentence.RequiredVocabulary,
		RequiredTense:      domain.ToTense(sentence.RequiredTense),
		IsCorrect:          sentence.IsCorrect,
		ErrorCode:          domain.ToSentenceErrorCode(sentence.ErrorCode),
		GrammarErrors:      make([]domain.SentenceGrammarError, 0),
		Translated:         language.TranslatedLanguages{},
		Sentiment:          domain.Sentiment{},
		Clauses:            make([]domain.SentenceClause, 0),
		CreatedAt:          sentence.CreatedAt,
		UpdatedAt:          sentence.UpdatedAt,
	}

	if err := json.Unmarshal([]byte(sentence.GrammarErrors), &result.GrammarErrors); err != nil {
		return nil, err
	}

	if err := json.Unmarshal([]byte(sentence.Translated), &result.Translated); err != nil {
		return nil, err
	}

	if err := json.Unmarshal([]byte(sentence.Sentiment), &result.Sentiment); err != nil {
		return nil, err
	}

	if err := json.Unmarshal([]byte(sentence.Clauses), &result.Clauses); err != nil {
		return nil, err
	}

	return result, nil
}

func (CommunitySentenceDraftMapper) FromDomainToModel(sentence domain.CommunitySentenceDraft) (*model.CommunitySentenceDrafts, error) {
	var result = &model.CommunitySentenceDrafts{
		ID:                 sentence.ID,
		UserID:             sentence.UserID,
		VocabularyID:       sentence.VocabularyID,
		Content:            sentence.Content,
		RequiredVocabulary: sentence.RequiredVocabulary,
		RequiredTense:      sentence.RequiredTense.String(),
		IsCorrect:          sentence.IsCorrect,
		ErrorCode:          sentence.ErrorCode.String(),
		GrammarErrors:      "",
		Translated:         "",
		Sentiment:          "",
		Clauses:            "",
		CreatedAt:          sentence.CreatedAt,
		UpdatedAt:          sentence.UpdatedAt,
	}

	grammarErrors := make([]SentenceGrammarError, 0)
	for _, grammarError := range sentence.GrammarErrors {
		grammarErrors = append(grammarErrors, SentenceGrammarError{
			Message:     grammarError.Message,
			Translated:  grammarError.Translated,
			Segment:     grammarError.Segment,
			Replacement: grammarError.Replacement,
		})
	}
	if data, err := json.Marshal(grammarErrors); err != nil {
		return nil, err
	} else {
		result.GrammarErrors = string(data)
	}

	if data, err := json.Marshal(sentence.Translated); err != nil {
		return nil, err
	} else {
		result.Translated = string(data)
	}

	sentiment := Sentiment{
		Polarity:     sentence.Sentiment.Polarity,
		Subjectivity: sentence.Sentiment.Subjectivity,
	}
	if data, err := json.Marshal(sentiment); err != nil {
		return nil, err
	} else {
		result.Sentiment = string(data)
	}

	clauses := make([]SentenceClause, 0)
	for _, clause := range sentence.Clauses {
		clauses = append(clauses, SentenceClause{
			Clause:         clause.Clause,
			Tense:          clause.Tense.String(),
			IsPrimaryTense: clause.IsPrimaryTense,
		})
	}
	if data, err := json.Marshal(clauses); err != nil {
		return nil, err
	} else {
		result.Clauses = string(data)
	}

	return result, nil
}
