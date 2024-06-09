package apperrors

import "errors"

var Staff = struct {
	InvalidStaffID error
	StaffNotFound  error
}{
	InvalidStaffID: errors.New("staff_invalid_id"),
	StaffNotFound:  errors.New("staff_not_found"),
}
