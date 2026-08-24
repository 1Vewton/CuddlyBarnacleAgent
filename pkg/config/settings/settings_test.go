package settings

import (
	"testing"
)

// Test the env reading of string var
func TestEnvReadingString(t *testing.T) {
	var res string
	res = GetEnvString("test", "T")
	if res != "T" {
		t.Errorf("Expected T, got %s", res)
	}
	t.Setenv("test", "KK3S")
	res = GetEnvString("test", "T")
	if res != "KK3S" {
		t.Errorf("Expected KK3S, got %s", res)
	}
}

// Test the env setting
func TestEnvSettingString(t *testing.T) {
	var tSetting *settings = &settings{}
	t.Setenv("INI_FILE", "s.ini")
	res := tSetting.GetIniFile()
	if res != "s.ini" {
		t.Errorf("Expected s.ini, got %s", res)
	}
}

// Test Application env getting
func TestApplicationEnvGetting(t *testing.T) {
	var tSetting1 *settings = &settings{}
	t.Setenv("APPLICATION_MODE", "GUI")
	result := tSetting1.GetApplicationMode()
	if result != GUI {
		t.Errorf(
			"Expected %s, got %s",
			GUI.ToString(),
			result.ToString(),
		)
	}
	t.Setenv("APPLICATION_MODE", "AAA")
	var tSetting2 *settings = &settings{}
	result = tSetting2.GetApplicationMode()
	if result != CMD {
		t.Errorf(
			"Expected %s, got %s",
			CMD.ToString(),
			result.ToString(),
		)
	}
	t.Setenv("APPLICATION_MODE", "CMD")
	var tSetting3 *settings = &settings{}
	result = tSetting3.GetApplicationMode()
	if result != CMD {
		t.Errorf(
			"Expected %s, got %s",
			CMD.ToString(),
			result.ToString(),
		)
	}
}
