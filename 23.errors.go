package main

import (
	"errors"
	"fmt"
)

func divide(x, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("CANNOT DIVIDE BY ZERO")
	}
	return x / y, nil
}

func error_test() {
	fmt.Println("Error test in go")

	x, err := divide(100, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(x)
}
