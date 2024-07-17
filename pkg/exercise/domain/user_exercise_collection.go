package domain

import "github.com/namhq1989/vocab-booster-utilities/language"

type UserExerciseCollection struct {
	ID              string
	Name            string
	Slug            string
	Translated      language.TranslatedLanguages
	Order           int
	Image           string
	StatsExercises  int
	StatsInteracted int
}
