package article

import (
	"os"
	"testing"
)

const testDataPath = "testdata/README.md"

// TestArticleFromFile tests creating article from file
func TestArticleFromFile(t *testing.T) {
	t.Parallel()
	tmpDir := t.TempDir()
	newArticle, err := CreateNewSingleArticle(
		testDataPath,
		tmpDir,
	)
	if err != nil {
		t.Error(err)
	}
	if newArticle.GetTitle() != "README" {
		t.Errorf(
			"expected title of the article to be %s, got %s",
			"README",
			newArticle.GetTitle(),
		)
	}
}

// TestShowArticle tests the article showing
func TestShowArticle(t *testing.T) {
	t.Parallel()
	// Create Article
	tmpDir := t.TempDir()
	article, err := CreateNewSingleArticle(
		testDataPath,
		tmpDir,
	)
	if err != nil {
		t.Error(err)
	}
	// Read expected split result
	data, err := os.ReadFile(
		"testdata/README_split.md",
	)
	if err != nil {
		t.Error(err)
	}
	stringSplitResult := string(data)
	if stringSplitResult != article.ShowLines() {
		t.Errorf(
			"expected %s, got %s",
			stringSplitResult,
			article.ShowLines(),
		)
	}
}
