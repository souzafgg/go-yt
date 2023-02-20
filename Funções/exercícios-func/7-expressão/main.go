package main

import "fmt"

var x float64

func main() {
	fmt.Println("Digite um valor a ser multiplicado por 3.14: ")
	fmt.Scan(&x)

	f := func(x float64) {
		c := x * 3.14
		fmt.Println(c)
	}
	f(x)
}
