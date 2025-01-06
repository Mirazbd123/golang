package main

import "fmt"

func main() {
	fmt.Println("Hello! Welcome to the maps in Golang")
	// Creating a map
	grade := make(map[string]int)
	grade["John"] = 90
	grade["Alice"] = 80
	grade["Bob"] = 85
	grade["Charlie"] = 95
	grade["John"] = 92
	fmt.Println("Grades are : ", grade)
	fmt.Printf("Type of grade is : %T\n", grade)
	// Accessing a value
	fmt.Println("John's grade is : ", grade["John"])
	//deleting an element
	delete(grade, "Alice")
	fmt.Println("Grades after deleting Alice : ", grade)

	// Iterating over a map
	fmt.Println("Iterating over a map")
	for key, value := range grade {
		fmt.Println(key, ":", value)
	}

}
