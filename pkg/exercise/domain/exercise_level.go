package domain

type ExerciseLevel string

const (
	ExerciseLevelUnknown      ExerciseLevel = ""
	ExerciseLevelBeginner     ExerciseLevel = "beginner"
	ExerciseLevelIntermediate ExerciseLevel = "intermediate"
	ExerciseLevelAdvanced     ExerciseLevel = "advanced"
)

func (s ExerciseLevel) IsValid() bool {
	return s != ExerciseLevelUnknown
}

func (s ExerciseLevel) String() string {
	return string(s)
}

func ToExerciseLevel(value string) ExerciseLevel {
	switch value {
	case ExerciseLevelBeginner.String():
		return ExerciseLevelBeginner
	case ExerciseLevelIntermediate.String():
		return ExerciseLevelIntermediate
	case ExerciseLevelAdvanced.String():
		return ExerciseLevelAdvanced
	default:
		return ExerciseLevelUnknown
	}
}
