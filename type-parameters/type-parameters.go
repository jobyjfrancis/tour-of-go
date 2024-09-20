package main

import "fmt"

func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		if v == x {
			return i
		}
	}
	return -1
}

func main() {
	si := []int{10, 20, 30, 40, 50}
	fmt.Println(Index(si, 90))

	ss := []string{"foo", "bar", "saz"}
	fmt.Println(Index(ss, "bar"))
}
