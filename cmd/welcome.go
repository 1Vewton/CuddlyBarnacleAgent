package cmd

import (
	"github.com/fatih/color"
)

// Welcome text
func Welcome() {
	printer := color.New(color.FgBlue, color.Bold)
	printer.Println(
		`
   ____   ____       _    
  / ___| | __ )     / \   
 | |     |  _ \    / _ \  
 | |___  | |_) |  / ___ \ 
  \____| |____/  /_/   \_\
		`,
	)
}
