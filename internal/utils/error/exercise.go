package apperrors

import "errors"

var Exercise = struct {
	ExerciseNotFound           error
	InvalidExerciseID          error
	InvalidVocabularyExampleID error
	InvalidAudio               error
	InvalidLevel               error
	InvalidContent             error
	InvalidVocabulary          error
	InvalidCorrectAnswer       error
	InvalidOptions             error
}{
	ExerciseNotFound:           errors.New("exercise_not_found"),
	InvalidExerciseID:          errors.New("exercise_invalid_id"),
	InvalidVocabularyExampleID: errors.New("exercise_invalid_vocabulary_example_id"),
	InvalidAudio:               errors.New("exercise_invalid_audio"),
	InvalidLevel:               errors.New("exercise_invalid_level"),
	InvalidContent:             errors.New("exercise_invalid_content"),
	InvalidVocabulary:          errors.New("exercise_invalid_vocabulary"),
	InvalidCorrectAnswer:       errors.New("exercise_invalid_correct_answer"),
	InvalidOptions:             errors.New("exercise_invalid_options"),
}
