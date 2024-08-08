package dto

import (
	"github.com/namhq1989/vocab-booster-english-hub/internal/genproto/vocabularypb"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/domain"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ConvertCommunitySentenceDraftFromDomainToGrpc(sentence domain.CommunitySentenceDraft) *vocabularypb.CommunitySentenceDraft {
	clauses := make([]*vocabularypb.SentenceClause, len(sentence.Clauses))
	for i, clause := range sentence.Clauses {
		clauses[i] = &vocabularypb.SentenceClause{
			Clause:         clause.Clause,
			Tense:          clause.Tense.String(),
			IsPrimaryTense: clause.IsPrimaryTense,
		}
	}

	grammarErrors := make([]*vocabularypb.SentenceGrammarError, len(sentence.GrammarErrors))
	for i, grammarError := range sentence.GrammarErrors {
		grammarErrors[i] = &vocabularypb.SentenceGrammarError{
			Message:     ConvertMultilingualToGrpcData(grammarError.Message),
			Segment:     grammarError.Segment,
			Replacement: grammarError.Replacement,
		}
	}

	result := &vocabularypb.CommunitySentenceDraft{
		Id:                   sentence.ID,
		Content:              ConvertMultilingualToGrpcData(sentence.Content),
		RequiredVocabularies: sentence.RequiredVocabularies,
		RequiredTense:        sentence.RequiredTense.String(),
		Clauses:              clauses,
		IsCorrect:            sentence.IsCorrect,
		ErrorCode:            sentence.ErrorCode.String(),
		Sentiment: &vocabularypb.Sentiment{
			Polarity:     sentence.Sentiment.Polarity,
			Subjectivity: sentence.Sentiment.Subjectivity,
		},
		Level:     sentence.Level.String(),
		Errors:    grammarErrors,
		CreatedAt: timestamppb.New(sentence.CreatedAt),
	}

	return result
}
