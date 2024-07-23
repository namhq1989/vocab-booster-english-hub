package dto

import (
	"github.com/namhq1989/vocab-booster-english-hub/internal/genproto/vocabularypb"
	"github.com/namhq1989/vocab-booster-utilities/language"
)

func ConvertMultilingualToGrpcData(multilingual language.Multilingual) *vocabularypb.Multilingual {
	return &vocabularypb.Multilingual{
		English:    multilingual.English,
		Vietnamese: multilingual.Vietnamese,
	}
}

func ConvertGrpcDataToMultilingual(data *vocabularypb.Multilingual) language.Multilingual {
	return language.Multilingual{
		English:    data.GetEnglish(),
		Vietnamese: data.GetVietnamese(),
	}
}
