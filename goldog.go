package main

import (
	"fmt"
	"math"
)

func colorize(i int) (int, int, int) {
	var f = 0.1
	return int(math.Sin(f*float64(i)+0)*127 + 128),
		int(math.Sin(f*float64(i)+2*math.Pi/3)*127 + 128),
		int(math.Sin(f*float64(i)+4*math.Pi/3)*127 + 128)
}

func display(text []rune) {
	for j := 0; j < len(text); j++ {
		r, g, b := colorize(j)
		fmt.Printf("\033[38;2;%d;%d;%dm%c\033[0m", r, g, b, text[j])
	}
	fmt.Println()
}
