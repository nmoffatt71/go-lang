package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")
	fmt.Println("`.-::::::-.`")
	fmt.Println(".:-::::::::::::::-:.")
	fmt.Println("`_:::    ::    :::_`")
	var line4 string = ".:( ^   :: ^   ):."
	fmt.Println(line4)
	integer := 23.1
	fmt.Printf("%v\n", integer)
	fmt.Printf("%v %d %x %o %b\n", integer, integer, integer, integer, integer)
	var flavorScale float32 = 5.8
	fmt.Println(flavorScale)

	// fmt.Println(math.Round(x*100) / 100)
	amount, unit := 10, "doll hairs"

	fmt.Println(amount, unit, ", that's expensive...")

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// var websites map[string]string
	websites := make(map[string]string)

	fmt.Println(len(websites))
	websites["google"] = "https://www.google.com"
	websites["microsoft"] = "https://www.microsoft.com"

	fmt.Println(websites)

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sum := 0
	for _, number := range numbers {
		sum += number
		println(sum)
	}
}
