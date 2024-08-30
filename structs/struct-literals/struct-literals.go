package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{10, 20}
	v2 = Vertex{X: 100}
	v3 = Vertex{}
	p  = &Vertex{500, 1000}
)

func main() {
	fmt.Println(v1, v2, v3, p, *p)
}
