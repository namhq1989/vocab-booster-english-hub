package domain

import (
	"github.com/namhq1989/vocab-booster-english-hub/internal/database"
	"github.com/namhq1989/vocab-booster-utilities/language"
)

var SystemCollections = []ExerciseCollection{
	{
		ID: database.NewStringID(),
		Name: language.Multilingual{
			English:    "All",
			Vietnamese: "Tất cả",
		},
		Slug:           "all",
		Criteria:       "",
		IsFromSystem:   true,
		StatsExercises: 0,
		Order:          0,
		Image:          "all.svg",
	},
	{
		ID: database.NewStringID(),
		Name: language.Multilingual{
			English:    "Beginner",
			Vietnamese: "Mới toe",
		},
		Slug:           "beginner",
		Criteria:       "level=beginner",
		IsFromSystem:   true,
		StatsExercises: 0,
		Order:          1,
		Image:          "beginner.svg",
	},
	{
		ID: database.NewStringID(),
		Name: language.Multilingual{
			English:    "Intermediate",
			Vietnamese: "Tầm trung",
		},
		Slug:           "intermediate",
		Criteria:       "level=intermediate",
		IsFromSystem:   true,
		StatsExercises: 0,
		Order:          2,
		Image:          "intermediate.svg",
	},
	{
		ID: database.NewStringID(),
		Name: language.Multilingual{
			English:    "Advanced",
			Vietnamese: "Rành rọt",
		},
		Slug:           "advanced",
		Criteria:       "level=advanced",
		IsFromSystem:   true,
		StatsExercises: 0,
		Order:          3,
		Image:          "advanced.svg",
	},
	{
		ID: database.NewStringID(),
		Name: language.Multilingual{
			English:    "Essential picks",
			Vietnamese: "Từ vựng cốt yếu",
		},
		Slug:           "essential-picks",
		Criteria:       "frequency=1000",
		IsFromSystem:   true,
		StatsExercises: 0,
		Order:          10,
		Image:          "advanced.svg",
	},
	{
		ID: database.NewStringID(),
		Name: language.Multilingual{
			English:    "Common words",
			Vietnamese: "Từ vựng thông dụng",
		},
		Slug:           "common-words",
		Criteria:       "frequency=500",
		IsFromSystem:   true,
		StatsExercises: 0,
		Order:          11,
		Image:          "advanced.svg",
	},
	{
		ID: database.NewStringID(),
		Name: language.Multilingual{
			English:    "Must-knows",
			Vietnamese: "Từ vựng cần biết",
		},
		Slug:           "must-knows",
		Criteria:       "frequency=100",
		IsFromSystem:   true,
		StatsExercises: 0,
		Order:          12,
		Image:          "advanced.svg",
	},
	{
		ID: database.NewStringID(),
		Name: language.Multilingual{
			English:    "Comprehensive lexicon",
			Vietnamese: "Từ điển toàn diện",
		},
		Slug:           "comprehensive-lexicon",
		Criteria:       "frequency=10",
		IsFromSystem:   true,
		StatsExercises: 0,
		Order:          13,
		Image:          "advanced.svg",
	},
}
