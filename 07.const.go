package main

import "fmt"

func const_variable() {
	const minInHour = 60
	const secInMin = 60
	const senInHour = minInHour * secInMin
	fmt.Println(senInHour)
}
