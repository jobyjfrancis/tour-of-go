package main

import "fmt"

func main() {
	i := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(i)

	s := i[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)
}
