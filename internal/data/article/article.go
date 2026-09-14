package article

import (
	"fmt"
	"os"

	"github.com/1Vewton/CuddlyBarnacleAgent/pkg/osoperation"
)

// SingleArticle defines the data structure for storing a single article
type SingleArticle struct {
	Lines          []string `json:"lines"`
	Title          string   `json:"title"`
	PreviewLine    string   `json:"preview_line"`
	TargetFilePath string   `json:"target_file_path"`
}

// CreateNewSingleArticle creates new single article
func CreateNewSingleArticle(
	originalFilePath string,
	targetDirName string,
) (*SingleArticle, error) {
	// Get title
	title := osoperation.GetFileName(
		originalFilePath,
	)
	// Copy file and separate to lines
	resultFilePath, err := osoperation.CopyFileTo(
		originalFilePath,
		title,
		targetDirName,
	)
	if err != nil {
		return nil, err
	}
	lines, err := osoperation.ReadFileToLines(
		resultFilePath,
	)
	if err != nil {
		return nil, err
	}
	// Get preview line
	var previewLine string
	if len(lines) < 1 {
		previewLine = title
	} else {
		previewLine = lines[0]
	}
	return &SingleArticle{
		PreviewLine:    previewLine,
		Title:          title,
		Lines:          lines,
		TargetFilePath: resultFilePath,
	}, nil
}

// Reload reloads the file from the single article
func (article *SingleArticle) Reload() error {
	_, err := os.Stat(article.TargetFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf(
				"%s file not exists",
				article.TargetFilePath,
			)
		}
		return err
	}
	lines, err := osoperation.ReadFileToLines(
		article.TargetFilePath,
	)
	if err != nil {
		return err
	}
	article.Lines = lines
	return nil
}
