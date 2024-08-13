package main

import (
	"fmt"
	"os"
	"strconv"
	"two-moons-divi/division"
)

// args: 1, value, 2, preset optional - or json string

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("Please provide an argument - int or float value, preset string, json string {donation: float, cashflow: float, joint: float, owen: float, lola: float}")
		return
	}

	div := division.DefineDivision(args)
	// div.printReport(strconv.ParseFloat(args[0], 32))
	if sum, err := strconv.ParseFloat(args[0], 32); err == nil {
		div.PrintReport(float32(sum))
	} else {
		fmt.Println("Please provide a valid number")
	}

}
