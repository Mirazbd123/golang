package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println("Hello , Sir welcome to the loop!!")

	fmt.Println(strings.Repeat("-", 20))
	fmt.Println("Mathod 1: ")
	fmt.Println(strings.Repeat("-", 20))

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	sort.Strings(days)
	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}

	fmt.Println(strings.Repeat("-", 20))
	fmt.Println("Mathod 2: ")
	fmt.Println(strings.Repeat("-", 20))

	month := []string{"January", "February", "March", "April", "May", "June", "July", "August",
		"September", "October", "November", "December"}
	sort.Strings(month)

	for i := range month {
		fmt.Println(month[i])
	}
}
