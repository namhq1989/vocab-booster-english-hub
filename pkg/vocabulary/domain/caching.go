package domain

import "github.com/namhq1989/vocab-booster-utilities/appcontext"

type CachingRepository interface {
	GetVocabularyByTerm(ctx *appcontext.AppContext, term string) (*Vocabulary, error)
	SetVocabularyByTerm(ctx *appcontext.AppContext, term string, vocabulary *Vocabulary) error

	GetVocabularyExamplesByVocabularyID(ctx *appcontext.AppContext, vocabularyID string) ([]VocabularyExample, error)
	SetVocabularyExamplesByVocabularyID(ctx *appcontext.AppContext, vocabularyID string, examples []VocabularyExample) error

	GetWordOfTheDay(ctx *appcontext.AppContext, country string) (*ExtendedWordOfTheDay, error)
	SetWordOfTheDay(ctx *appcontext.AppContext, wotd ExtendedWordOfTheDay, country string) error
	DeleteWordOfTheDay(ctx *appcontext.AppContext, country string) error
}
