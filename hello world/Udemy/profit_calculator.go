package main

import (
	"fmt"
	"math"
)

func main() {
	var revenue, expenses, tax_rate, ebt, profit, ratio float64
	fmt.Scan(&revenue, &expenses, &tax_rate)
	// Taxes
	ebt = revenue - expenses
	profit = ebt - (tax_rate * ebt)
	ratio = ebt / profit

	//Output
	fmt.Println("Earning before Tax: ", ebt)
	fmt.Println("Profit: ", profit)
	fmt.Println("Ratio", math.Round(ratio))
}
