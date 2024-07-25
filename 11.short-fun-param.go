package main

import "fmt"

func short_add(x, y, z int, p string) int {
	fmt.Println(p)
	return x + y + z
}

func short_function_params() {
	fmt.Println(short_add(10, 230, 30, "riyad"))
}
