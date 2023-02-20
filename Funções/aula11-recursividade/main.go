package main

import "fmt"

func main() {
	fmt.Println(fatorial(5))
	fmt.Println(fatorialLoops(5))
}

func fatorial(x int) int {
	if x == 1 {
		return x
	}
	return x * fatorial(x-1)
}

func fatorialLoops(x int) int {
	total := 1
	for x > 1 {
		total *= x // total = total * x
		x--
	}
	return total
}
