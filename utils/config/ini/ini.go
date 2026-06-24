package ini

import (
	"gopkg.in/ini.v1"
)

// InitializeConfig initialize the config file
func InitializeConfig() {
	cfg := ini.Empty()
	llmSettingsSection := cfg.Section("LLM")
	llmSettingsSection.NewKey("LLMModelName", "YOUR_MODEL_NAME")
}
