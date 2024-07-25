package main

import "fmt"

func static_type() {
	var username string = "riyad"
	var password string = "1233445"

	fmt.Println("Basic Authentication ", username+":"+password)
}
