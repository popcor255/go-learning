package main

import (
	"fmt"
)

type IP []byte
type IPMask []byte

type Time []byte

//unmarshal data uses pointer semantics
/* func (t *Time) UnmarshalBinary() error         {}
func (t *Time) GobDecode(data []byte) error    {}
func (t *Time) UnmarshalJSON(data []bye) error {}
func (t *Time) UnmarshalText(data []bye) error {}
*/
// Factory functions dictate the semantics that will be used. The Open function
type Person struct {
	Name string
	Age  int
}

func (p Person) Greet() {
	fmt.Printf("Hi! My name is %s", p.Name)
}

func NewPerson(name string, age int) Person {
	return Person{
		Name: name,
		Age:  age,
	}
}

func main() {
	//Use users as ref and change the values inside will still be using value semantics
	users := []Person{
		NewPerson("eric", 10),
		NewPerson("bill", 10),
	}

	users[0].Greet()
	fmt.Println(users[0].Name)

	//a reciever is the first paramter of a function as reference in order to
	// know where it should append
	//d.setAge(45) == (*data).setAge(&d. 45)
}
