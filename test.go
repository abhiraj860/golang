package main

import "fmt"

var x int = 5

func f2() (int, int) {
	return 5, 6
}


func main() {
	add := func(x, y int) int {
		return x + y
	}
	fmt.Println(add(1, 2))
}

func average(xs [] float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}