package main

import (
	"fmt"
	"os"
)

func main() {
	var choice int

	fmt.Println("Welcome")
	fmt.Println("Select an option")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Exit")

	fmt.Scan(&choice)

	if choice == 1 {
		fmt.Println("Chosen 1")
	} else if choice == 2 {
		fmt.Println("Chosen 2")
	} else {
		fmt.Println("Other")
	}

	fmt.Println("Chosen item: ", choice)

	for count2 := 0; count2 < 10; count2++ {
		fmt.Println(count2)
	}

	switch choice {
	case 1:
		fmt.Println("Chosen 1111")
	case 2:
		fmt.Println("Chosen 2222")
	}
	choice_t := fmt.Sprint(choice, " hello")
	os.WriteFile("neil.txt", []byte(choice_t), 0644)
}
