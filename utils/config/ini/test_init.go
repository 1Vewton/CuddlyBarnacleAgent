//go:build test

package ini

import (
	"github.com/1Vewton/CuddlyBarnacleAgent/utils/config/settings"
)

func init() {
	iniConfigManager = newConfigManager(
		false,
		settings.Settings.GetIniFile(),
	)
	IniConfig = NewConfigStruct(iniConfigManager)
}
