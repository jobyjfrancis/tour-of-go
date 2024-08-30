package main

import "fmt"

func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}

	a := names[0:2]
	b := names[1:3]

	fmt.Println("names: ", names)
	fmt.Println("a: ", a)
	fmt.Println("b: ", b)

	b[0] = "XXX"

	fmt.Println("names: ", names)
	fmt.Println("a: ", a)
	fmt.Println("b: ", b)
}
