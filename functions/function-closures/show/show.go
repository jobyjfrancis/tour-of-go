package main

import "fmt"

func show() func() int {
	i := 0

	return func() int {
		i += 10
		return i
	}
}

func main() {
	a := show()
	fmt.Println("a=", a())

	b := show()
	fmt.Println("b=", b())

	b()

	fmt.Println("a=", a())
	fmt.Println("b=", b())
}
