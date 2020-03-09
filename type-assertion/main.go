package main

import "fmt"

func main() {
	var variable interface{} = "64"
	_, ok := variable.(string)
	fmt.Println(ok)
	_, ok = variable.(float64)
	fmt.Println(ok)
}
