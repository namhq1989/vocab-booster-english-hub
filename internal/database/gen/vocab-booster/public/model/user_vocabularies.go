//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

import (
	"time"
)

type UserVocabularies struct {
	ID           string
	UserID       string `sql:"primary_key"`
	VocabularyID string `sql:"primary_key"`
	CreatedAt    time.Time
}
