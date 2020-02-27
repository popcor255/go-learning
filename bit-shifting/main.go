package main

import (
	"fmt"
)

func main()  {
	fmt.Println(1 << 1, ": 0 << 1 = 01")
	fmt.Println(8 >> 1, ": 1 >> 111= 11")
	fmt.Println(3 << 1, ": 0 >> 011= 0011")
}