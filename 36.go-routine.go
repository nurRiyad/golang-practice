package main

import "fmt"

func go_routine() {
	fmt.Println("Go routine practice")

	go func() {
		fmt.Println("1 This is under go routine")
	}()

	fmt.Println("2 this is outside of go routine")
}
