package main

import (
	"fmt"
)

func main() {
	var arr []int = []int{2, 4, 8, 16}

	for i := range arr {
		fmt.Println(arr[i])
	}

	//_ is for a variable that you dont care about
	for _, v := range arr {
		fmt.Println(v)
	}
}
