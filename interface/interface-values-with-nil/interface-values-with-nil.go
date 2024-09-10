package main

import "fmt"

type I interface {
	M()
}

type T struct {
	s string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<NIL>")
		return
	}
	fmt.Println(t.s)
}

func describe(i I) {
	fmt.Printf("Value: %v, Type: %T\n", i, i)
}

func main() {
	var i I
	var t *T

	i = t
	describe(i)
	i.M()

	t = &T{"Hello"}
	i = t
	describe(i)
	i.M()
}
