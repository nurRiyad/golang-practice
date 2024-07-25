package main

import "fmt"

func slice_test() {
	fmt.Println("Slice are dynamic array in go")

	// full variable
	var apples []string
	apples = append(apples, "bananaApple")
	fmt.Println(apples)

	// short variable
	guava := []string{"a", "b", "c"}
	fmt.Println(guava)

	// using make
	zilla := make([]int, 10)
	fmt.Println(zilla)

	// array to slice
	arr := [4]int{1, 2, 3, 4}
	sls := arr[1:3]
	fmt.Println(sls)

	// traverse slice with loop
	for i := 0; i < len(guava); i++ {
		fmt.Println(i, "->", guava[i])
	}

	// traverse slice with range
	for i, b := range guava {
		fmt.Println(i, " range ->", b)
	}
}
