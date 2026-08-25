package tools

import (
	"context"
	"encoding/json"

	"github.com/1Vewton/CuddlyBarnacleAgent/internal/textresult"
)

// ProvideResult gets the result
func ProvideResult(
	ctx context.Context,
	argumentsInJSON string,
) (string, error) {
	var inputs []textresult.TextError
	errUnmarshal := json.Unmarshal(
		[]byte(argumentsInJSON),
		&inputs,
	)
	if errUnmarshal != nil {
		return "", nil
	}
	return "", nil
}
