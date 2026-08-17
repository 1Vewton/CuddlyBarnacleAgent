//go:build !cli

package main

import (
	"github.com/1Vewton/CuddlyBarnacleAgent/cmd"
)

func main() {
	cmd.Execute()
}
