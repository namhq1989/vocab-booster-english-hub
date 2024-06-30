package mapping

import (
	"github.com/namhq1989/vocab-booster-english-hub/internal/database/gen/vocab-booster/public/model"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/domain"
)

type CommunitySentenceLikeMapper struct{}

func (CommunitySentenceLikeMapper) FromModelToDomain(sentence model.CommunitySentenceLikes) (*domain.CommunitySentenceLike, error) {
	var result = &domain.CommunitySentenceLike{
		ID:         sentence.ID,
		UserID:     sentence.UserID,
		SentenceID: sentence.SentenceID,
		CreatedAt:  sentence.CreatedAt,
	}

	return result, nil
}

func (CommunitySentenceLikeMapper) FromDomainToModel(sentence domain.CommunitySentenceLike) (*model.CommunitySentenceLikes, error) {
	var result = &model.CommunitySentenceLikes{
		ID:         sentence.ID,
		UserID:     sentence.UserID,
		SentenceID: sentence.SentenceID,
		CreatedAt:  sentence.CreatedAt,
	}

	return result, nil
}
