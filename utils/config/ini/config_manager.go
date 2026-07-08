package ini

import (
	"os"

	"gopkg.in/ini.v1"
)

type configManager struct {
	createFile bool
	config     *ini.File
	configFile string
}

// newConfigManager creates new config manager
func newConfigManager(
	createFile bool,
	configFile string,
) *configManager {
	res := &configManager{
		createFile: createFile,
		config:     ini.Empty(),
		configFile: configFile,
	}
	err := res.InitializeConfig()
	if err != nil {
		panic(err.Error())
	}
	return res
}

// InitializeConfig initialize the config file
func (cfg *configManager) InitializeConfig() error {
	_, err := os.Stat(cfg.configFile)
	if os.IsNotExist(err) {
		configLogger.Info("Create new ini file")
		// Create new ini
		llmSettingsSection := cfg.config.Section("LLM")
		llmSettingsSection.NewKey("LLMModelName", "YOUR_MODEL_NAME")
		llmSettingsSection.NewKey("LLMAPIKey", "YOUR_API_KEY")
		llmSettingsSection.NewKey("LLMBaseURL", "YOUR_BASE_URL")
	} else {
		cfg.config, err = ini.Load(cfg.configFile)
		if err != nil {
			return err
		}
		// Fill the empty fields
		cfg.fillEmpty("LLM", "LLMModelName", "YOUR_MODEL_NAME")
		cfg.fillEmpty("LLM", "LLMAPIKey", "YOUR_API_KEY")
		cfg.fillEmpty("LLM", "LLMBaseURL", "YOUR_BASE_URL")
	}
	if cfg.createFile {
		cfg.config.SaveTo(cfg.configFile)
	}
	return nil
}

// fillEmpty fills the empty fields of config
func (cfg *configManager) fillEmpty(
	section string,
	key string,
	defaultValue string,
) {
	if !cfg.config.HasSection(section) || !cfg.config.Section(section).HasKey(key) {
		cfg.config.Section(section).NewKey(key, defaultValue)
		if cfg.createFile {
			cfg.config.SaveTo(cfg.configFile)
		}
	}
}

// RemoveAll removes all the fields of config
func (cfg *configManager) RemoveAll() {
	removedCfg := ini.Empty()
	removedCfg.SaveTo(cfg.configFile)
}

// GetConfigString get a field in config file
func (cfg *configManager) GetConfigString(
	section string,
	key string,
	defaultValue string,
) string {
	if !cfg.config.HasSection(section) ||
		!cfg.config.Section(section).HasKey(key) {
		cfg.config.Section(section).NewKey(key, defaultValue)
		if cfg.createFile {
			cfg.config.SaveTo(cfg.configFile)
		}
		return defaultValue
	}
	return cfg.config.Section(section).Key(key).String()
}
