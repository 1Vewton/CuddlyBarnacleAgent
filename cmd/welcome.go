package cmd

import (
	"fmt"
)

// Welcome text
func Welcome() {
	fmt.Println(
		`
   ____   ____       _    
  / ___| | __ )     / \   
 | |     |  _ \    / _ \  
 | |___  | |_) |  / ___ \ 
  \____| |____/  /_/   \_\
		`,
	)
}
