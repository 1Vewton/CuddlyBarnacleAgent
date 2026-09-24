package article

import (
	"context"
	"fmt"

	"github.com/1Vewton/CuddlyBarnacleAgent/internal/data/vectordb"
)

// UploadConfig defines the struct for uploading article
type UploadConfig struct {
	ArticleManager *Manager
	VectorDB       *vectordb.VectorDB
}

// UploadArticle uploads single article
func (cfg *UploadConfig) UploadArticle(
	ctx context.Context,
	originalFilePath string,
	targetDirPath string,
) error {
	id, err := cfg.ArticleManager.AddNewArticle(
		originalFilePath,
		targetDirPath,
	)
	if err != nil {
		return err
	}
	articles := cfg.ArticleManager.GetAllArticles()
	article, exists := articles[id]
	if !exists {
		return fmt.Errorf(
			"article with %s id not exists",
			id,
		)
	}
	err = cfg.VectorDB.UploadArticleByLines(
		ctx,
		article.Lines,
		article.Title,
	)
	return err
}
