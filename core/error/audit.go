package apperrors

import "errors"

var Audit = struct {
	AuditNotFound error
	InvalidActor  error
	InvalidEntity error
	InvalidAction error
}{
	AuditNotFound: errors.New("audit_not_found"),
	InvalidActor:  errors.New("audit_invalid_actor"),
	InvalidEntity: errors.New("audit_invalid_entity"),
	InvalidAction: errors.New("audit_invalid_action"),
}
