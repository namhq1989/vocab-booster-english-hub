//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

type UserExerciseCollectionStatus struct {
	ID                  string
	UserID              string `sql:"primary_key"`
	CollectionID        string `sql:"primary_key"`
	InteractedExercises int32
}
