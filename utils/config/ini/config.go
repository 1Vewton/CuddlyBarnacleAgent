package ini

import (
	"errors"
	"fmt"

	"github.com/fatih/color"
)

// ConfigStruct stores the configuration of the application
type ConfigStruct struct {
	cfg *configManager
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

// SetLLMModelName sets the LLM Model Name
func (cfg *ConfigStruct) SetLLMModelName(
	newModelName string,
) error {
	cfg.cfg.SetConfigString(
		"LLM",
		"LLMModelName",
		newModelName,
	)
	return nil
}

// GetLLMAPIKey gets the LLM Model Name
func (cfg *ConfigStruct) GetLLMAPIKey() string {
	return cfg.cfg.GetConfigString(
		"LLM",
		"LLMAPIKey",
		"YOUR_API_KEY",
	)
}

// SetLLMAPIKey sets the API key of the LLM
func (cfg *ConfigStruct) SetLLMAPIKey(
	newAPIKey string,
) error {
	cfg.cfg.SetConfigString(
		"LLM",
		"LLMAPIKey",
		newAPIKey,
	)
	return nil
}

// GetLLMBaseURL gets the base url for the llm
func (cfg *ConfigStruct) GetLLMBaseURL() string {
	return cfg.cfg.GetConfigString(
		"LLM",
		"LLMBaseURL",
		"YOUR_BASE_URL",
	)
}

// SetLLMBaseURL sets the LLM Model Name
func (cfg *ConfigStruct) SetLLMBaseURL(
	newBaseURL string,
) error {
	cfg.cfg.SetConfigString(
		"LLM",
		"LLMBaseURL",
		newBaseURL,
	)
	return nil
}

// GetEmbeddingModelName gets the model name for the embedding model
func (cfg *ConfigStruct) GetEmbeddingModelName() string {
	return cfg.cfg.GetConfigString(
		"Embedding",
		"EmbeddingModelName",
		"YOUR_MODEL_NAME",
	)
}

// SetEmbeddingModelName sets the Embedding Model Name
func (cfg *ConfigStruct) SetEmbeddingModelName(
	newModelName string,
) error {
	cfg.cfg.SetConfigString(
		"Embedding",
		"EmbeddingModelName",
		newModelName,
	)
	return nil
}

// GetEmbeddingModelAPIKey gets the api key for the embedding model
func (cfg *ConfigStruct) GetEmbeddingModelAPIKey() string {
	return cfg.cfg.GetConfigString(
		"Embedding",
		"EmbeddingModelAPIKey",
		"YOUR_API_KEY",
	)
}

// SetEmbeddingModelAPIKey sets the API Key of the Embedding Model Name
func (cfg *ConfigStruct) SetEmbeddingModelAPIKey(
	newAPIKey string,
) error {
	cfg.cfg.SetConfigString(
		"Embedding",
		"EmbeddingModelAPIKey",
		newAPIKey,
	)
	return nil
}

// GetEmbeddingModelBaseURL gets the base url for the embedding model
func (cfg *ConfigStruct) GetEmbeddingModelBaseURL() string {
	return cfg.cfg.GetConfigString(
		"Embedding",
		"EmbeddingModelBaseURL",
		"YOUR_BASE_URL",
	)
}

// SetEmbeddingModelBaseURL sets the base url of the Embedding Model Name
func (cfg *ConfigStruct) SetEmbeddingModeBaseURL(
	newBaseURL string,
) error {
	cfg.cfg.SetConfigString(
		"Embedding",
		"EmbeddingModelAPIKey",
		newBaseURL,
	)
	return nil
}

// GetEmbeddingChunkSize gets the chunk size for textsplitter
func (cfg *ConfigStruct) GetEmbeddingChunkSize() (int, error) {
	return cfg.cfg.GetConfigInt(
		"Embedding",
		"EmbeddingChunkSize",
		200,
	)
}

// ShowAllConfig shows all the config used in CLI
func (cfg *ConfigStruct) ShowAllConfig() {
	header := color.New(color.FgBlue, color.Bold)
	header.Println("LLM")
	fmt.Printf(
		"LLM Model Name: %s\n",
		cfg.GetLLMModelName(),
	)
	fmt.Printf(
		"LLM API Key: %s\n",
		cfg.GetLLMAPIKey(),
	)
	fmt.Printf(
		"LLM Base URL: %s\n",
		cfg.GetLLMBaseURL(),
	)
	header.Println("Embedding")
	fmt.Printf(
		"Embedding Model Name: %s\n",
		cfg.GetEmbeddingModelName(),
	)
	fmt.Printf(
		"Embedding Model API Key: %s\n",
		cfg.GetEmbeddingModelAPIKey(),
	)
	fmt.Printf(
		"Embedding Model Base URL: %s\n",
		cfg.GetEmbeddingModelBaseURL(),
	)
}

// SetAnyConfig sets the configuration without needing to call single function
func (cfg *ConfigStruct) SetAnyConfig(
	field string,
	value string,
) error {
	switch field {
	case "LLMModelName":
		cfg.SetLLMModelName(value)
	case "LLMAPIKey":
		cfg.SetLLMAPIKey(value)
	case "LLMBaseURL":
		cfg.SetLLMBaseURL(value)
	case "EmbeddingModelName":
		cfg.SetEmbeddingModelName(value)
	case "EmbeddingModelAPIKey":
		cfg.SetEmbeddingModelAPIKey(value)
	case "EmbeddingModelBaseURL":
		cfg.SetEmbeddingModeBaseURL(value)
	default:
		return errors.New("No such field exists")
	}
	return nil
}
