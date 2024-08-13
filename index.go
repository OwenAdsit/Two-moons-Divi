package main

import (
	"fmt"
	"os"
	"strconv"
)

// args: 1, value, 2, preset optional - or json string

type Division struct {
	donation float32
	cashflow float32
	joint    float32
	owen     float32
	lola     float32
}

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("Please provide an argument - int or float value, preset string, json string {donation: float, cashflow: float, joint: float, owen: float, lola: float}")
		return
	}

	div := defineDivision(args)
	// div.printReport(strconv.ParseFloat(args[0], 32))
	if sum, err := strconv.ParseFloat(args[0], 32); err == nil {
		div.printReport(float32(sum))
	} else {
		fmt.Println("Please provide a valid number")
	}

}

func (div Division) printReport(sum float32) {
	donation := sum * div.donation
	cashflow := (sum - donation) * div.cashflow
	joint := (sum - donation - cashflow) * div.joint
	owen := (sum - donation - cashflow - joint) * div.owen
	lola := (sum - donation - cashflow - joint) * div.lola
	report := Division{
		donation: donation,
		cashflow: cashflow,
		joint:    joint,
		owen:     owen,
		lola:     lola,
	}
	// iterate over the struct and print the values
	fmt.Printf("Total: %.2f\n", sum)
	fmt.Printf("Donation: %.2f\n", report.donation)
	fmt.Printf("Cashflow: %.2f\n", report.cashflow)
	fmt.Printf("Joint: %.2f\n", report.joint)
	fmt.Printf("Owen: %.2f\n", report.owen)
	fmt.Printf("Lola: %.2f\n", report.lola)
}

func defineDivision(args []string) (div Division) {
	if len(args) == 1 {
		div = Division{
			donation: 0.03,
			cashflow: 0.1,
			joint:    0.15,
			owen:     0.6,
			lola:     0.4,
		}
	} else if len(args) == 2 {
		switch args[1] {
		case "branding":
			div = Division{
				donation: 0,
				cashflow: 0.1,
				joint:    0.15,
				owen:     0.05,
				lola:     0.95,
			}
		default:
			fmt.Println("Please provide a valid preset or no preset at all")
			div = Division{}
		}
	}
	return
}
