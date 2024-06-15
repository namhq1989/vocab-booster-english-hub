package apperrors

import "errors"

var Vocabulary = struct {
	VocabularyNotFound     error
	InvalidVocabularyID    error
	InvalidTerm            error
	InvalidAuthor          error
	InvalidDefinition      error
	InvalidPartOfSpeech    error
	InvalidIPA             error
	InvalidAudioName       error
	InvalidExampleLanguage error
	InvalidExamplePosTags  error
}{
	VocabularyNotFound:     errors.New("vocabulary_not_found"),
	InvalidVocabularyID:    errors.New("vocabulary_invalid_id"),
	InvalidTerm:            errors.New("vocabulary_invalid_term"),
	InvalidAuthor:          errors.New("vocabulary_invalid_author"),
	InvalidDefinition:      errors.New("vocabulary_invalid_definition"),
	InvalidPartOfSpeech:    errors.New("vocabulary_invalid_part_of_speech"),
	InvalidIPA:             errors.New("vocabulary_invalid_ipa"),
	InvalidAudioName:       errors.New("vocabulary_invalid_audio_name"),
	InvalidExampleLanguage: errors.New("vocabulary_invalid_example_language"),
	InvalidExamplePosTags:  errors.New("vocabulary_invalid_example_pos_tags"),
}
