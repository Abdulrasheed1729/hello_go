package main

// This program takes in a slice of [int] as input and
// calculates the minimum of the values in the input.
//
// ```[go]
// Example:
// x := []int{10,3,7,8,12,}
// minOfSlice(x)
// ```
// will returns 3
func minOfSlice(slice []int) int {
	var min int
	for i, value := range slice {
		if i == 0 || value < min {
			min = value
		}
	}
	return min
}

// This program takes in a slice of [int] as input and
// calculates the maximum of the values in the input.
func maxOfSlice(slice []int) int {
	var min int
	for i, value := range slice {
		if i == 0 || value > min {
			min = value
		}
	}
	return min
}
