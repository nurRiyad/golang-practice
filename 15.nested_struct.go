package main

import "fmt"

type bookSize struct {
	height int
	width  int
	depth  int
}

type books struct {
	name      string
	author    string
	publishAt string
	size      bookSize
}

func nested_struct() {
	fmt.Println("Nested Struct")

	book1 := books{
		name:      "Himu",
		author:    "Humayon",
		publishAt: "2020",
		size: bookSize{
			height: 10,
			width:  20,
			depth:  30}}

	fmt.Println(book1)
}
