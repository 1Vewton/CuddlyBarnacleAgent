package prompts

import (
	"github.com/1Vewton/CuddlyBarnacleAgent/internal/agents"
)

var TextResultParam = map[string]agents.ParameterDescription{
	"Line": {
		Description: "The position of the error",
		Required:    true,
	},
	"Level": {
		Description: "The reason you believe this part is an error",
		Required:    true,
	},
}
