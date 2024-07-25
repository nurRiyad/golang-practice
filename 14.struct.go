package main

import "fmt"

type book struct {
	name        string
	author      string
	publishedAt string
}

func struct_test() {
	book1 := book{name: "api topu", author: "Jafor iqbal", publishedAt: "10-11"}

	fmt.Println(book1)
}
