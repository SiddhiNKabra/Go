package main

import (
	"fmt"
)

var num = 2

func main() {
	fmt.Println("Greetings!")

	switch num {
	case 1:
		fmt.Println("One")

	case 2:
		fmt.Println("Two")
	default:
		fmt.Println("None")
	}
	if num <= 5 {
		fmt.Println("Hii", num)
	}
	fmt.Println("Bye")

}
