package main

import (
	"fmt"
	"io"
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
	variance := c.K % len(c.A)
	if variance == 0 {
		c.rotatedA = c.A
		return
	}
	c.rotatedA = make([]int, len(c.A))
	diffFromVariance := len(c.A) - variance

	_ = copy(c.rotatedA[variance:], c.A[:diffFromVariance])
	_ = copy(c.rotatedA[:variance], c.A[diffFromVariance:])

}

func (c *Circular) printSherlocksMQueries() {
	for _, v := range c.M {
		fmt.Printf("%d\n", c.rotatedA[v])
	}
}
