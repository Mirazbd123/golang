package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var t int
	fmt.Fscan(reader, &t)

	for i := 0; i < t; i++ {
		var n string
		fmt.Fscan(reader, &n)

		if isDivisibleByNine(n) {
			fmt.Fprintln(writer, "YES")
		} else {
			fmt.Fprintln(writer, "NO")
		}
	}
}

func isDivisibleByNine(n string) bool {
	sum := 0

	for _, digit := range n {
		num := int(digit - '0')
		if num == 3 || num == 6 || num == 9 {
			return true
		}
		sum += num
	}

	return sum%9 == 0
}
