package textresult

// TextError defines an error occurs in the text, including logic and other errors
type TextError struct {
	Level  int
	Type   int
	Line   int
	Reason string
}
