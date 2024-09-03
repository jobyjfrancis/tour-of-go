package main

import "fmt"

func calculate() func() int {
	num := 0

	return func() int {
		num += 2
		return num
	}
}

func main() {
	even := calculate()

	fmt.Println(even())
	fmt.Println(even())
	fmt.Println(even())
	fmt.Println(even())
	fmt.Println(even())

	even2 := calculate()
	fmt.Println(even2())
}
