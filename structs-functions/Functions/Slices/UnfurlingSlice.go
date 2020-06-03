package main

import "fmt"

func main() {
	xi := []int{1, 2, 3, 4, 5, 6, 7}
	//
	sum := sum(xi...)
	fmt.Println(sum)
	fmt.Println(xi)
}

func sum(x ...int) int {
	sum := 0
	x[1] = 10
	for v := range x {
		sum += v
	}
	return sum
}
