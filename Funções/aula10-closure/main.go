package main

import "fmt"

func main() {
	a := i()
	fmt.Println(a())
}

func i() func() int {
	x := 0

	return func() int {
		x++
		return x
	}
}
