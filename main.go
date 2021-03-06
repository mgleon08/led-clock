package main

import (
	"fmt"
	"time"

	"github.com/mgleon08/led-clock/placeholder"
)

func main() {
	fmt.Print("\033[2J")
	for {
		fmt.Print("\033[H")
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
			for index, digit := range clock {
				if digit == placeholder.Colon && sec%2 == 0 {
					fmt.Print("   ", "  ")
				} else {
					fmt.Printf("\033[1;32m%v\033[m  ", clock[index][line])
				}
			}
			fmt.Println()
		}

		time.Sleep(1 * time.Second)
	}

}
