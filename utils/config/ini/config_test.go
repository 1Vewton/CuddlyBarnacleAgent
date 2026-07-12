package ini

import (
	"testing"
)

// Test Initialization
func TestInitialization(t *testing.T) {
	NewConfig := newConfigManager(
		true,
		"testdata/test.ini",
	)
	defer NewConfig.RemoveAll()
	err := NewConfig.InitializeConfig()
	if err != nil {
		t.Fatal(err.Error())
	}
	result := NewConfig.GetConfigString(
		"LLM",
		"LLMModelName",
		"YOUR_MODEL_NAMES",
	)
	if result != "YOUR_MODEL_NAME" {
		t.Errorf("Expected YOUR_MODEL_NAME, got %s", result)
	}
}

// Test GetConfigString
func TestConfigString(t *testing.T) {
	NewConfig := newConfigManager(
		true,
		"testdata/test.ini",
	)
	defer NewConfig.RemoveAll()
	err := NewConfig.InitializeConfig()
	if err != nil {
		t.Fatal(err.Error())
	}
	NewConfig.GetConfigString(
		"test",
		"test1",
		"test3",
	)
	result := NewConfig.GetConfigString(
		"test",
		"test1",
		"test4",
	)
	if result != "test3" {
		t.Errorf("Expected test3, got %s", result)
	}
}

// Test the fillEmpty function
func TestFillEmpty(t *testing.T) {
	NewConfig := newConfigManager(
		true,
		"testdata/test.ini",
	)
	defer NewConfig.RemoveAll()
	err := NewConfig.InitializeConfig()
	if err != nil {
		t.Fatal(err.Error())
	}
	NewConfig.fillEmpty(
		"test",
		"test1",
		"test2",
	)
	result := NewConfig.GetConfigString(
		"test",
		"test1",
		"test3",
	)
	t.Log(result)
	if result != "test2" {
		t.Errorf("Expected test2, got %s", result)
	}
}

func TestSetConfig(t *testing.T) {
	NewConfig := newConfigManager(
		true,
		"testdata/test.ini",
	)
	defer NewConfig.RemoveAll()
	err := NewConfig.InitializeConfig()
	if err != nil {
		t.Fatal(err.Error())
	}
	NewConfig.SetConfigString(
		"test",
		"test",
		"1",
	)
	result := NewConfig.GetConfigString(
		"test",
		"test",
		"2",
	)
	if result != "1" {
		t.Errorf("Expected 1, got %s", result)
	}
}
