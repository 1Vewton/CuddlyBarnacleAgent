package textresult

// TextError defines an error occurs in the text, including logic and other errors
type TextError struct {
	Level  int
	Type   int
	Line   int
	Reason string
}

// ListResult defines the list of errors
type ListResult struct {
	Array                 []TextError
	OverWriteOriginalData bool
}

// StoredTextError defines the error stored in the program
type StoredTextError struct {
	Level  Levels
	Type   ErrorType
	Line   int
	Reason string
}
