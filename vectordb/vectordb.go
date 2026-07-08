package vectordb

import (
	"github.com/philippgille/chromem-go"
)

// VectorDB controls a vector database
type VectorDB struct {
	dbPath string
	db     *chromem.DB
}

// NewVectorDB creates a new VectorDB
func NewVectorDB(
	dbPath string,
) *VectorDB {
	return &VectorDB{
		dbPath: dbPath,
	}
}

// InitializeDB initializes a vector databse
