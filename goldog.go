package main

import (
	"fmt"
	"math"
	"os"
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

func main() {
	inputInfo, _ := os.Stdin.Stat()
	// var input []rune

	if inputInfo.Mode()&os.ModeCharDevice != 0 {
		warn := []rune("goldog works with pipe inputs.\nUsage: command | goldog")
		display(warn)
	}
}
