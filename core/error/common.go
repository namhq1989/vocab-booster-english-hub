package apperrors

import "errors"

var Common = struct {
	Success      error
	BadRequest   error
	NotFound     error
	Unauthorized error
	Forbidden    error
	InvalidID    error
}{
	Success:      errors.New("success"),
	BadRequest:   errors.New("bad_request"),
	NotFound:     errors.New("not_found"),
	Unauthorized: errors.New("unauthorized"),
	Forbidden:    errors.New("forbidden"),
	InvalidID:    errors.New("invalid_id"),
}
