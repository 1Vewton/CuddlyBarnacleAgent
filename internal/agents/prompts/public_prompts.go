package prompts

const (
	// TextErrorDescription defines the problem of the text error
	TextErrorDescription string = `
This data structure defines one of the problem in the text provided by the user. 
	`
	// TextErrorArrayDescription defines the description for the array of text error
	TextErrorArrayDescription string = `
This contains all the problems you find in the text for this round. 
	`
	// OverwriteDataDescription describes the variable to control whether the provided text errors will be added to the existing data or overwrite the data
	OverwriteDataDescription string = `
If this value is set to true, the result you provided will overwrite the original data. 
	`
)
