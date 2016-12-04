package main

import "fmt"

type Matrix struct {
	dimension int
	data      [][]int
}

func main() {
	matrix := newMatrix(6)

	err := matrix.Read()
	if err != nil {
		panic("FUCK")
	}

	maxHourGlass := matrix.getMaxHoursGlass()
	fmt.Println(maxHourGlass)

}

func newMatrix(dimension int) *Matrix {
	m := new(Matrix)

	m.dimension = dimension

	m.data = make([][]int, m.dimension)

	for i := range m.data {
		m.data[i] = make([]int, m.dimension)
	}

	return m
}

func (m *Matrix) Read() error {
	for i := 0; i < m.dimension; i++ {
		for j := 0; j < m.dimension; j++ {
			_, err := fmt.Scanf("%d", &m.data[i][j])
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (m *Matrix) getMaxHoursGlass() int {
	var total int
	var first *int

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			currentSum := m.getHourGlassSum(i, j)
			if first == nil {
				total = currentSum
				first = &currentSum
			}
			if currentSum > total {
				total = currentSum
			}
		}
	}
	return total
}

func (m *Matrix) getHourGlassSum(i, j int) int {
	sum := 0
	sum += m.data[i][j]
	sum += m.data[i][j+1]
	sum += m.data[i][j+2]
	sum += m.data[i+1][j+1]
	sum += m.data[i+2][j]
	sum += m.data[i+2][j+1]
	sum += m.data[i+2][j+2]

	return sum
}
