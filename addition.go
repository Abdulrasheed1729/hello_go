package main

func add(args ...float64) float64 {
	total := 0.0
	for _, v := range args {
		total += v
	}
	return total
}
