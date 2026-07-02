package ini

// ConfigStruct stores the configuration of the application
type ConfigStruct struct {
	llmModelName *string
}

// GetLLMModelName gets the LLM Model Name
func (cfg *ConfigStruct) GetLLMModelName() string {
	return iniConfigManager.GetConfigString(
		"LLM",
		"LLMModelName",
		"YOUR_MODEL_NAME",
	)
}
