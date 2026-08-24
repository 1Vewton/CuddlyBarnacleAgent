package settings

// ApplicationMode defines the mode of running of the application
type ApplicationMode string

const (
	// CMD mode runs the application in terminal.
	CMD ApplicationMode = "CMD"
	// GUI mode runs the application with GUI.
	GUI ApplicationMode = "GUI"
)

// ToString converts ApplicationMode to string
func (mode ApplicationMode) ToString() string {
	switch mode {
	case CMD:
		return "CMD"
	case GUI:
		return "GUI"
	default:
		return "N/A"
	}
}

// ToApplicationMode converts string to ApplicationMode
func ToApplicationMode(mode string) ApplicationMode {
	switch mode {
	case "CMD":
		return CMD
	case "GUI":
		return GUI
	default:
		return CMD
	}
}
