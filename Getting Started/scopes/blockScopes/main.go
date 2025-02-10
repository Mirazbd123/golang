package main

import "fmt"

func main() {
	if true {
		y := 10
		fmt.Println("The value of this block scope is : ", y)
	}
	//fmt.Println("The vlaue of block scope that is outside the block is : ", y) // this line will show error bcz the value of y is removed by the
	// ram when the block (if) is finised.
}
