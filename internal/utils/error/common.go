package apperrors

import "errors"

var Common = struct {
	Success                   error
	BadRequest                error
	NotFound                  error
	Unauthorized              error
	Forbidden                 error
	AlreadyExisted            error
	InvalidID                 error
	InvalidName               error
	InvalidTense              error
	InvalidContent            error
	InvalidRequiredVocabulary error
}{
	Success:                   errors.New("success"),
	BadRequest:                errors.New("bad_request"),
	NotFound:                  errors.New("not_found"),
	Unauthorized:              errors.New("unauthorized"),
	Forbidden:                 errors.New("forbidden"),
	AlreadyExisted:            errors.New("already_existed"),
	InvalidID:                 errors.New("invalid_id"),
	InvalidName:               errors.New("invalid_name"),
	InvalidTense:              errors.New("invalid_tense"),
	InvalidContent:            errors.New("invalid_content"),
	InvalidRequiredVocabulary: errors.New("invalid_required_vocabulary"),
}
