package main

import "fmt"

func updater(x int) {
	defer fmt.Println("I am calling with the defer keyword")
	if x == 0 {
		fmt.Println("Zero")
	} else if x%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
}

func defer_test() {
	fmt.Println("Golang defer test")

	updater(0)
	updater(3)
	updater(10)

}
