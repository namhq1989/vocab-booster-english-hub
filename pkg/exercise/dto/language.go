package dto

import (
	"github.com/namhq1989/vocab-booster-english-hub/internal/genproto/exercisepb"
	"github.com/namhq1989/vocab-booster-utilities/language"
)

// func ConvertTranslatedLanguagesToGrpcData(translatedLanguages language.TranslatedLanguages) *exercisepb.TranslatedLanguages {
// 	return &exercisepb.TranslatedLanguages{
// 		Vietnamese: translatedLanguages.Vietnamese,
// 	}
// }

func ConvertGrpcDataToTranslatedLanguages(data *exercisepb.TranslatedLanguages) language.TranslatedLanguages {
	return language.TranslatedLanguages{
		Vietnamese: data.GetVietnamese(),
	}
}
