package mapping

import (
	"github.com/goccy/go-json"
	"github.com/namhq1989/vocab-booster-english-hub/internal/database"
	"github.com/namhq1989/vocab-booster-english-hub/internal/database/gen/vocab-booster/public/model"
	apperrors "github.com/namhq1989/vocab-booster-english-hub/internal/utils/error"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/domain"
	"github.com/namhq1989/vocab-booster-utilities/language"
)

type WordOfTheDayMapper struct{}

func (WordOfTheDayMapper) FromModelToDomain(word model.WordOfTheDay) (*domain.WordOfTheDay, error) {
	var result = &domain.WordOfTheDay{
		ID:           word.ID,
		VocabularyID: word.VocabularyID,
		Country:      word.Country,
		Information:  language.Multilingual{},
		Date:         word.Date,
	}

	if err := json.Unmarshal([]byte(word.Information), &result.Information); err != nil {
		return nil, err
	}

	return result, nil
}

func (WordOfTheDayMapper) FromDomainToModel(word domain.WordOfTheDay) (*model.WordOfTheDay, error) {
	if !database.IsValidID(word.VocabularyID) {
		return nil, apperrors.Vocabulary.InvalidVocabularyID
	}

	var result = &model.WordOfTheDay{
		ID:           word.ID,
		VocabularyID: word.VocabularyID,
		Country:      word.Country,
		Information:  "",
		Date:         word.Date,
	}

	if data, err := json.Marshal(word.Information); err != nil {
		return nil, err
	} else {
		result.Information = string(data)
	}

	return result, nil
}

type ExtendedWordOfTheDay struct {
	WordOfTheDay model.WordOfTheDay `alias:"wotd"`
	Vocabulary   model.Vocabularies `alias:"v"`
}

type ExtendedWordOfTheDayMapper struct{}

func (ExtendedWordOfTheDayMapper) FromModelToDomain(ewotd ExtendedWordOfTheDay) (*domain.ExtendedWordOfTheDay, error) {
	var vocabularyMapper = VocabularyMapper{}
	vocabulary, err := vocabularyMapper.FromModelToDomain(ewotd.Vocabulary)
	if err != nil {
		return nil, err
	}

	var wordOfTheDayMapper = WordOfTheDayMapper{}
	wotd, err := wordOfTheDayMapper.FromModelToDomain(ewotd.WordOfTheDay)
	if err != nil {
		return nil, err
	}

	return &domain.ExtendedWordOfTheDay{
		Vocabulary:   *vocabulary,
		WordOfTheDay: *wotd,
	}, nil
}
