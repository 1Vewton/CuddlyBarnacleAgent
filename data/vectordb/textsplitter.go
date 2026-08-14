package vectordb

import (
	"context"

	"github.com/1Vewton/CuddlyBarnacleAgent/utils/config/ini"
	"github.com/1Vewton/textsplitter"
	"github.com/1Vewton/textsplitter/fixedsplitter"
	"github.com/google/uuid"
	"github.com/philippgille/chromem-go"
)

// Splitter defines the text splitter
type Splitter struct {
	chunkSize int
	overlap   int
	method    ini.TextSplittingMethod
}

// GetSplitter gets the required text spliter
func (sc *Splitter) GetSplitter() textsplitter.TextSplitter {
	switch sc.method {
	case ini.Recursive:
		return fixedsplitter.NewFixedSplitter(
			sc.chunkSize,
			sc.overlap,
		)
	default:
		return fixedsplitter.NewFixedSplitter(
			sc.chunkSize,
			sc.overlap,
		)
	}
}

// SplitDoc splits the document
func (sc *Splitter) SplitDoc(
	ctx context.Context,
	document string,
) ([]chromem.Document, error) {
	result, err := sc.GetSplitter().SplitText(ctx, document)
	if err != nil {
		return nil, err
	}
	var documents []chromem.Document = []chromem.Document{}
	for _, i := range result {
		documents = append(
			documents,
			chromem.Document{
				Content: i,
				ID:      uuid.NewString(),
			},
		)
	}
	return documents, nil
}
