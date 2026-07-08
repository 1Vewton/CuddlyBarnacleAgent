package settings

import (
	"github.com/joho/godotenv"
)

// settings stores the config of the program
type settings struct {
	iniFile *string
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
