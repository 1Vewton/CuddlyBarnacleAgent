package agents

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
