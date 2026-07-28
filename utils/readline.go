package utils

import (
	"fmt"
)

// ReadLine reads the inputs of string
func ReadLine(
	tip string,
	target ...*string,
) error {
	fmt.Println(tip)
	n, err := fmt.Scanln(target)
	if n != len(target) {
		return fmt.Errorf(
			"Expected %d target, got %d",
			len(target),
			n,
		)
	}
	return err
}
