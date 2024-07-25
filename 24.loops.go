package main

import "fmt"

func loops() {
	fmt.Println("Testing Loops")

	// for loop
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		if i == 5 {
			continue
		}
	}

	//while loop
	xy := 0
	for xy < 10 {
		fmt.Println("while", xy)
		if xy >= 5 {
			break
		}
		xy += 1
	}

	// infinite loop
	test := 0
	for {
		fmt.Println(test)
		test += 1

		if test >= 10 {
			break
		}
	}

}
