package main

import (
	"fmt"
	"math"
)

type about interface {
	show_shape() string
}

type cal interface {
	area() float64
	ttlArea(amount int) float64
}

type circle struct {
	name  string
	redis int
}

func (c circle) show_shape() string {
	return fmt.Sprintf("Circle named %v has redis of %v", c.name, c.redis)
}

func (c circle) area() float64 {
	return math.Pi * float64(c.redis) * float64(c.redis)
}

func (c circle) ttlArea(amount int) float64 {
	return float64(amount) * math.Pi * float64(c.redis) * float64(c.redis)
}

type square struct {
	name   string
	length int
}

func (s square) show_shape() string {
	return fmt.Sprintf("Square named %v has length of %v", s.name, s.length)
}

func (s square) area() float64 {
	return float64(s.length * s.length)
}

func (s square) ttlArea(amount int) float64 {
	return float64(s.length * s.length * amount)
}

func show(a about) {
	fmt.Println(a.show_shape())
}

func area(c cal) {
	fmt.Println(c.area())
}

func ttlArea(c cal, amount int) {
	fmt.Println(c.ttlArea(amount))
}

func multiple_interface() {
	fmt.Println("A struct an implement multiple interface")

	c := circle{name: "Lalbitto", redis: 10}
	s := square{name: "SadaSware", length: 20}

	fmt.Println(c, s)
	show(c)
	show(s)

	area(c)
	area(s)

	ttlArea(c, 10)
	ttlArea(s, 10)
}
