package mapping

import (
	"github.com/goccy/go-json"
	"github.com/namhq1989/vocab-booster-english-hub/internal/database/gen/vocab-booster/public/model"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/domain"
	"github.com/namhq1989/vocab-booster-utilities/language"
)

type CommunitySentenceDraftMapper struct{}

func (CommunitySentenceDraftMapper) FromModelToDomain(sentence model.CommunitySentenceDrafts) (*domain.CommunitySentenceDraft, error) {
	var result = &domain.CommunitySentenceDraft{
		ID:                   sentence.ID,
		UserID:               sentence.UserID,
		VocabularyID:         sentence.VocabularyID,
		Content:              language.Multilingual{},
		RequiredVocabularies: sentence.RequiredVocabularies,
		RequiredTense:        domain.ToTense(sentence.RequiredTense),
		IsCorrect:            sentence.IsCorrect,
		ErrorCode:            domain.ToSentenceErrorCode(sentence.ErrorCode),
		GrammarErrors:        make([]domain.SentenceGrammarError, 0),
		Sentiment:            domain.Sentiment{},
		Clauses:              make([]domain.SentenceClause, 0),
		Level:                domain.ToSentenceLevel(sentence.Level),
		CreatedAt:            sentence.CreatedAt,
		UpdatedAt:            sentence.UpdatedAt,
	}

	if err := json.Unmarshal([]byte(sentence.GrammarErrors), &result.GrammarErrors); err != nil {
		return nil, err
	}

	if err := json.Unmarshal([]byte(sentence.Content), &result.Content); err != nil {
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
		ID:                   sentence.ID,
		UserID:               sentence.UserID,
		VocabularyID:         sentence.VocabularyID,
		Content:              "",
		RequiredVocabularies: sentence.RequiredVocabularies,
		RequiredTense:        sentence.RequiredTense.String(),
		IsCorrect:            sentence.IsCorrect,
		ErrorCode:            sentence.ErrorCode.String(),
		GrammarErrors:        "",
		Sentiment:            "",
		Clauses:              "",
		Level:                sentence.Level.String(),
		CreatedAt:            sentence.CreatedAt,
		UpdatedAt:            sentence.UpdatedAt,
	}

	grammarErrors := make([]SentenceGrammarError, 0)
	for _, grammarError := range sentence.GrammarErrors {
		grammarErrors = append(grammarErrors, SentenceGrammarError{
			Message:     grammarError.Message,
			Segment:     grammarError.Segment,
			Replacement: grammarError.Replacement,
		})
	}
	if data, err := json.Marshal(grammarErrors); err != nil {
		return nil, err
	} else {
		result.GrammarErrors = string(data)
	}

	if data, err := json.Marshal(sentence.Content); err != nil {
		return nil, err
	} else {
		result.Content = string(data)
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
