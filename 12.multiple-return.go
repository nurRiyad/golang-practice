package main

import "fmt"

func getCoordinate() (string, string) {
	return "nur", "riyad"
}

func multiple_return() {
	firstName, _ := getCoordinate()
	fmt.Println(firstName)
}
