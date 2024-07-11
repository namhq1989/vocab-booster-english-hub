package database

import "go.mongodb.org/mongo-driver/bson/primitive"

func NewStringID() string {
	return primitive.NewObjectID().Hex()
}

func ObjectIDFromString(id string) (primitive.ObjectID, error) {
	return primitive.ObjectIDFromHex(id)
}

func IsValidID(id string) bool {
	_, err := ObjectIDFromString(id)
	return err == nil
}
