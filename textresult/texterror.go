package textresult

// TextError defines an error occurs in the text, including logic and other errors
type TextError struct {
	Level  Levels
	Type   ErrorTypes
	Line   int
	Reason string
}
