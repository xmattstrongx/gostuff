package main

import (
	"fmt"
	"io"
	"math"
)

type Circular struct {
	N        int
	K        int
	Q        int
	A        []int
	M        []int
	rotatedA []int
}

func main() {
	c := Circular{}
	c.readInData()
	c.circularRotate()
	c.printSherlocksMQueries()
}

func (c *Circular) readInData() {
	c.getNKQ()
	c.getA()
	c.getM()
}

func (c *Circular) getNKQ() {
	fmt.Scanf("%d%d%d", &c.N, &c.K, &c.Q)
}

func (c *Circular) getA() {
	c.A = make([]int, c.N)

	for i := 0; i < c.N; i++ {
		fmt.Scanf("%d", &c.A[i])
	}
}

func (c *Circular) getM() {
	var current int

	for {
		_, err := fmt.Scanf("%d", &current)
		if err != nil {
			if err == io.EOF {
				break
			}
		}
		c.M = append(c.M, current)
	}
}

func (c *Circular) circularRotate() {
	length := len(c.A)
	shiftLength := c.K
	c.rotatedA = make([]int, length)

	for i, v := range c.A {
		if (i + shiftLength) <= (length - 1) {
			c.rotatedA[i+shiftLength] = v
		} else if shiftLength > length {
			c.rotatedA[c.calculateComplexNewLocation(i)] = v
		} else {
			shiftChange := i + shiftLength
			newLocation := shiftChange - (length)
			c.rotatedA[newLocation] = v
		}
	}
}

func (c *Circular) calculateComplexNewLocation(i int) int {
	remainder := c.K % len(c.A)
	newLocation := remainder + i
	if newLocation >= len(c.A) {
		newLocation = int(math.Abs(float64(newLocation - len(c.A))))
	}
	return newLocation
}

func (c *Circular) printSherlocksMQueries() {
	for _, v := range c.M {
		fmt.Printf("%d\n", c.rotatedA[v])
	}
}
