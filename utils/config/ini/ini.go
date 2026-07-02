package ini

import (
	"github.com/1Vewton/CuddlyBarnacleAgent/utils/logger"
)

// IniConfig exports the config stored in the ini file
var IniConfig *ConfigStruct = &ConfigStruct{}
var iniConfigManager *configManager = &configManager{}
var configLogger *logger.Logger = logger.NewLogger(
	"Config",
	nil,
)
