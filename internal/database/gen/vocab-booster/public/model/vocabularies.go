//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

import (
	"github.com/namhq1989/vocab-booster-english-hub/internal/database"
	"time"
)

type Vocabularies struct {
	ID            string `sql:"primary_key"`
	AuthorID      string
	Term          string
	PartsOfSpeech database.StringArray
	Ipa           string
	Audio         string
	Synonyms      database.StringArray
	Antonyms      database.StringArray
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Frequency     float64
}
