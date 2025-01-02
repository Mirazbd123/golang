package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello sir welcome...")
	fmt.Println("Please see the Present time and date first.")
	presentTime := time.Now()
	fmt.Println(presentTime) // Printing the present time
	fmt.Println("I don't understand this time format")

	// Converting the time to desired format

	fmt.Println(presentTime.Format("02-01-2006 15:04:05 Monday")) // always have to remember this format
	// 02 is for date, 01 is for month and 2006 is for year
	fmt.Println("Is that ok sir!!")

	fmt.Println("Yes , it is !")

	fmt.Println("Now tell me why do you come here at that time ? Don't you have any common sence!!")

	// Here in terminal if i run go env it will show me a bunch of things where i can get help . But if i
	// want to build a executable file then i have to run go build and if i want to build a file for windows
	// or any os than i have to run GOOS='windows' go build and it will create a windows executable file.
	// Here mytime file is for linux and mytime.exe is for windows.
}
