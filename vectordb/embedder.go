package vectordb

import (
	"context"
	"errors"

	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
)

// Embedder data struct stores the config of embedder
type Embedder struct {
	client    *openai.Client
	modelName string
	dimension int64
}

// NewEmbedder creates a new embedder
func NewEmbedder(
	baseURL string,
	modelName string,
	APIKey string,
	dimension int64,
) *Embedder {
	client := openai.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey(APIKey),
	)
	return &Embedder{
		client:    &client,
		modelName: modelName,
		dimension: dimension,
	}
}

// EmbeddingFunc embeds the text
func (embedder *Embedder) EmbeddingFunc(
	context context.Context,
	text string,
) ([]float32, error) {
	res, err := embedder.client.Embeddings.New(
		context,
		openai.EmbeddingNewParams{
			Input: openai.EmbeddingNewParamsInputUnion{
				OfString: openai.String(text),
			},
			Model:          embedder.modelName,
			Dimensions:     openai.Int(embedder.dimension),
			EncodingFormat: openai.EmbeddingNewParamsEncodingFormatFloat,
		},
	)
	if err != nil {
		return nil, err
	}
	if len(res.Data) < 1 {
		return nil, errors.New("The returned data is empty")
	}
	result := make([]float32, len(res.Data[0].Embedding))
	for i := 0; i < len(res.Data[0].Embedding); i++ {
		result[i] = float32(res.Data[0].Embedding[i])
	}
	return result, nil
}
