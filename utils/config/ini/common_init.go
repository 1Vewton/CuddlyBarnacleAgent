//go:build !test

package ini

func init() {
	iniConfigManager = newConfigManager(true)
}
