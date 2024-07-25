package main

import "fmt"

func divideBy(x, y int) (a, b int) {
	a = x / y
	b = x % y

	return
}

func name_return() {
	fmt.Println(divideBy(10, 3))
}
