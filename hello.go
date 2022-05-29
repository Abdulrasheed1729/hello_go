package main

import "fmt"

func main() {

	for i := 1; i < 100; i++ {
		if i%3 == 0 {
			fmt.Printf("%d ", i)
		}
	}

	var x [5]int
	for i := 0; i < len(x); i++ {
		x[i] = i + 1
	}
	var total = 0
	for _, v := range x {
		total += v
	}

	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1, slice2)
	fmt.Println(cap(slice2))
	mySlice := make([]int, 3, 9)
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))

}
