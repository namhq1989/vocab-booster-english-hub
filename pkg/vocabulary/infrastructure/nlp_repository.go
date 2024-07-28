package infrastructure

import (
	"strings"

	"github.com/namhq1989/vocab-booster-english-hub/internal/nlp"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/domain"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
	"github.com/namhq1989/vocab-booster-utilities/language"
)

type NlpRepository struct {
	nlp nlp.Operations
}

func NewNlpRepository(nlp *nlp.NLP) NlpRepository {
	return NlpRepository{
		nlp: nlp,
	}
}

func (r NlpRepository) AnalyzeSentence(ctx *appcontext.AppContext, sentence, term string) (*domain.NlpSentenceAnalysisResult, error) {
	result, err := r.nlp.AnalyzeSentence(ctx, sentence, term)
	if err != nil {
		return nil, err
	}

	// convert result to domain data
	posTags := make([]domain.PosTag, 0)
	for _, posTag := range result.PosTags {
		posTags = append(posTags, domain.PosTag{
			Word:  strings.ToLower(posTag.Word),
			Value: domain.ToPartOfSpeech(posTag.Value),
			Level: posTag.Level,
		})
	}

	dependencies := make([]domain.Dependency, 0)
	for _, dependency := range result.Dependencies {
		dependencies = append(dependencies, domain.Dependency{
			Word:   strings.ToLower(dependency.Word),
			DepRel: strings.ToLower(dependency.DepRel),
			Head:   strings.ToLower(dependency.Head),
		})
	}

	verbs := make([]domain.Verb, 0)
	for _, verb := range result.Verbs {
		verbs = append(verbs, domain.Verb{
			Base:                strings.ToLower(verb.Base),
			Past:                strings.ToLower(verb.Past),
			PastParticiple:      strings.ToLower(verb.PastParticiple),
			Gerund:              strings.ToLower(verb.Gerund),
			ThirdPersonSingular: strings.ToLower(verb.ThirdPersonSingular),
		})
	}

	return &domain.NlpSentenceAnalysisResult{
		Translated: result.Translated,
		MainWord: domain.VocabularyMainWord{
			Word: strings.ToLower(result.MainWord.Word),
			Base: strings.ToLower(result.MainWord.Base),
			Pos:  domain.ToPartOfSpeech(result.MainWord.Pos),
		},
		PosTags: posTags,
		Sentiment: domain.Sentiment{
			Polarity:     result.Sentiment.Polarity,
			Subjectivity: result.Sentiment.Subjectivity,
		},
		Dependencies: dependencies,
		Verbs:        verbs,
		Level:        domain.ToSentenceLevel(result.Level),
	}, nil
}

func (r NlpRepository) TranslateDefinition(ctx *appcontext.AppContext, definition string) (*language.Multilingual, error) {
	return r.nlp.Translate(ctx, definition)
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

func (r NlpRepository) GrammarCheck(ctx *appcontext.AppContext, sentence string) ([]domain.SentenceGrammarError, error) {
	result, err := r.nlp.GrammarCheck(ctx, sentence)
	if err != nil {
		return nil, err
	}

	grammarErrors := make([]domain.SentenceGrammarError, 0)
	for _, ge := range result.Errors {
		sge, _ := domain.NewSentenceGrammarError(ge.Message, ge.Segment, ge.Replacement)
		if sge != nil {
			grammarErrors = append(grammarErrors, *sge)
		}
	}

	return grammarErrors, nil
}
