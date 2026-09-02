package task

import (
	"testing"

	"github.com/1Vewton/CuddlyBarnacleAgent/pkg/osoperation"
)

// TestManagerFileOperation tests the file operation of task manager
func TestManagerFileOperation(t *testing.T) {
	rawTmpDir := t.TempDir()
	tmpDir := osoperation.PathToWindows(rawTmpDir)
	newManager := NewRawManager()
	err := newManager.Load(tmpDir, "test")
	if err != nil {
		t.Error(err)
	}
	newManager.Data["1"] = "abc"
	err = newManager.SaveFile()
	if err != nil {
		t.Error(err)
	}
	reloadedManager := NewRawManager()
	err = reloadedManager.Load(tmpDir, "test")
	if err != nil {
		t.Error(err)
	}
	if reloadedManager.Data["1"] != newManager.Data["1"] {
		t.Errorf(
			"expected %s, got %s",
			newManager.Data["1"],
			reloadedManager.Data["1"],
		)
	}
}
