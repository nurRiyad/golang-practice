package main

import "fmt"

func conditions() {

	if length := 20; length < 10 {
		fmt.Println("Too small")
	} else if length >= 10 && length < 30 {
		fmt.Println("Good size")
	} else {
		fmt.Println("Height is too big")
	}

	length := 30
	if length < 10 {
		fmt.Println("Too small")
	} else if length >= 10 && length < 30 {
		fmt.Println("Good size")
	} else {
		fmt.Println("Height is too big")
	}

}
