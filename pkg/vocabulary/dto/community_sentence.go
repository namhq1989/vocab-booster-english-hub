package dto

import (
	"github.com/namhq1989/vocab-booster-english-hub/internal/genproto/vocabularypb"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/domain"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ConvertCommunitySentencesFromDomainToGrpc(sentences []domain.ExtendedCommunitySentence) []*vocabularypb.CommunitySentence {
	var result = make([]*vocabularypb.CommunitySentence, len(sentences))

	for index, sentence := range sentences {
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

		result[index] = &vocabularypb.CommunitySentence{
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
	}

	return result
}
