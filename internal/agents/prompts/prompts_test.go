package prompts

import (
	"testing"
)

const (
	testTipsPrompts string = "The user wants you to follow the following tips:\n1. abc\n"
)

// TestTipsPrompts tests the tips prompt
func TestTipsPrompts(
	t *testing.T,
) {
	t.Parallel()
	newTipsList := []string{
		"abc",
	}
	resultPrompt := GetTipsInfoCommand(newTipsList)
	if resultPrompt != testTipsPrompts {
		t.Errorf(
			"Expected %s, got %s",
			testTipsPrompts,
			resultPrompt,
		)
	}
}
