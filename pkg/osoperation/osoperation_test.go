package osoperation

import (
	"testing"
)

const testDataPath = "testdata/游泳馆.md"

// TestCopyFile tests the file copying
func TestCopyFile(
	t *testing.T,
) {
	t.Parallel()
	tmpDir := t.TempDir()
	_, err := CopyFileTo(
		testDataPath,
		"test",
		tmpDir,
	)
	if err != nil {
		t.Error(err)
	}
}

// TestGetFileName tests the file name fetching from path function
func TestGetFileName(
	t *testing.T,
) {
	t.Parallel()
	res := GetFileName(
		testDataPath,
	)
	if res != "游泳馆" {
		t.Errorf(
			"expected %s, got %s",
			"游泳馆",
			res,
		)
	}
}

// TestToLines tests the read file to lines function
func TestToLines(
	t *testing.T,
) {
	t.Parallel()
	lines, err := ReadFileToLines(
		testDataPath,
	)
	if err != nil {
		t.Error(err)
	}
	if len(lines) != 31 {
		t.Errorf(
			"expected total line number to be %d, got %d",
			31,
			len(lines),
		)
	}
}
