package settings

import (
	"github.com/1Vewton/CuddlyBarnacleAgent/utils/logger"
)

// Config stores the configuration of this program.
var Config *settings
var configLogger *logger.Logger = logger.NewLogger(
	"config",
	nil,
)
