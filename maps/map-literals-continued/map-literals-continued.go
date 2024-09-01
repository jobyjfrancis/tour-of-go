package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs":  {20.345, 19.789},
	"Google LLC": {18.478, 45.390},
}

func main() {
	fmt.Println(m)
}
