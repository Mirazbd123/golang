package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hello , Sir welcome to the controlflow!!")
	// Process of taking input starts from here
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a number: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading number ", err)
	}
	input = strings.TrimSpace(input)
	num, _ := strconv.Atoi(input)
	// and end in here. Instead we can use scan function
	fmt.Printf("The type is %T\n", num)

	// we can use this for taking input
	fmt.Print("Enter another number : ")
	var num2 int
	_, err = fmt.Scan(&num2)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	if num%2 == 0 {
		fmt.Println("The number is even")
	} else {
		fmt.Println("The number is odd")
	}

	if num2%2 == 0 {
		fmt.Println("The new number is even")
	} else {
		fmt.Println("The new number is odd")
	}

	if numm := num + num2; numm%2 == 0 {
		fmt.Println("The sum is even")
	} else {
		fmt.Println("The sum is odd")
	}

}
