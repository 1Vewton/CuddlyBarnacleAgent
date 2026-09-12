package main

import (
	"github.com/1Vewton/CuddlyBarnacleAgent/internal/data/task"
	"github.com/1Vewton/CuddlyBarnacleAgent/internal/data/vectordb"
	"github.com/1Vewton/CuddlyBarnacleAgent/pkg/config/ini"
)

// Initialize
func init() {
	vectordb.VectorDataBase.InitializeDB()
	err := task.TaskManager.Load(
		ini.IniConfig.GetDataDir(),
		ini.IniConfig.GetTaskManagerFile(),
	)
	if err != nil {
		panic(err)
	}
}
