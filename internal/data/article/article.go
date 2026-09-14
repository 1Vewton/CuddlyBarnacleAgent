package article

// SingleArticle defines the data structure for storing a single article
type SingleArticle struct {
	Lines          []string `json:"lines"`
	Title          string   `json:"title"`
	PreviewLine    string   `json:"preview_line"`
	TargetFilePath string   `json:"target_file_path"`
}

// CreateNewSingleArticle creates new single article
func CreateNewSingleArticle(
	filePath string,
) (*SingleArticle, error) {
	return nil, nil
}
