package settings

import (
	"github.com/joho/godotenv"
)

// settings stores the config of the program
type settings struct {
	iniFile         *string
	applicationMode *string
}

// Initialize reads the env file setted
func (s *settings) Initialize(filePath string) error {
	err := godotenv.Load(filePath)
	return err
}

// GetIniFile gets the path of the configuration file
func (s *settings) GetIniFile() string {
	return SetConfigString(
		"INI_FILE",
		"conf.ini",
		&s.iniFile,
	)
}

// GetApplicationMode gets the mode of the application
func (s *settings) GetApplicationMode() ApplicationMode {
	return ToApplicationMode(
		SetConfigString(
			"APPLICATION_MODE",
			"CMD",
			&s.applicationMode,
		),
	)
}
