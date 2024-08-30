package main

import "fmt"

type Vertex struct {
	X, Y int
}

func main() {
	v := Vertex{10, 20}
	fmt.Println("Vertex before: ", v)
	p := &v
	p.X = 200
	fmt.Println("Vertex after: ", v)
}
