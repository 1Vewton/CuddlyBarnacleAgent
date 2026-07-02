package ini

import (
	"os"

	"github.com/1Vewton/CuddlyBarnacleAgent/utils/config/settings"
	"gopkg.in/ini.v1"
)

type configManager struct {
	createFile bool
	config     *ini.File
}

// Create new config manager
func newConfigManager(
	createFile bool,
) *configManager {
	return &configManager{
		createFile: createFile,
		config:     ini.Empty(),
	}
}

// InitializeConfig initialize the config file
func (cfg *configManager) InitializeConfig() error {
	_, err := os.Stat(settings.Settings.GetIniFile())
	if os.IsNotExist(err) {
		configLogger.Info("Create new ini file")
		// Create new ini
		llmSettingsSection := cfg.config.Section("LLM")
		llmSettingsSection.NewKey("LLMModelName", "YOUR_MODEL_NAME")
		llmSettingsSection.NewKey("LLMAPIKey", "YOUR_API_KEY")
		llmSettingsSection.NewKey("LLMBaseURL", "YOUR_BASE_URL")
		if cfg.createFile {
			cfg.config.SaveTo(settings.Settings.GetIniFile())
		}
	} else {
		cfg.config, err = ini.Load(settings.Settings.GetIniFile())
		if err != nil {
			return err
		}
	}
	return nil
}

// GetConfigString get a field in config file
func (cfg *configManager) GetConfigString(
	section string,
	key string,
	defaultValue string,
) string {
	if !cfg.config.HasSection(section) {
		cfg.config.Section(section).NewKey(key, defaultValue)
		return defaultValue
	}
	if !cfg.config.Section(section).HasKey(key) {
		return defaultValue
	}
	return cfg.config.Section(section).Key(key).String()
}
