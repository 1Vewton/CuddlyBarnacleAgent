//go:build !test

package ini

import (
	"github.com/1Vewton/CuddlyBarnacleAgent/utils/config/settings"
)

func init() {
	iniConfigManager = newConfigManager(
		true,
		settings.Settings.GetIniFile(),
	)
	iniConfigManager.InitializeConfig()
	IniConfig = NewConfigStruct(iniConfigManager)
}
