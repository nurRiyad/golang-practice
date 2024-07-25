package main

import (
	"fmt"
)

type welDev struct {
	employee int
	place    string
}

func (w *welDev) changeEmployed(x int) {
	w.employee = x
}

func changePlace(place string, st *welDev) {
	st.place = place
}

func pointer_test() {
	fmt.Println("Testing Pointer in go")

	x := 10
	var ptr *int = &x
	fmt.Println(ptr, *ptr)
	*ptr = 100
	fmt.Println(ptr, *ptr)

	dev := welDev{employee: 140, place: "Banany"}
	fmt.Println(dev)

	changePlace("Uttora", &dev)
	fmt.Println(dev)

	dev.changeEmployed(1000)
	fmt.Println(dev)

}
