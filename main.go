package main

import (
	"fmt"
	"time"

	"github.com/mgleon08/led-clock/placeholder"
	"github.com/inancgumus/screen"
)

func main() {
	screen.Clear()
	for {
		screen.MoveTopLeft()
		now := time.Now()
		hour, min, sec := now.Hour(), now.Minute(), now.Second()

		clock := [...][5]string{
			placeholder.Digits[hour/10], placeholder.Digits[hour%10],
			placeholder.Colon,
			placeholder.Digits[min/10], placeholder.Digits[min%10],
			placeholder.Colon,
			placeholder.Digits[sec/10], placeholder.Digits[sec%10],
		}

		for line := range clock[0] {
			for digit := range clock {
				fmt.Print(clock[digit][line], "  ")
			}
			fmt.Println()
		}

		time.Sleep(1 * time.Second)
	}

}
