package dto

import (
	"github.com/namhq1989/vocab-booster-english-hub/internal/genproto/vocabularypb"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/domain"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ConvertCommunitySentenceBriefFromDomainToGrpc(sentence domain.ExtendedCommunitySentence) *vocabularypb.CommunitySentenceBrief {
	result := &vocabularypb.CommunitySentenceBrief{
		Id:        sentence.ID,
		Content:   ConvertMultilingualToGrpcData(sentence.Content),
		Level:     sentence.Level.String(),
		StatsLike: int32(sentence.StatsLike),
		IsLiked:   sentence.IsLiked,
	}

	return result
}

func ConvertCommunitySentenceFromDomainToGrpc(sentence domain.ExtendedCommunitySentence) *vocabularypb.CommunitySentence {
	clauses := make([]*vocabularypb.SentenceClause, len(sentence.Clauses))
	for i, clause := range sentence.Clauses {
		clauses[i] = &vocabularypb.SentenceClause{
			Clause:         clause.Clause,
			Tense:          clause.Tense.String(),
			IsPrimaryTense: clause.IsPrimaryTense,
		}
	}

	posTags := make([]*vocabularypb.PosTag, len(sentence.PosTags))
	for i, posTag := range sentence.PosTags {
		posTags[i] = &vocabularypb.PosTag{
			Word:  posTag.Word,
			Value: posTag.Value.String(),
			Level: int32(posTag.Level),
		}
	}

	dependencies := make([]*vocabularypb.Dependency, len(sentence.Dependencies))
	for i, dependency := range sentence.Dependencies {
		dependencies[i] = &vocabularypb.Dependency{
			Word:   dependency.Word,
			DepRel: dependency.DepRel,
			Head:   dependency.Head,
		}
	}

	verbs := make([]*vocabularypb.Verb, len(sentence.Verbs))
	for i, verb := range sentence.Verbs {
		verbs[i] = &vocabularypb.Verb{
			Base:                verb.Base,
			Past:                verb.Past,
			PastParticiple:      verb.PastParticiple,
			Gerund:              verb.Gerund,
			ThirdPersonSingular: verb.ThirdPersonSingular,
		}
	}

	result := &vocabularypb.CommunitySentence{
		Id:                 sentence.ID,
		VocabularyId:       sentence.VocabularyID,
		Content:            ConvertMultilingualToGrpcData(sentence.Content),
		RequiredVocabulary: sentence.RequiredVocabulary,
		RequiredTense:      sentence.RequiredTense.String(),
		Clauses:            clauses,
		PosTags:            posTags,
		Sentiment: &vocabularypb.Sentiment{
			Polarity:     sentence.Sentiment.Polarity,
			Subjectivity: sentence.Sentiment.Subjectivity,
		},
		Dependencies: dependencies,
		Verbs:        verbs,
		Level:        sentence.Level.String(),
		StatsLike:    int32(sentence.StatsLike),
		IsLiked:      sentence.IsLiked,
		CreatedAt:    timestamppb.New(sentence.CreatedAt),
	}

	return result
}
