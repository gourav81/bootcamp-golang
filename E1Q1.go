package main

import (
	"encoding/json"
	"fmt"
)

type Matrix struct {
	rows    int
	columns int
	matrix  [][]int
}

func (a *Matrix) getRowsNumber() int {
	return a.rows
}

func (a *Matrix) getColumnsNumber() int {
	return a.columns
}

func (a *Matrix) setElement(i, j, value int) {
	a.matrix[i][j] = value
}

func (a *Matrix) addMatrices(b [][]int) {
	for i := 0; i < a.rows; i++ {
		for j := 0; j < a.columns; j++ {
			a.matrix[i][j] = a.matrix[i][j] + b[i][j]
		}
	}
}

func main() {
	n := 2
	m := 1
	mat1 := Matrix{rows: n, columns: m, matrix: make([][]int, n)}
	for i := 0; i < n; i++ {
		mat1.matrix[i] = make([]int, m)
	}
	mat1.setElement(0, 0, 5)
	mat1.setElement(1, 0, 2)

	mat2 := Matrix{rows: n, columns: m, matrix: make([][]int, n)}
	for i := 0; i < n; i++ {
		mat2.matrix[i] = make([]int, m)
	}
	mat2.setElement(0, 0, 5)
	mat2.setElement(1, 0, 2)
	mat1.addMatrices(mat2.matrix)
	jsonMatrix, _ := json.Marshal(mat1.matrix)
	fmt.Printf("%s\n", jsonMatrix)
}
