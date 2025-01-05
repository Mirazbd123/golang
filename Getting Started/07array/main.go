package main

import "fmt"

func main() {
	fmt.Println("Welcome to array!!!!")
	var arr [5]int // Array declaration
	arr[0] = 10
	arr[1] = 20
	arr[2] = 30

	fmt.Println("Array declaration is : ", arr)
	fmt.Println("Array length is : ", len(arr)) // As here i declared length is 5

	var arrstr [5]string
	arrstr[0] = "Arman"
	arrstr[1] = "Miraz"
	arrstr[2] = "Hero"

	fmt.Println("Array declaration is : ", arrstr)
	fmt.Println("Array length is : ", len(arrstr)) // As here i declared

}
