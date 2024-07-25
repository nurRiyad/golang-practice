package main

import (
	"fmt"
)

type about2 interface {
	show_shape() string
}

// circle type
type circle2 struct {
	name  string
	redis int
}

func (c circle2) show_shape() string {
	return fmt.Sprintf("Circle named %v has redis of %v", c.name, c.redis)
}

// square type
type square2 struct {
	name   string
	length int
}

func (s square2) show_shape() string {
	return fmt.Sprintf("Square named %v has length of %v", s.name, s.length)
}

func typeOfStruct1(a1 about2) string {
	switch v1 := a1.(type) {
	case square2:
		return v1.name
	case circle2:
		return v1.name
	default:
		return "None"
	}

}

func type_switch() {
	fmt.Println("Type asserting is used to find the type of the interface is. I.E find which struct is used to create the interface")

	c := circle2{name: "Lalbitto", redis: 10}
	s := square2{name: "SadaSware", length: 20}

	fmt.Println(c, s)
	show(c)
	show(s)

	fmt.Println(typeOfStruct1(c))
	fmt.Println(typeOfStruct1(s))

}
