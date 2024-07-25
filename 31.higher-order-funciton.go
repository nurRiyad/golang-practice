package main

import "fmt"

func addition(x, y int) int {
	return x + y
}

func subtract(x, y int) int {
	return x - y
}

// cal is higher order function
func aggregate(x, y, z int, cal func(int, int) int) int {
	return cal(x, cal(y, z))
}

func higher_order_function() {
	fmt.Println("Higher order function test ")
	fmt.Println(addition(1, 2))

	// here we used as anonymous function
	sum := aggregate(1, 2, 3, addition)
	sub := aggregate(20, 1, 1, subtract)

	fmt.Println(sum, sub)
}
