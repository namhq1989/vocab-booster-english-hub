package infrastructure

import (
	"github.com/namhq1989/vocab-booster-english-hub/core/appcontext"
	"github.com/namhq1989/vocab-booster-english-hub/core/language"
	"github.com/namhq1989/vocab-booster-english-hub/internal/nlp"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/domain"
)

type NlpRepository struct {
	nlp nlp.Operations
}

func NewNlpRepository(nlp *nlp.NLP) NlpRepository {
	return NlpRepository{
		nlp: nlp,
	}
}

func (r NlpRepository) AnalyzeSentence(ctx *appcontext.AppContext, sentence string) (*domain.NlpSentenceAnalysisResult, error) {
	result, err := r.nlp.AnalyzeSentence(ctx, sentence)
	if err != nil {
		return nil, err
	}

	// convert result to domain data
	posTags := make([]domain.PosTag, 0)
	for _, posTag := range result.PosTags {
		posTags = append(posTags, domain.PosTag{
			Word:  posTag.Word,
			Value: domain.ToPartOfSpeech(posTag.Value),
			Level: posTag.Level,
		})
	}

	dependencies := make([]domain.Dependency, 0)
	for _, dependency := range result.Dependencies {
		dependencies = append(dependencies, domain.Dependency{
			Word:   dependency.Word,
			DepRel: dependency.DepRel,
			Head:   dependency.Head,
		})
	}

	verbs := make([]domain.Verb, 0)
	for _, verb := range result.Verbs {
		verbs = append(verbs, domain.Verb{
			Base:                verb.Base,
			Past:                verb.Past,
			PastParticiple:      verb.PastParticiple,
			Gerund:              verb.Gerund,
			ThirdPersonSingular: verb.ThirdPersonSingular,
		})
	}

	return &domain.NlpSentenceAnalysisResult{
		Translated: result.Translated,
		PosTags:    posTags,
		Sentiment: domain.Sentiment{
			Polarity:     result.Sentiment.Polarity,
			Subjectivity: result.Sentiment.Subjectivity,
		},
		Dependencies: dependencies,
		Verbs:        verbs,
		Level:        domain.ToSentenceLevel(result.Level),
	}, nil
}

func (r NlpRepository) TranslateDefinition(ctx *appcontext.AppContext, definition string) (*language.TranslatedLanguages, error) {
	result, err := r.nlp.TranslateDefinition(ctx, definition)
	if err != nil {
		return nil, err
	}

	return &language.TranslatedLanguages{
		Vi: result.Vi,
	}, nil
}
