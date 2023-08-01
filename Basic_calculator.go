package main

import (
	"fmt"
	// "math"
	// "string"
)

func main() {
	var a float32
	var b float32
	var c int

	fmt.Println("===Claculator===")
	fmt.Println("Enter first number: ")
	fmt.Scan(&a)
	fmt.Println("Enter second number: ")
	fmt.Scan(&b)
	fmt.Println("Which of the following operation you want to perform on above numbers: ")
	fmt.Println("1. Addition")
	fmt.Println("2. Subtraction")
	fmt.Println("3. Multiplication")
	fmt.Println("4. Division")
	fmt.Scan(&c)

	if c == 1 {
		add(a, b)
	} else if c == 2 {
		sub(a, b)
	} else if c == 3 {
		multiply(a, b)
	} else if c == 4 {
		divide(a, b)
	} else {
		fmt.Println("Enter Correct choice")
	}

}

func add(x float32, y float32) {
	fmt.Printf("Addition of %v and %v is: %v", x, y, x+y)

}

func sub(x float32, y float32) {
	fmt.Printf("Subtraction of %v and %v is: %v", x, y, x-y)

}

func multiply(x float32, y float32) {
	fmt.Printf("Multiplication of %v and %v is: %v", x, y, x*y)

}

func divide(x float32, y float32) {
	fmt.Printf("Division of %v and %v is: %v", x, y, x/y)

}
