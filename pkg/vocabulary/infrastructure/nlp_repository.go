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

func (r NlpRepository) EvaluateSentence(ctx *appcontext.AppContext, sentence, tense string, vocabulary []string) (*domain.NlpSentenceEvaluationResult, error) {
	result, err := r.nlp.EvaluateSentence(ctx, sentence, tense, vocabulary)
	if err != nil {
		return nil, err
	}

	// convert result to domain data
	clauses := make([]domain.SentenceClause, 0)
	for _, clause := range result.Clauses {
		c, _ := domain.NewSentenceClause(clause.Clause, clause.Tense, clause.IsPrimaryTense)
		if c != nil {
			clauses = append(clauses, *c)
		}
	}

	return &domain.NlpSentenceEvaluationResult{
		IsEnglish:           result.IsEnglish,
		IsVocabularyCorrect: result.IsVocabularyCorrect,
		IsTenseCorrect:      result.IsTenseCorrect,
		Sentiment: domain.Sentiment{
			Polarity:     result.Sentiment.Polarity,
			Subjectivity: result.Sentiment.Subjectivity,
		},
		Translated: result.Translated,
		Clauses:    clauses,
	}, nil
}
