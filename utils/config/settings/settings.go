package settings

// settings stores the config of the program
type settings struct {
	iniFile *string
}

// GetIniFile gets the path of the configuration file
func (s *settings) GetIniFile() string {
	return SetConfigString(
		"INI_FILE",
		"conf.ini",
		&s.iniFile,
	)
}
