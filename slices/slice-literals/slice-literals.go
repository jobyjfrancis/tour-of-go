package main

import "fmt"

func main() {
	i := []int{10, 20, 30, 40, 50}
	fmt.Println(i)

	b := []bool{true, false, true, true, false, true}
	fmt.Println(b)

	s := []struct {
		i int
		b bool
	}{
		{10, true},
		{20, false},
		{30, true},
		{40, false},
		{50, true},
	}
	fmt.Println(s)
}
