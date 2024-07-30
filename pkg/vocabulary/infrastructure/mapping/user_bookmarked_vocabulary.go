package mapping

import (
	"time"

	"github.com/namhq1989/vocab-booster-english-hub/internal/database"
	"github.com/namhq1989/vocab-booster-english-hub/internal/database/gen/vocab-booster/public/model"
	apperrors "github.com/namhq1989/vocab-booster-english-hub/internal/utils/error"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/domain"
)

type UserBookmarkedVocabularyMapper struct{}

func (UserBookmarkedVocabularyMapper) FromModelToDomain(ubv model.UserBookmarkedVocabulary) (*domain.UserBookmarkedVocabulary, error) {
	return &domain.UserBookmarkedVocabulary{
		ID:           ubv.ID,
		UserID:       ubv.UserID,
		VocabularyID: ubv.VocabularyID,
		BookmarkedAt: ubv.BookmarkedAt,
	}, nil
}

func (UserBookmarkedVocabularyMapper) FromDomainToModel(ubv domain.UserBookmarkedVocabulary) (*model.UserBookmarkedVocabulary, error) {
	if !database.IsValidID(ubv.ID) {
		return nil, apperrors.Common.InvalidID
	}

	if !database.IsValidID(ubv.UserID) {
		return nil, apperrors.User.InvalidUserID
	}

	if !database.IsValidID(ubv.VocabularyID) {
		return nil, apperrors.Vocabulary.InvalidVocabularyID
	}

	return &model.UserBookmarkedVocabulary{
		ID:           ubv.ID,
		UserID:       ubv.UserID,
		VocabularyID: ubv.VocabularyID,
		BookmarkedAt: ubv.BookmarkedAt,
	}, nil
}

//
// EXTENDED
//

type ExtendedUserBookmarkedVocabulary struct {
	Vocabulary   model.Vocabularies `alias:"v"`
	BookmarkedAt time.Time          `alias:"ubv.bookmarked_at"`
}

type ExtendedUserBookmarkedVocabularyMapper struct{}

func (ExtendedUserBookmarkedVocabularyMapper) FromModelToDomain(eubv ExtendedUserBookmarkedVocabulary) (*domain.ExtendedUserBookmarkedVocabulary, error) {
	var vocabularyMapper = VocabularyMapper{}
	vocabulary, err := vocabularyMapper.FromModelToDomain(eubv.Vocabulary)
	if err != nil {
		return nil, err
	}

	return &domain.ExtendedUserBookmarkedVocabulary{
		Vocabulary:   *vocabulary,
		BookmarkedAt: eubv.BookmarkedAt,
	}, nil
}
