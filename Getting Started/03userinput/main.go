package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Hello sir!"
	println(welcome)

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your name: ")
	input, err := reader.ReadString('\n') // use _ in err if you don't care about errors and here \n means
	// until you press enter it will keep reading the input
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Hello, " + input)

	fmt.Printf("Type of your name : %T\n", input)
}
