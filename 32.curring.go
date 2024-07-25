package main

import "fmt"

func zog(x, y int) int {
	return x + y
}

func gun(x, y int) int {
	return x * y
}

func selfMath(mathH func(int, int) int) func(int) int {
	return func(x int) int {
		return mathH(x, x)
	}
}

func currying_test() {
	fmt.Println("Testing currying in go")

	doubleFn := selfMath(zog)

	squareFun := selfMath(gun)

	fmt.Println(doubleFn(10), squareFun(5))
}
