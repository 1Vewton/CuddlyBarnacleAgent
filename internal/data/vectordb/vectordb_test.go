package vectordb

import (
	"testing"
)

// Test the creation of database
func TestCreateDB(t *testing.T) {
	db := NewVectorDB(
		"./testdb",
		"test",
	)
	err := db.InitializeDB()
	if err != nil {
		t.Error(err.Error())
	}
}
