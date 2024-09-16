package main

import "fmt"

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("Cannot do Sqrt for negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x > 0 {
		return 0, nil
	} else {
		return 1, ErrNegativeSqrt(x)
	}
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
