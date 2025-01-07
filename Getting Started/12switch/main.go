package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Hello , Sir welcome to the 12switch!!")
	// Switch case is a type of control flow
	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("The dice number is ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice number is 1")
	case 2:
		fmt.Println("Dice number is 2")
	case 3:
		fmt.Println("Dice number is 3")
	case 4:
		fmt.Println("Dice number is 4")
	case 5:
		fmt.Println("Dice number is 5")
	case 6:
		fmt.Println("Dice number is 6 and you won the game")
	default:
		fmt.Println("What kind of dice is that?")
	}
}
