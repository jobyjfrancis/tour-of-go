package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v old)", p.Name, p.Age)
}

func main() {
	a := Person{"John", 52}
	b := Person{"Joby", 38}

	fmt.Println(a, b) //Note here that we don't have to call method like a.String(), b.String()
}
