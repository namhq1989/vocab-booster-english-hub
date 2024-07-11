package dto

import (
	"github.com/namhq1989/vocab-booster-english-hub/internal/genproto/vocabularypb"
	"github.com/namhq1989/vocab-booster-utilities/language"
)

func ConvertTranslatedLanguagesToGrpcData(translatedLanguages language.TranslatedLanguages) *vocabularypb.TranslatedLanguages {
	return &vocabularypb.TranslatedLanguages{
		Vietnamese: translatedLanguages.Vietnamese,
	}
}

func ConvertGrpcDataToTranslatedLanguages(data *vocabularypb.TranslatedLanguages) language.TranslatedLanguages {
	return language.TranslatedLanguages{
		Vietnamese: data.GetVietnamese(),
	}
}
