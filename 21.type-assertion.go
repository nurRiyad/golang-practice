package main

import (
	"fmt"
)

type about1 interface {
	show_shape() string
}

// circle type
type circle1 struct {
	name  string
	redis int
}

func (c circle1) show_shape() string {
	return fmt.Sprintf("Circle named %v has redis of %v", c.name, c.redis)
}

// square type
type square1 struct {
	name   string
	length int
}

func (s square1) show_shape() string {
	return fmt.Sprintf("Square named %v has length of %v", s.name, s.length)
}

func typeOfStruct(a1 about1) string {
	c1, ok := a1.(circle1)
	if ok {
		return "This is circle and it's name is " + c1.name
	}
	s1, ok := a1.(square1)
	if ok {
		return "This is square type and it's name if " + s1.name
	}
	return "None"
}

func type_assertion() {
	fmt.Println("Type asserting is used to find the type of the interface is. I.E find which struct is used to create the interface")

	c := circle1{name: "Lalbitto", redis: 10}
	s := square1{name: "SadaSware", length: 20}

	fmt.Println(c, s)
	show(c)
	show(s)

	fmt.Println(typeOfStruct(c))
	fmt.Println(typeOfStruct(s))

}
