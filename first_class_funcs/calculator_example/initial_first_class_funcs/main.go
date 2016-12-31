package main

import "fmt"
import "math"

type Calculator struct {
	acc float64
}

type opfunc func(float64, float64) float64

func (c *Calculator) Do(op opfunc, v float64) float64 {
	c.acc = op(c.acc, v)
	return c.acc
}

func Add(a, b float64) float64 {
	return a + b
}

func Sub(a, b float64) float64 {
	return a - b
}

func Mul(a, b float64) float64 {
	return a * b
}

func Div(a, b float64) float64 {
	return a / b
}

func main() {
	var c Calculator
	fmt.Println(c.Do(Add, 5))
	fmt.Println(c.Do(Sub, 3))
	fmt.Println(c.Do(Mul, 8))
	fmt.Println(c.Do(Div, 2))

	var c2 Calculator
	fmt.Println(c2.Do(Add, 16))
	fmt.Println(c2.Do(Sqrt, 0))
}

// In this method we can extend the calculator by only adding another function that can be of type opfunc

func Sqrt(n, _ float64) float64 {
	return math.Sqrt(n)
}
