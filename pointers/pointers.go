package main

import "fmt"

func main() {
	i, j := 100, 200
	p := &i
	fmt.Println("i: ", i)
	fmt.Println("Address of i: ", p)
	fmt.Println("value of i through pointer p: ", *p)

	*p = 500 // changing value of i via pointer
	fmt.Println("\ni after changing value via pointer: ", i)
	fmt.Println("Address of i: ", p)
	fmt.Println("value of i through pointer p: ", *p)

	p = &j
	fmt.Println("\nj: ", j)
	fmt.Println("Address of j: ", p)
	fmt.Println("value of j through pointer p: ", *p)

	*p = *p / 100
	fmt.Println("\nj after changing value via pointer: ", j)
	fmt.Println("Address of j: ", p)
	fmt.Println("value of j through pointer p: ", *p)

}
