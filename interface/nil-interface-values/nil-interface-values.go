package main

import "fmt"

type I interface {
	M()
}

func describe(i I) {
	fmt.Printf("Value: %v, Type: %T", i, i)
}

func main() {
	var i I
	describe(i)
	i.M()
}
