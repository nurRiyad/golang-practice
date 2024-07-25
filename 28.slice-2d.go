package main

import "fmt"

func makeMatrix(row, col int) [][]int {
	mat := make([][]int, 0)
	for i := 0; i < row; i++ {
		row := make([]int, 0)
		for j := 0; j < col; j++ {
			row = append(row, j)
		}
		mat = append(mat, row)
	}
	return mat
}

func slice_2d() {
	fmt.Println("2d Slice")

	mat := makeMatrix(3, 4)
	fmt.Println(mat)
}
