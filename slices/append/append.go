package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v address=%p\n", len(s), cap(s), s, &s)
}

func main() {
	var s []int
	printSlice(s)

	s = append(s, 10)
	printSlice(s)

	s = append(s, 20, 30, 40, 50)
	printSlice(s)

	s1 := []int{1, 2, 3, 4, 5, 6, 7}
	printSlice(s1)

	s1 = append(s1, 8)
	printSlice(s1)
}
