package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hellow sir, Welcome to our pizza shop!")
	fmt.Print("Please enter the number of slices you want to buy: ")

	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	input = strings.TrimSpace(input)
	fmt.Print("You've ordered ", input, " slices of pizza.\n")

	fmt.Println("Oh wait ! you are my new customer , so I give you an extra!")

	number, err := strconv.ParseInt(strings.TrimSpace(input), 10, 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Your pizza slice is : ", number+1)
	fmt.Printf("The type of pizza slice is : %T\n", number)

	fmt.Print("Would you like to order anything else? (yes/no): ")

	anything, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	anything = strings.TrimSpace(anything)
	if anything == strings.ToLower("no") {
		fmt.Println("Thank you for visiting our pizza shop!")
	} else {
		fmt.Println("Then visit tomorrow! Now get out and fuck life!")
	}

}
