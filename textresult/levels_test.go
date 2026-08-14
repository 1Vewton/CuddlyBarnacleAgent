package textresult

import (
	"testing"
)

// Test to int function
func TestToInt(
	t *testing.T,
) {
	r1, _ := Warning.ToInt()
	if r1 != 0 {
		t.Errorf(
			"Expected %d, got %d",
			0,
			r1,
		)
	}
}
