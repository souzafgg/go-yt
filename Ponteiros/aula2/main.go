package main

import "fmt"

func main() {
	y := 10

	recebeValor(y)
	recebePointer(&y)
}

func recebeValor(x int) {
	x++
	fmt.Printf("Valor da var: %v\n", x)
}

func recebePointer(x *int) {
	*x++
	fmt.Printf("Valor do pointer: %v", x)
}
