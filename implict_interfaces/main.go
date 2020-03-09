package main

import "fmt"

type Duck interface {
	Quack()
}

type Donald struct {
}

func (d Donald) Quack() {
	fmt.Println("Quack Quack, Im DONALD THE DUCK!")
}

type Daisy struct {
}

func (d Daisy) Quack() {
	fmt.Println("Quack Quack, Im DAISY THE DUCK!")
}

func sayQuack(d Duck) {
	d.Quack()
}

func main() {
	donald := Donald{}
	sayQuack(donald)

	daisy := Daisy{}
	sayQuack(daisy)
}
