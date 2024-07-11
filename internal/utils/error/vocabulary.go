package apperrors

import "errors"

var Vocabulary = struct {
	VocabularyNotFound                error
	InvalidVocabularyID               error
	InvalidTerm                       error
	InvalidAuthor                     error
	InvalidDefinition                 error
	InvalidPartOfSpeech               error
	InvalidIPA                        error
	InvalidAudioName                  error
	InvalidExampleContent             error
	InvalidExampleTranslatedLanguages error
	InvalidExamplePosTags             error
	InvalidExampleLevel               error
	InvalidVerbConjugation            error
	InvalidVerbForm                   error
	InvalidSentence                   error
	SentenceIsAlreadyCorrect          error
	CannotPromoteDraftSentence        error
	SentenceIsTooSimple               error
}{
	VocabularyNotFound:                errors.New("vocabulary_not_found"),
	InvalidVocabularyID:               errors.New("vocabulary_invalid_id"),
	InvalidTerm:                       errors.New("vocabulary_invalid_term"),
	InvalidAuthor:                     errors.New("vocabulary_invalid_author"),
	InvalidDefinition:                 errors.New("vocabulary_invalid_definition"),
	InvalidPartOfSpeech:               errors.New("vocabulary_invalid_part_of_speech"),
	InvalidIPA:                        errors.New("vocabulary_invalid_ipa"),
	InvalidAudioName:                  errors.New("vocabulary_invalid_audio_name"),
	InvalidExampleContent:             errors.New("vocabulary_invalid_example_content"),
	InvalidExampleTranslatedLanguages: errors.New("vocabulary_invalid_example_translated_languages"),
	InvalidExamplePosTags:             errors.New("vocabulary_invalid_example_pos_tags"),
	InvalidExampleLevel:               errors.New("vocabulary_invalid_example_level"),
	InvalidVerbConjugation:            errors.New("vocabulary_invalid_verb_conjugation"),
	InvalidVerbForm:                   errors.New("vocabulary_invalid_verb_form"),
	InvalidSentence:                   errors.New("vocabulary_invalid_sentence"),
	SentenceIsAlreadyCorrect:          errors.New("vocabulary_sentence_is_already_correct"),
	CannotPromoteDraftSentence:        errors.New("vocabulary_cannot_promote_draft_sentence"),
	SentenceIsTooSimple:               errors.New("vocabulary_sentence_is_too_simple"),
}
