package main

import "fmt"

type Shape interface {
	area() float64
}
type Person struct {
	Name string
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}

type Android struct {
	Person
	Model string
}

func main() {
	a := new(Android)
	a.Talk()
}
