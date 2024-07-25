package main

import (
	"fmt"
)

func anonyms_struct() {
	fmt.Println("Anonyms Struct")

	book := struct {
		name   string
		author string
		size   struct {
			height int
			width  int
		}
	}{
		name:   "Riyad",
		author: "Mr. Riyad",
		size: struct {
			height int
			width  int
		}{height: 10, width: 20},
	}

	fmt.Println(book)

}
