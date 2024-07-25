package main

import "fmt"

func range_test() {
	fmt.Println("Go range test")

	// traverse array with range

	// lopping over an array with range
	banana := [5]int{1, 2, 3, 4, 5}
	for i, ba := range banana {
		fmt.Println("Range", i, " -> ", ba)
	}

	// traverse slice with range
	guava := []string{"a", "b", "c"}
	for i, b := range guava {
		fmt.Println(i, " range ->", b)
	}
}
