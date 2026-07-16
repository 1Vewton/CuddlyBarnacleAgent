package main

import (
	"fyne.io/fyne/v2/app"
	"github.com/1Vewton/CuddlyBarnacleAgent/cmd"
	"github.com/1Vewton/CuddlyBarnacleAgent/utils/config/settings"
)

// Main program
func main() {
	if settings.Settings.GetApplicationMode() == settings.GUI {
		a := app.New()
		w := a.NewWindow("CuddlyBarnacleAgent")
		w.ShowAndRun()
	} else {
		cmd.Execute()
	}
}
