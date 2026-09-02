package textresult

import (
	"testing"

	"github.com/1Vewton/CuddlyBarnacleAgent/internal/agents/agenttypes"
)

// TestConvert tests the converting
func TestConvert(t *testing.T) {
	t.Parallel()
	testData := []TextError{
		{
			0,
			0,
			114514,
			"1919810",
		},
		{
			0,
			0,
			114514,
			"1919810",
		},
		{
			0,
			0,
			114514,
			"1919810",
		},
	}
	_, err := MultipleRawToProcessed(
		testData,
		agenttypes.FactualCheck,
	)
	if err != nil {
		t.Error(err)
	}
}
