package apperrors

import "errors"

var UserVocabularyCollection = struct {
	CollectionNotFound    error
	InvalidCollectionID   error
	MaxCollectionsReached error
}{
	CollectionNotFound:    errors.New("user_vocabulary_collection_not_found"),
	InvalidCollectionID:   errors.New("user_vocabulary_collection_invalid_id"),
	MaxCollectionsReached: errors.New("user_vocabulary_collection_max_collections_reached"),
}
