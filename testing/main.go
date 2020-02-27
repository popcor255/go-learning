package main

import "fmt"

func sum(n int) int {
	var result int
	for i := 0; i <= n; i++ {
		result += i
	}

	return result
}

func quickSum(n int) int {
	return (n * (n + 1)) / 2
}

func main(){
	fmt.Println(quickSum(5))
	fmt.Println(sum(5))
}