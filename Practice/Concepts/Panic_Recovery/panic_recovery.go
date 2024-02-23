package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(division(10, 5))
	fmt.Println(division(10, 2))
	fmt.Println(division(10, 0))
	fmt.Println(division(10, 10))
	fmt.Println(division(80, 5))
}

func division(numerator, denominator int) float64 {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	if denominator == 0 {
		panic(errors.New("denominator can't be 0"))
	}

	return float64(numerator) / float64(denominator)
}
