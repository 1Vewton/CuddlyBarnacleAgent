package vectordb

import (
	"context"
	"runtime"

	"github.com/1Vewton/CuddlyBarnacleAgent/pkg/config/settings"
	"github.com/google/uuid"
	"github.com/philippgille/chromem-go"
)

// VectorDataBase provides the public knowledge base entrypoint
var VectorDataBase = NewVectorDB(
	settings.Settings.GetKnowledgeBasePath(),
	"knowledgeBase",
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

// SetEmbeddingFunc sets the embedding function
func (vDB *VectorDB) SetEmbeddingFunc(
	embeddingFunc chromem.EmbeddingFunc,
) {
	vDB.embeddingFunc = embeddingFunc
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

// Query queries the related texts
func (vDB *VectorDB) Query(
	ctx context.Context,
	queryNum int,
	text string,
	title *string,
) ([]chromem.Result, error) {
	var queryMetadata map[string]string
	if title == nil {
		queryMetadata = nil
	} else {
		queryMetadata = map[string]string{
			"title": *title,
		}
	}
	return vDB.collection.Query(
		ctx,
		text,
		queryNum,
		queryMetadata,
		nil,
	)
}

// UploadArticleByLines uploads the article by lines
func (vDB *VectorDB) UploadArticleByLines(
	ctx context.Context,
	lines []string,
	title string,
) error {
	resultDocuments := []chromem.Document{}
	for _, line := range lines {
		id := uuid.NewString()
		tmpDocument := chromem.Document{
			ID:      id,
			Content: line,
			Metadata: map[string]string{
				"title": title,
			},
		}
		resultDocuments = append(resultDocuments, tmpDocument)
	}
	err := vDB.UploadDocuments(
		ctx,
		resultDocuments,
	)
	return err
}
