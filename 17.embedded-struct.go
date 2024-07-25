package main

import "fmt"

type user struct {
	name   string
	age    int
	gender string
}

type student struct {
	user
	class   string
	academy string
	subject string
}

func embedded_struct() {
	fmt.Println("Embedded Struct")

	riyad := student{
		user: user{
			name:   "riyad",
			age:    24,
			gender: "male",
		},
		class:   "Bsc",
		academy: "IIUC",
		subject: "Cse",
	}

	fmt.Println(riyad, riyad.name, riyad.academy, riyad.academy)
}
