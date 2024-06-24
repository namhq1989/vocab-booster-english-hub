package database

import (
	"github.com/segmentio/ksuid"
)

func NewStringID() string {
	return ksuid.New().String()
}

func IsValidID(value string) bool {
	_, err := ksuid.Parse(value)
	return err == nil
}
