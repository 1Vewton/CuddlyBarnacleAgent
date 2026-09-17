package article

import (
	"testing"
)

// TestLoadAndSave tests load and save
func TestLoadAndSave(
	t *testing.T,
) {
	t.Parallel()
	const fileName = "114.json"
	tmpDir := t.TempDir()
	newManager := NewManager()
	err := newManager.Load(
		tmpDir,
		fileName,
	)
	if err != nil {
		t.Error(err)
	}
	id, err := newManager.AddNewArticle(
		testDataPath,
		tmpDir,
	)
	if err != nil {
		t.Error(err)
	}
	result := newManager.GetAllArticles()
	_, exists := result[id]
	if !exists {
		t.Errorf(
			"%s article does not exists",
			id,
		)
	}
	err = newManager.Save()
	if err != nil {
		t.Error(err)
	}
	reloadedManager := NewManager()
	err = reloadedManager.Load(
		tmpDir,
		fileName,
	)
	if err != nil {
		t.Error(err)
	}
	result = reloadedManager.GetAllArticles()
	_, exists = result[id]
	if !exists {
		t.Errorf(
			"%s article does not exists",
			id,
		)
	}
}
