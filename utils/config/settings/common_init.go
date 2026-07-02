package settings

import (
	"github.com/1Vewton/CuddlyBarnacleAgent/utils/logger"
)

// Settings stores the basic settings of this program.
var Settings *settings = &settings{}
var configLogger *logger.Logger = logger.NewLogger(
	"config",
	nil,
)
