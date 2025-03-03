package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Welcome!")
	prices := []float64{10, 20.88, 30}
	taxrates := []float64{0, 0.0734, 0.1, 0.15}
	var total_prices float64 = 0.00
	var rounded_price float64 = 0.00

	result := make(map[float64][]float64)

	for _, taxrate := range taxrates {

		var taxIncludedPrices []float64 = make([]float64, len(prices))
		for count_price, price := range prices {
			// tax calc
			taxIncludedPrices[count_price] = price * (1 + taxrate)
			fmt.Println(taxIncludedPrices[count_price])
			// rax rounded
			ratio := math.Pow(10, float64(2))
			rounded_price := math.Round(taxIncludedPrices[count_price]*ratio) / ratio
			fmt.Println(rounded_price)
			// fmt.Println(taxIncludedPrices[count_price])
			total_prices += rounded_price

		}
		result[taxrate] = rounded_price
	}
	fmt.Println(total_prices)

}
