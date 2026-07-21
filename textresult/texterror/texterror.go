package texterror

// TextError defines an error occurs in the text, including logic and other errors
type TextError struct {
	Type   ErrorTypes
	Line   int
	Reason string
}
