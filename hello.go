package main

import "fmt"

func half(x int) (int, bool) {
	return x / 2, x%2 == 0
}

func greatest(args ...int) int {
	var max int = 0
	for _, v := range args {
		if v == 0 || v > max {
			max = v
		}
	}
	return max
}

func swap(x, y *int) {
	*x, *y = *y, *x

}

func square(x *float64) {
	*x = *x * *x
}

func main() {
	w := 1.5
	square(&w)
	x, y := 5, 6
	swap(&x, &y)

	fmt.Println(half(9))
	fmt.Println(x, y)
	fmt.Println(w)
}
