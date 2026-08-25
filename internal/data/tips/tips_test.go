package tips

import (
	"testing"
)

// TestFileOperation tests the operation on tips file
func TestFileOperation(
	t *testing.T,
) {
	testTips := NewTips(
		"testdata/test.json",
	)
	emptyTips := NewTips(
		"testdata/test.json",
	)
	defer emptyTips.SaveFile()
	errLoad := testTips.FromFile()
	if errLoad != nil {
		t.Error(errLoad)
	}
	testTips.NewTip("abc")
	testTips.NewTip("bcd")
	testTips.NewTip("ace")
	testTips.NewTip("ache")
	errSave := testTips.SaveFile()
	if errLoad != nil {
		t.Error(errSave)
	}
	reloadedTestTip := NewTips(
		"testdata/test.json",
	)
	errReloadedLoad := reloadedTestTip.FromFile()
	if errReloadedLoad != nil {
		t.Error(errReloadedLoad)
	}
	fetchedTips := reloadedTestTip.ReturnTipsThroughIdx(
		[]int{
			0,
			1,
			3,
			5,
		},
	)
	expectedResult := []string{
		"abc",
		"bcd",
		"ache",
	}
	for i, fetched := range fetchedTips {
		if i >= len(expectedResult) {
			t.Errorf(
				"%d: %s is not supposed to appear",
				i,
				fetched,
			)
		}
		if fetched != expectedResult[i] {
			t.Errorf(
				"In %d, expected %s, got %s",
				i,
				expectedResult[i],
				fetched,
			)
		}
	}
}
