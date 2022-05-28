package main

import (
	"fmt"
)

func main() {
	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)
	output := input + 2
	fmt.Println(output)

	for i := 1; i < 100; i++ {
		if i%3 == 0 {
			fmt.Printf("%d ", i)
		}
	}
	FizzBuzz(100)
}
