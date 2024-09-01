package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{40.33456, -17.8934}
	fmt.Println(m["Bell Labs"])

	// Checking what happens if elements are added to a Nil map
	/*var m1 map[string]string
	fmt.Println(m1)
	m1["Name"] = "Joby"
	m1["Age"] = "38"*/
}
