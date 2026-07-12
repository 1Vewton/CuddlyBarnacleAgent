package main

import (
	"fyne.io/fyne/v2/app"
)

// Main program
func main() {
	a := app.New()
	w := a.NewWindow("CuddlyBarnacleAgent")
	w.ShowAndRun()
}
