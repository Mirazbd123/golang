package main

import "fmt"

const Publickeyword string = "github"

func main() {
	var username string = "Miraz"
	fmt.Printf("The type of username is : %T\n", username)
	fmt.Println("username is : ", username, "\n")

	var isBool bool = true
	fmt.Printf("The type of isBool is : %T\n", isBool)
	fmt.Println("isBool is : ", isBool, "\n")

	myInt := 56
	fmt.Printf("The type of myInt is : %T\n", myInt)
	fmt.Println("myInt is : ", myInt, "\n")

	fmt.Println(Publickeyword)
}
