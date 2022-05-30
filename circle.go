package main

import "math"

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

type Circle struct {
	x, y, r float64
}
