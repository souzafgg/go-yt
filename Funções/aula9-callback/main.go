package main

import "fmt"

func main() {
	t := somenteImpar(soma, []int{50, 51, 52, 53, 54, 55}...)
	fmt.Println(t)
}

func soma(x ...int) int {
	n := 0
	for _, v := range x {
		n = +v
	}
	return n
}

func somenteImpar(f func(x ...int), y ...int) int {
	var slice []int
	for _, v := range y {
		if v%2 != 0 {
			slice = append(slice, v)
		}

	}
	total := f(slice...)
	return total
}
