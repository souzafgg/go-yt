package main

import (
	"fmt"
)

var pl = fmt.Println
var dia string

func main() {
	// valor := soma(10, 10)
	// pl(valor)

	// basica()

	total, length, _ := argumento("tarde", 5, 10, 5, 20, 25)
	pl(total, length)
}

func argumento(s string, y ...int) (int, int, string) {
	if s == "manhã" {
		pl("bom dia")
	} else if s == "tarde" {
		pl("boa tarde")
	} else {
		pl("boa noite")
	}

	soma := 0
	for _, v := range y {
		soma += v
	}
	return soma, len(y), s
}
