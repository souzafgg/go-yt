package main

import "fmt"

func main() {
	n := []int{10, 20, 30, 40, 50}
	fmt.Println(intVariadico(n...))

	s := somaSlice([]int{10, 20, 30, 40, 50})
	fmt.Println(s)
}

func intVariadico(x ...int) int {
	t := 0
	for _, v := range x {
		t += v
	}
	return t
}

func somaSlice(x []int) int {
	n := 0
	for _, v := range x {
		n += v
	}
	return n
}
