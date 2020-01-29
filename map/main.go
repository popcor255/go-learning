package main

import "fmt"

func main() {
	my_map := make(map[string]car)
	my_map["a"] = newCar("vroom", "blue")
	my_map["b"] = newCar("ferrrreeee", "red")
	my_map["c"] = newCar("land rooomer", "green")
	my_map["d"] = car{name: "woot", color : "gold"};
	fmt.Println(my_map)
}

func newCar(name, color string) (new_car car) {
	new_car.name = name
	new_car.color = color
	return
}

type car struct {
	name  string
	color string
}
