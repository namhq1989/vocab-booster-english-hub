package dto

import (
	"github.com/namhq1989/vocab-booster-english-hub/internal/genproto/exercisepb"
	"github.com/namhq1989/vocab-booster-utilities/language"
)

func ConvertMultilingualToGrpcData(multilingual language.Multilingual) *exercisepb.Multilingual {
	return &exercisepb.Multilingual{
		English:    multilingual.English,
		Vietnamese: multilingual.Vietnamese,
	}
}

func ConvertGrpcDataToMultilingual(data *exercisepb.Multilingual) language.Multilingual {
	return language.Multilingual{
		English:    data.GetEnglish(),
		Vietnamese: data.GetVietnamese(),
	}
}
