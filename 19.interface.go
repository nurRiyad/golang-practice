package main

import "fmt"

type msg interface {
	show() string
}

type animal struct {
	name  string
	speed string
	place string
}

func (a animal) show() string {
	return fmt.Sprintf("Animal name %v live in %v and run %v in a hour", a.name, a.place, a.speed)
}

type car struct {
	name   string
	brand  string
	milage string
}

func (c car) show() string {
	return fmt.Sprintf("Car name %v created by %v has milage of %v", c.name, c.brand, c.milage)
}

func PrintStruct(s msg) {
	fmt.Println(s.show())
}

func interface_test() {
	fmt.Println("Interface in go is just collection of method signature")

	toyota := car{name: "Corola", brand: "Toyota", milage: "20km"}
	tiger := animal{name: "RBT", place: "Sundorbar", speed: "120km/h"}

	fmt.Println(toyota, tiger)
	PrintStruct(toyota)
	PrintStruct(tiger)
}
