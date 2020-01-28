package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int = 15
	var result string = ""
	for i := 1; i <= n; i++ {
		output := ""
		if i%3 == 0 {
			output += "fizz"
		}
		if i%5 == 0 {
			output += "buzz"
		}

		if output == "" {
			output += strconv.Itoa(i)
		}

		result += output + "\n"
	}

	fmt.Println(result)
}
