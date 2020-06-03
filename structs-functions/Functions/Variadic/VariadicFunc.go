package main

import (
	"fmt"
)

func main() {
	primes := []int{2, 3, 5, 7, 11, 13}
	sum := sum(primes...)
	fmt.Println(sum)
	fmt.Println(primes)
}

func sum(x ...int) int {

	x[0] = 10
	fmt.Println(x)

	fmt.Printf("%T\n", x)

	fmt.Println(len(x))
	sum := 0

	for i := 0; i < len(x); i++ {
		sum += x[i]
	}
	return sum
}
