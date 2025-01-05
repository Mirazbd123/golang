package main

import "fmt"

func main() {
	fmt.Println("Welcome to pointers!")

	var ptr *int
	fmt.Println("The value of pointer : ", ptr)
	fmt.Println("The address of pointer : ", &ptr)

	myvalue := 22
	fmt.Println("The value of myvalue is : ", myvalue)

	ptr2 := &myvalue
	fmt.Println("The value of ptr2 : ", *ptr2)
	fmt.Println("The address of ptr2 : ", ptr2)

	*ptr2 = *ptr2 * 2
	fmt.Println("The new value of myvalue is : ", myvalue)
	fmt.Println("The address of muvalue is : ", ptr2)
}
