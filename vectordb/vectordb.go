package vectordb

import (
	"context"
	"runtime"

	"github.com/philippgille/chromem-go"
)

// VectorDB controls a vector database
type VectorDB struct {
	dbPath         string
	db             *chromem.DB
	collectionName string
	collection     *chromem.Collection
	embeddingFunc  chromem.EmbeddingFunc
}

// NewVectorDB creates a new VectorDB
func NewVectorDB(
	dbPath string,
	collectionName string,
) *VectorDB {
	return &VectorDB{
		dbPath:         dbPath,
		collectionName: collectionName,
	}
}

// InitializeDB initializes a vector databse
func (vDB *VectorDB) InitializeDB() error {
	var errDB error
	var errCollection error
	vDB.db, errDB = chromem.NewPersistentDB(
		vDB.dbPath,
		true,
	)
	if errDB != nil {
		return errDB
	}
	vDB.collection, errCollection = vDB.db.GetOrCreateCollection(
		vDB.collectionName,
		nil,
		vDB.embeddingFunc,
	)
	if errCollection != nil {
		return errCollection
	}
	return nil
}

// UploadDocuments uploads chunked document list
func (vDB *VectorDB) UploadDocuments(
	ctx context.Context,
	documentList []chromem.Document,
) error {
	err := vDB.collection.AddDocuments(
		ctx,
		documentList,
		runtime.NumCPU(),
	)
	return err
}
