package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSQLiteStorage(t *testing.T) {
	// Test with a valid database file
	// dbName := "valid_database.db"
	storage, err := NewSQLiteStorage()
	assert.NoError(t, err)
	assert.NotNil(t, storage)
	assert.NotNil(t, storage.db)
}

// func TestNewSQLiteStorageError(t *testing.T) {
// 	// Test with an invalid database file
// 	// dbName := "invalid_database.txt"
// 	storage, err := NewSQLiteStorage()
// 	assert.Error(t, err)
// 	assert.Nil(t, storage)
// 	// assert.True(t, errors.Is(err, errors.New("invalid database file")))
// }