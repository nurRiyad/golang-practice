package main

import "fmt"

func string_formatter() {
	name := "riyad"
	fmt.Printf("Hello i am %v\n", name)
	msg := fmt.Sprintf("My age is %d", 20)
	fmt.Println(msg)
}
