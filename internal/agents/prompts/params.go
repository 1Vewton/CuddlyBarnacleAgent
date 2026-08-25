package prompts

import (
	"github.com/1Vewton/CuddlyBarnacleAgent/internal/agents"
)

const (
	lineDescription  string = "The position of the error"
	levelDescription string = `
The level of this problem. 
- 0: Warning
- 1: Error
	`
	typeDescription string = `
The type of the problem. 
- 0: Uncategorized Error
	`
	reasonDescription string = "The reason you believe this part is an error"
)

// TextResultParam defines the param of text result
var TextResultParam = map[string]agents.ParameterDescription{
	"Line": {
		Description: lineDescription,
		Required:    true,
	},
	"Level": {
		Description: levelDescription,
		Required:    true,
	},
	"Type": {
		Description: typeDescription,
		Required:    true,
	},
	"Reason": {
		Description: reasonDescription,
		Required:    true,
	},
}
