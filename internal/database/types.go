package database

import (
	"database/sql/driver"
	"strings"
)

type StringArray []string

func (s *StringArray) Scan(value interface{}) error {
	input := strings.Trim(value.(string), "{}")
	*s = strings.Split(input, ",")

	return nil
}

func (s *StringArray) Value() (driver.Value, error) {
	return "{" + strings.Join(*s, ",") + "}", nil
}

type CountResult struct {
	Total int64 `json:"total"`
}
