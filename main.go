package main

import (
	"fmt"

	"github.com/mgleon08/led-clock/placeholder"
)

func main() {
	for line := 0; line < len(placeholder.Digits[0]); line++ {
		for digit := range placeholder.Digits {
			fmt.Print(placeholder.Digits[digit][line], "  ")
		}
		fmt.Println()
	}
}
