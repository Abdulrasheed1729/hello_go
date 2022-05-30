package main

import (
	"fmt"
)

func main() {
	c := Circle{0, 0, 5}
	fmt.Println(c.area())
	r := Rectangle{0, 0, 10, 10}
	fmt.Println(r.area())
}
