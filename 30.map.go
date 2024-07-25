package main

import "fmt"

func map_test() {
	// map key can be bool,number,string,struct
	// map key can not be slice function map
	// map value can be anything
	fmt.Println("Testing map in go lang")

	// init map type 1
	user := make(map[string]int)
	user["ashik"] = 1
	user["riyad"] = 2
	fmt.Println(user)

	// init map type 2
	country := map[string]string{
		"bangladesh": "dhaka",
		"india":      "dilli",
	}
	fmt.Println(country)

	// map length
	countryLen := len(country)
	fmt.Println(countryLen)

	// traverse through an map
	for i, cap := range country {
		fmt.Println(i, cap)
	}

	// insert an element in map
	country["nepal"] = "kathmudu"

	// get a value from map
	mapVal, ok := country["sdklfj"]
	mapVal1 := country["nepal"]

	fmt.Println(ok, "->", mapVal)
	fmt.Println(mapVal1)

	// delete a key
	delete(country, "nepal")
	fmt.Println(country)

	type animal struct {
		name string
		leg  int
	}

	animals := map[animal]string{
		{name: "cow", leg: 20}:  "Hi I am an cow",
		{name: "snake", leg: 0}: "Mewao",
	}
	fmt.Println(animals)
	fmt.Println(animals[animal{name: "cow", leg: 20}])

}
