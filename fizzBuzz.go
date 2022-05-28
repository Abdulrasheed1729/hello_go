package main

import "fmt"

// This program print a list of numbers between 1 and [limit]
// and it prints "Fizz" if a number is divisible by 3
// it returns "Buzz" if a number is divisible by 5 and returns
// "FizzBuzz" if a number is both divisible by 3 and 5
func FizzBuzz(limit int) {
	for i := 1; i < limit; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Printf("%d\n", i)
		}
	}
}
