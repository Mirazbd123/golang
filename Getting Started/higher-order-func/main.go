package main

import "fmt"

// higher order function
func applyOperation(a int, b int, operation func(int, int) int) int {
	return operation(a, b)
}

// first order function
func add(a int, b int) int {
	return a + b
}

func multiply(a int, b int) int {
	return a * b
}

func main() {
	sum := applyOperation(10, 20, add)
	fmt.Println("Sum:", sum)

	mul := applyOperation(10, 20, multiply)
	fmt.Println("Multiplication:", mul)
}
