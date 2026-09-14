package osoperation

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// PathToWindows converts path form to windows type
func PathToWindows(
	path string,
) string {
	return strings.ReplaceAll(
		path,
		`\`,
		"/",
	)
}

// CopyFileTo copies certain file to target directory.
// For targetFileName, no suffix needed.
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
	_, err = os.Stat(
		targetFilePath,
	)
	if os.IsExist(err) {
		return targetFilePath, fmt.Errorf(
			"%s already exists",
			targetFilePath,
		)
	}
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

// SplitToLines splits the text into lines
func SplitToLines(
	text string,
) []string {
	return strings.Split(
		text,
		"\n",
	)
}

// ReadFileToLines reads file and splits it to lines
func ReadFileToLines(
	filePath string,
) ([]string, error) {
	data, err := os.ReadFile(
		filePath,
	)
	if err != nil {
		return nil, err
	}
	dataToString := string(data)
	result := SplitToLines(
		dataToString,
	)
	return result, nil
}

// GetFileName gets the file name
func GetFileName(
	filePath string,
) string {
	base := filepath.Base(filePath)
	suffix := filepath.Ext(filePath)
	result := strings.ReplaceAll(
		base,
		suffix,
		"",
	)
	return result
}
