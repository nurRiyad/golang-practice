package main

import "fmt"

func array_test() {
	fmt.Println("Testing array")

	// users
	var users [5]int
	users[0] = 1000

	// bananas
	banana := [5]int{1, 2, 3, 4, 5}

	// user := [4]int{1, 2, 3, 4}
	fmt.Println(users)
	fmt.Println(banana)

	// looping over array
	for i := 0; i < len(banana); i++ {
		fmt.Println(i, banana[i])
	}

	// lopping over an array with range
	for i, ba := range banana {
		fmt.Println("Range", i, " -> ", ba)
	}

}
