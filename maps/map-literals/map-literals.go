package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": Vertex{
		20.113, 48.983,
	},
	"Google": Vertex{
		12.345, 67.892,
	},
}

func main() {
	fmt.Println(m)
}
