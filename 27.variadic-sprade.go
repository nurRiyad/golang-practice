package main

import "fmt"

// variadic params
// it must be last params
func printAll(nums ...int) {
	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i], "->", i)
	}
}

func variadic_spread() {
	fmt.Println("This is variadic function")

	x := []int{1, 2, 3, 4}
	printAll()

	// spread operation
	printAll(x...)
	printAll(4, 4)
}
