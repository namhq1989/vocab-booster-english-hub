package apperrors

import "errors"

var Collection = struct {
	CollectionNotFound    error
	InvalidCollectionID   error
	MaxCollectionsReached error
	DescriptionTooLong    error
}{
	CollectionNotFound:    errors.New("collection_not_found"),
	InvalidCollectionID:   errors.New("collection_invalid_id"),
	MaxCollectionsReached: errors.New("collection_max_collections_reached"),
	DescriptionTooLong:    errors.New("collection_description_too_long"),
}
