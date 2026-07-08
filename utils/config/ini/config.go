package ini

// ConfigStruct stores the configuration of the application
type ConfigStruct struct {
	llmModelName *string
	LLMAPIKey    *string
	cfg          *configManager
}

// NewConfigStruct creates a new ConfigStruct
func NewConfigStruct(cfg *configManager) *ConfigStruct {
	return &ConfigStruct{
		cfg: cfg,
	}
}

// GetLLMModelName gets the LLM Model Name
func (cfg *ConfigStruct) GetLLMModelName() string {
	return cfg.cfg.GetConfigString(
		"LLM",
		"LLMModelName",
		"YOUR_MODEL_NAME",
	)
}

// GetLLMAPIKey gets the LLM Model Name
func (cfg *ConfigStruct) GetLLMAPIKey() string {
	return cfg.cfg.GetConfigString(
		"LLM",
		"LLMAPIKey",
		"YOUR_API_KEY",
	)
}

// GetLLMBaseURL gets the base url for the llm
func (cfg *ConfigStruct) GetLLMBaseURL() string {
	return cfg.cfg.GetConfigString(
		"LLM",
		"LLMAPIKey",
		"YOUR_API_KEY",
	)
}

// GetEmbeddingModelName gets the model name for the embedding model
func (cfg *ConfigStruct) GetEmbeddingModelName() string {
	return cfg.cfg.GetConfigString(
		"Embedding",
		"EmbeddingModelName",
		"YOUR_MODEL_NAME",
	)
}

// GetEmbeddingModelAPIKey gets the api key for the embedding model
func (cfg *ConfigStruct) GetEmbeddingModelAPIKey() string {
	return cfg.cfg.GetConfigString(
		"Embedding",
		"EmbeddingModelAPIKey",
		"YOUR_API_KEY",
	)
}

// GetEmbeddingModelBaseURL gets the base url for the embedding model
func (cfg *ConfigStruct) GetEmbeddingModelBaseURL() string {
	return cfg.cfg.GetConfigString(
		"Embedding",
		"EmbeddingModelBaseURL",
		"YOUR_BASE_URL",
	)
}
