package main

import "fmt"
import "math"

type Matrix struct {
	dimension int
	data      [][]int
}

func main() {
	matrix := newMatrix()
	matrix.getData()
	matrix.getDiagonalDifference()

	diagonalDifference := matrix.getDiagonalDifference()
	fmt.Printf("%d\n", diagonalDifference)
}

func newMatrix() *Matrix {
	m := new(Matrix)
	m.dimension = getMatrixDimension()

	m.data = make([][]int, m.dimension)

	for i := range m.data {
		m.data[i] = make([]int, m.dimension)
	}

	return m
}

func getMatrixDimension() int {
	var dimension int
	fmt.Scanf("%d", &dimension)
	return dimension
}

func (m *Matrix) getData() {
	for i := 0; i < m.dimension; i++ {
		for j := 0; j < m.dimension; j++ {
			_, err := fmt.Scanf("%d", &m.data[i][j])
			if err != nil {
				panic("FUCK")
			}
		}
	}
}

func (m *Matrix) getDiagonalDifference() int {
	primaryDiagonal, secondaryDiagonal := m.getPrimaryAndSecondaryDiagonals()
	return int(math.Abs(float64(primaryDiagonal) - float64(secondaryDiagonal)))
}

func (m *Matrix) getPrimaryAndSecondaryDiagonals() (int, int) {
	var primaryDiagonal int
	var secondaryDiagonal int

	distanceFromRight := m.dimension - 1

	for i := 0; i < m.dimension; i++ {
		// fmt.Printf("primaryDiagonal: %d secondaryDiagonal: %d\n", m.data[i][i], m.data[i][distanceFromRight])
		primaryDiagonal += m.data[i][i]
		secondaryDiagonal += m.data[i][distanceFromRight]
		distanceFromRight--
	}
	return primaryDiagonal, secondaryDiagonal
}
