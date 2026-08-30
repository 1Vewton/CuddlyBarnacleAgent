package textresult

// ErrorType defines the types of errors
type ErrorType int

const (
	// UncategorizedError defines the error that is not categorized.
	// 0
	UncategorizedError ErrorType = iota
)

// GetAllPossibleErrors gets all the possible types of errors
func GetAllPossibleErrors() []ErrorType {
	return []ErrorType{
		UncategorizedError,
	}
}

// ToErrorType converts integer to error
func ToErrorType(data int) ErrorType {
	switch data {
	case 0:
		return UncategorizedError
	default:
		return UncategorizedError
	}
}

// ToInt converts error type to integer
func (errorType ErrorType) ToInt() int {
	switch errorType {
	case UncategorizedError:
		return 0
	default:
		return 0
	}
}

// ToString converts error type to string
func (errorType ErrorType) ToString() string {
	switch errorType {
	case UncategorizedError:
		return "UncategorizedError"
	default:
		return "UncategorizedError"
	}
}
