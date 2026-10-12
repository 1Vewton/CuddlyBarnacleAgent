package agenttypes

// AgentType defines the type of agent
type AgentType int

const (
	// GrammaticalCheck (0) checks for grammar errors
	GrammaticalCheck AgentType = iota
	// LogicalCheck (1) checks for logical errors
	LogicalCheck
	// FactualCheck (2) checks for factual errors
	FactualCheck
	// TotalCheck checks for results of other checks
	TotalCheck
)

// ToString converts the AgentType to string
func (agentType AgentType) ToString() string {
	switch agentType {
	case GrammaticalCheck:
		return "GrammaticalCheck"
	case LogicalCheck:
		return "LogicalCheck"
	case FactualCheck:
		return "FactualCheck"
	case TotalCheck:
		return "TotalCheck"
	default:
		return "NotSupported"
	}
}
