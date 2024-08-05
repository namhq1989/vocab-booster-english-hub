package infrastructure

import (
	"fmt"
	"time"

	"github.com/goccy/go-json"
	"github.com/namhq1989/vocab-booster-english-hub/internal/caching"
	"github.com/namhq1989/vocab-booster-english-hub/internal/utils/manipulation"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/domain"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type CachingRepository struct {
	caching caching.Operations
}

func NewCachingRepository(caching *caching.Caching) CachingRepository {
	return CachingRepository{
		caching: caching,
	}
}

func (r CachingRepository) GetVocabularyByTerm(ctx *appcontext.AppContext, term string) (*domain.Vocabulary, error) {
	key := r.generateVocabularyKey(term)

	dataStr, err := r.caching.Get(ctx, key)
	if err != nil {
		return nil, err
	}

	var result *domain.Vocabulary
	if err = json.Unmarshal([]byte(dataStr), &result); err != nil {
		return nil, nil
	}

	return result, nil
}

func (r CachingRepository) SetVocabularyByTerm(ctx *appcontext.AppContext, term string, vocabulary *domain.Vocabulary) error {
	key := r.generateVocabularyKey(term)
	r.caching.SetTTL(ctx, key, vocabulary, 1*time.Hour)
	return nil
}

func (r CachingRepository) generateVocabularyKey(term string) string {
	return r.caching.GenerateKey("vocabulary", manipulation.Slugify(term))
}

func (r CachingRepository) GetVocabularyExamplesByVocabularyID(ctx *appcontext.AppContext, vocabularyID string) ([]domain.VocabularyExample, error) {
	key := r.generateVocabularyExamplesKey(vocabularyID)

	dataStr, err := r.caching.Get(ctx, key)
	if err != nil {
		return nil, err
	}

	var result []domain.VocabularyExample
	if err = json.Unmarshal([]byte(dataStr), &result); err != nil {
		return nil, nil
	}

	return result, nil
}

func (r CachingRepository) SetVocabularyExamplesByVocabularyID(ctx *appcontext.AppContext, vocabularyID string, examples []domain.VocabularyExample) error {
	key := r.generateVocabularyExamplesKey(vocabularyID)
	r.caching.SetTTL(ctx, key, examples, 1*time.Hour)
	return nil
}

func (r CachingRepository) generateVocabularyExamplesKey(vocabularyID string) string {
	return r.caching.GenerateKey("vocabularyExamples", vocabularyID)
}

func (r CachingRepository) GetWordOfTheDay(ctx *appcontext.AppContext, country string) (*domain.ExtendedWordOfTheDay, error) {
	key := r.generateWordOfTheDayKey(country)

	dataStr, err := r.caching.Get(ctx, key)
	if err != nil {
		return nil, err
	}

	var result *domain.ExtendedWordOfTheDay
	if err = json.Unmarshal([]byte(dataStr), &result); err != nil {
		return nil, nil
	}

	return result, nil
}

func (r CachingRepository) SetWordOfTheDay(ctx *appcontext.AppContext, wotd domain.ExtendedWordOfTheDay, country string) error {
	key := r.generateWordOfTheDayKey(country)
	r.caching.SetTTL(ctx, key, wotd, 8*time.Hour)
	return nil
}

func (r CachingRepository) DeleteWordOfTheDay(ctx *appcontext.AppContext, country string) error {
	key := r.generateWordOfTheDayKey(country)
	_, _ = r.caching.Del(ctx, key)
	return nil
}

func (r CachingRepository) generateWordOfTheDayKey(country string) string {
	return r.caching.GenerateKey("vocabulary", fmt.Sprintf("wotd:%s", country))
}
