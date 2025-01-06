package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Hello! Welcome to the slices in Golang")
	fruits := []string{}
	fruits = append(fruits, "Apple")
	fruits = append(fruits, "Banana")
	fmt.Printf("Type of fruits is : %T\n", fruits)
	fmt.Println("Fruits are : ", fruits)
	fmt.Println("Length of fruits is : ", len(fruits))

	fruits = append(fruits, "Mango", "Cheery")
	fmt.Println("Fruits are : ", fruits)
	fmt.Println("Length of fruits is : ", len(fruits))

	// Slicing the slice
	fruits_new := fruits[1:3]
	fmt.Println("Fruits are : ", fruits_new)
	fmt.Println("Length of fruits is : ", len(fruits_new))

	highscores := make([]int, 4)
	highscores[0] = 234
	highscores[1] = 345
	highscores[2] = 456
	highscores[3] = 567
	// highscores[4] = 678 // This will give error as the size of slice is 4
	fmt.Println("Highscores are : ", highscores)

	highscores = append(highscores, 123, 1235, 453, 23, 43)
	fmt.Println("Highscores are : ", highscores)
	if sort.IntsAreSorted(highscores) {
		fmt.Println("Highscores are sorted")
	} else {
		fmt.Println("Highscores are not sorted, So sort it!!")
		sort.Ints(highscores)
	}
	fmt.Println("Highscores are : ", highscores)

	// Removing elements from slice
	courses := []string{"Python", "Java", "Golang", "C++", "C"}
	sort.Strings(courses)
	fmt.Println("Courses in sorted array are : ", courses)
	index := 2 // Removing Golang
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println("Courses after removing Golang are : ", courses)

}
