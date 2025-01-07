package main

import "fmt"

func main() {
	fmt.Println("Hello , Sir welcome to the struct!!")
	// There is no inheritance and no super or parent in golang

	miraz := User{"Miraz", 25, "miraz@go.com", true}
	fmt.Println(miraz)
	fmt.Printf("Details about miraz: %+v\n", miraz)

	fmt.Printf("My name is %v and my email is %v\n", miraz.Name, miraz.Email)

}

type User struct {
	Name     string
	Age      int
	Email    string
	varified bool
}
