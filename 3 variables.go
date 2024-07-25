package main

import "fmt"

func variables() {
	var name string = "Riyad"
	var isMarried bool = false
	var age int = 28
	var salary uint = 1000
	var char byte = 1
	var runeVal rune = 1000
	var tax float32 = 20.20
	var com complex128 = 20.30

	fmt.Print(name, isMarried, age, salary, char, runeVal, tax, com)
}
