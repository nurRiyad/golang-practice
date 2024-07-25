package main

import "fmt"

type food struct {
	name   string
	calory int
	taste  string
}

func (f food) about_food(count int) string {
	return fmt.Sprintf(
		"%v %v taste like %d and has %v calory\n", count, f.name, f.calory*count, f.taste)
}

func struct_method() {
	fmt.Println("Struct Method")

	apple := food{name: "Apple", calory: 120, taste: "Sweet"}

	fmt.Println(apple)
	msg := apple.about_food(10)
	fmt.Println(msg)
}
