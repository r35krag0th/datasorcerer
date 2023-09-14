package internal_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/r35krag0th/datasorcerer/internal"
)

func TestGetDatabaseConnection(t *testing.T) {
	db, err := internal.GetDatabaseConnection()
	assert.NoError(t, err)
	assert.NotNil(t, db)
}

func TestAutoMigrateModels(t *testing.T) {
	db, err := internal.GetDatabaseConnection()
	assert.NoError(t, err)

	err = internal.AutoMigrateModels(db)
	assert.NoError(t, err)
}
