package main

import "fmt"

func main() {
	var my_map map[string]car
	my_map = make(map[string]car)
	my_map["a"] = newCar("vroom", "blue")
	my_map["b"] = newCar("ferrrreeee", "red")
	my_map["c"] = newCar("land rooomer", "green")
	my_map["d"] = car(honda{name: "woot", color: "gold"})

	my_map = map[string]car {
		"woot" : newCar("descoop", "poop"),
		"cool" : newCar("descoop", "poop"),
	}

	delete(my_map, "woot")

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

type honda struct {
	name  string
	color string
}
