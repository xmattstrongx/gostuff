package main

import "fmt"
import "math"

type Calculator struct {
	acc float64
}

func (c *Calculator) Do(op func(float64) float64) float64 {
	c.acc = op(c.acc)
	return c.acc
}

func Add(n float64) func(float64) float64 {
	return func(acc float64) float64 {
		return acc + n
	}
}

func Sub(n float64) func(float64) float64 {
	return func(acc float64) float64 {
		return acc - n
	}
}

func Mul(n float64) func(float64) float64 {
	return func(acc float64) float64 {
		return acc * n
	}
}

func Div(n float64) func(float64) float64 {
	return func(acc float64) float64 {
		return acc / n
	}
}

func Sqrt() func(float64) float64 {
	return func(n float64) float64 {
		return math.Sqrt(n)
	}
}

func main() {
	var c Calculator
	fmt.Println(c.Do(Add(5)))
	fmt.Println(c.Do(Sub(3)))
	fmt.Println(c.Do(Mul(8)))
	fmt.Println(c.Do(Div(2)))

	var c2 Calculator
	fmt.Println(c2.Do(Add(16)))
	fmt.Println(c2.Do(Sqrt()))
}
