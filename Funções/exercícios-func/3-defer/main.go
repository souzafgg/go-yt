package main

import (
	"fmt"
	"strings"
)

var n = 0
var n1 = 0
var x int

func main() {

	defer fmt.Println("Deseja iniciar novamente?")
	r := ""
	fmt.Scan(&r)
	if strings.HasPrefix(r, "s") {
		for n <= 2 {
			fmt.Println("Digite dois números para soma: ")
			fmt.Scan(&n, &n1)
			x = somaSimples(n, n1)
		}
		fmt.Println(x)

	} else {
		fmt.Println("Encerrando")
	}
}

func somaSimples(x int, y int) int {
	n := x * y
	return n
}
