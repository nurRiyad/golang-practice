package main

import "fmt"

func adder() func(int) int {
	// this is closures
	add := 0
	// the returned function has ref of the add variable
	return func(i int) int {
		add += i
		return add
	}
}

func closures_test() {
	fmt.Println("Test closures in go")

	addFn := adder()
	fmt.Println(addFn(10))
	fmt.Println(addFn(20))
	fmt.Println(addFn(30))
	fmt.Println(addFn(40))
}
