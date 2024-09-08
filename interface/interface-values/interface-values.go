package main

import (
	"fmt"
	"math"
)

type I interface {
	M()
}

type T struct {
	s string
}

func (t *T) M() {
	fmt.Println(t.s)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func describe(i I) {
	fmt.Printf("Value = %v, Type = %T\n", i, i)
}

func main() {
	var i I
	i = &T{"Hello"}
	describe(i)
	i.M()

	var f = F(math.Pi)
	describe(f)
	f.M()
}
