package main

import "fmt"

func describe(i interface{}) {
	fmt.Printf("Value: %v, Type: %T\n", i, i)
}

func main() {
	var i interface{}
	describe(i)

	i = 100
	describe(i)

	i = "Hello"
	describe(i)
}
