package main

import "fmt"

func fibonacci() func() int {
	// Implement logic here using closures
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
