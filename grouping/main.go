package main

import (
	"fmt"
)

type Animal struct {
	id int
	name string
}

type cat struct {
	Animal
	name string
}

type dog struct {
	Animal
	name string
}

func main(){
	bill := dog{Animal: Animal{1, "dog"}, name: "bill"}
	sam := cat{Animal: Animal{1, "cat"}, name: "sam"}
	fmt.Println(bill)
	fmt.Println(sam)
}