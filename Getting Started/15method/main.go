package main

import "fmt"

func main() {
	fmt.Println("Hello , Sir welcome to the struct!!")
	// There is no inheritance and no super or parent in golang

	miraz := User{"Miraz", 25, "miraz@go.com", false, "active"}
	fmt.Println(miraz)
	fmt.Printf("Details about miraz: %+v\n", miraz)

	fmt.Printf("My name is %v and my email is %v\n", miraz.Name, miraz.Email)
	fmt.Println("Is Miraz adult? ", miraz.isAdult())
	fmt.Println("Is Miraz active? ", miraz.getterIsActive())
	fmt.Println("status of Miraz is: ", miraz.status)
	miraz.setterActive("inactive")
	fmt.Println("new status of Miraz is: ", miraz.status)

}

type User struct {
	Name     string
	Age      int
	Email    string
	varified bool
	status   string
}

// key difference between the method and function is that method is associated with the struct and it take that as a receiver
//and function is not associated with the struct

func (r User) isAdult() string { // using value receiver for checking the age of the struct user --> Miraz
	if r.Age > 18 && r.varified {
		return "Yes"
	}
	return "NO"
}

func (r User) getterIsActive() string {
	if r.status == "active" {
		return "Yes"
	}
	return "No"
}

func (r *User) setterActive(status string) { // using pointer for updating the status property for the struct user --> Miraz
	r.status = status
}

// type Vertex struct {
// 	X int
// 	Y int
// }

// func main() {
// 	v := Vertex{1, 2}
// 	v.X = 4
// 	fmt.Println(v.X)
// }
