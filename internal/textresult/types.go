package textresult

// ErrorTypes defines the types of errors
type ErrorType int

const (
	// UncategorizedError defines the error that is not categorized
	UncategorizedError ErrorType = iota
)

// ToErrorType converts integer to error
func ToErrorType(data int) ErrorType {
	switch data {
	case 0:
		return UncategorizedError
	default:
		return UncategorizedError
	}
}
