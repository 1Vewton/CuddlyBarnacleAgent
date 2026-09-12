package article

import (
	"fmt"
	"os"
	"path/filepath"
)

// CopyFileTo copies certain file to target directory
func CopyFileTo(
	originalFilePath string,
	targetFileName string,
	targetDirName string,
) (string, error) {
	data, err := os.ReadFile(originalFilePath)
	if err != nil {
		return "", err
	}
	suffix := filepath.Ext(originalFilePath)
	targetFilePath := fmt.Sprintf(
		"%s/%s.%s",
		targetDirName,
		targetFileName,
		suffix,
	)
	err = os.WriteFile(
		targetFilePath,
		data,
		0644,
	)
	if err != nil {
		return targetFilePath, err
	}
	return targetFilePath, nil
}
