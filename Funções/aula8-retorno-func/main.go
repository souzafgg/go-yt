package main

import "fmt"

func main() {
	// x := retornaFuncao()
	// y := x(3)
	y := retornaFuncao()(3)
	fmt.Println(y)
}

func retornaFuncao() func(int) int {
	return func(i int) int {
		return i * 10
	}
}
