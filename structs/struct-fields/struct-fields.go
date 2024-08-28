package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{10, 20}
	v.X = 20
	fmt.Println(v)
}
