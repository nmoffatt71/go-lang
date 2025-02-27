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
}
