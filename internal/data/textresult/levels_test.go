package textresult

import (
	"testing"
)

// TestToInt tests to int function
func TestToInt(
	t *testing.T,
) {
	t.Parallel()
	r1, _ := Warning.ToInt()
	if r1 != 0 {
		t.Errorf(
			"Expected %d, got %d",
			0,
			r1,
		)
	}
	_, err := ToLevels(114514)
	if err == nil {
		t.Error("Expected to return error")
	}
}

// TestToLevels tests to level function
func TestToLevels(t *testing.T) {
	level, _ := ToLevels(0)
	if level != Warning {
		t.Errorf(
			"Expected %d, got %d",
			Warning,
			level,
		)
	}
	_, err := ToLevels(-1)
	if err == nil {
		t.Error("Expected to return error")
	}
}
