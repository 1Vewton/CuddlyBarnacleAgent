package article

import (
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
