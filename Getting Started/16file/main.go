package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Hello , Sir welcome to the file system!!")
	content := "Hello, Sir welcome to the file system,go out and have fun!!"

	file, err := os.Create("./myfile.txt") // this will create a file in the current directory and return a file pointer and error
	// if err != nil {
	// 	panic(err)

	// }
	handleError(err)

	length, err := io.WriteString(file, content) // this will write the content to the file and return the length of the content and error
	handleError(err)
	fmt.Println("Length of the content is: ", length)
	file.Close()
	fmt.Printf("Type of file is %T\n", file) // *os.File , this is a pointer to the file
	fmt.Println(&file)                       // this will print the address of the file pointer

	readFile("./myfile.txt") // this will read the file and print the content of the file
}

func readFile(filename string) {
	dataByte, err := ioutil.ReadFile(filename) // this will read the file and return the data in the form of byte and error
	handleError(err)
	fmt.Println("Data in the file is: ", string(dataByte)) // this will print the data in the form of string
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
