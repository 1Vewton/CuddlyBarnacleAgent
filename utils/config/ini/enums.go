package ini

// TextSplittingMethod defines the way of splitting the text
type TextSplittingMethod string

const (
	// Recursive defines the Recursive Splitting method
	Recursive TextSplittingMethod = "Recursive"
)

// ToString gets the string form of the method
func (method TextSplittingMethod) ToString() string {
	switch method {
	case Recursive:
		return "Recursive"
	default:
		return "N/A"
	}
}

// ToMethod converts the string to method
func ToMethod(method string) TextSplittingMethod {
	switch method {
	case "Recursive":
		return Recursive
	default:
		return Recursive
	}
}
