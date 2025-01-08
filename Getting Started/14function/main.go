package main

import "fmt"

func main() {
	fmt.Println("Hello , Sir welcome to the function sir!")
	fmt.Println("")

	sayHello()
	fmt.Println("")

	add := adder(2, 3)
	fmt.Println(add)
	fmt.Println("")

	proAdd := proAdder(1, 2, 3, 4, 5)
	fmt.Println(proAdd)
	fmt.Println("")

	fmt.Println("Using recursion!!!!!!")
	fmt.Println("Factorial of 5 is ", fact(5))
	fmt.Println("")

	a := 5
	b := 6
	fmt.Println("At first value of a and b is: ", a, " , ", b)
	fmt.Println("")

	valueA, valueB := PassByValue(a, b) // Pass by value
	fmt.Println("In Pass by value: ")
	fmt.Println("value A and Value B is: ", valueA, " , ", valueB)
	fmt.Println("value a and Value b is: ", a, " , ", b)
	fmt.Println("")

	referA, referB := PassByReferance(&a, &b) // Pass by referance
	fmt.Println("In Pass by Referance: ")
	fmt.Println("Referance A and Referance B is: ", referA, " , ", referB)
	fmt.Println("Referance a and Referance b is: ", a, " , ", b)

}

func sayHello() {
	fmt.Println("Hello")
}

func adder(a int, b int) int {
	fmt.Println("Adding two numbers", a, " and ", b)
	return a + b
}

func proAdder(val ...int) int { // here ... is called variadic parameter it will take slice of int
	fmt.Println("Adding multiple numbers")
	sum := 0
	for _, i := range val {
		sum += i
	}
	fmt.Printf("Type of val is %T\n", val)
	return sum
}

func fact(n int) int { // Recursion
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func PassByValue(a, b int) (int, int) {
	a = a * 2
	b = b * 2

	return a, b
}

func PassByReferance(a, b *int) (int, int) {
	*a = *a * 2
	*b = *b * 2

	return *a, *b
}
