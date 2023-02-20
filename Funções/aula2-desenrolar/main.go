package main

import (
	"fmt"
)

// passado uma slice como argumento na função soma, naturalmente a função não a aceita pois deve-se desenrolar a slice antes passando
// a variável, no caso s, com ... (s...)
func main() {
	s := []int{5, 10, 15, 20, 25}

	total := soma(s...)
	fmt.Println(total)
}

func soma(y ...int) int {
	soma := 0
	for _, v := range y {
		soma += v
	}
	return soma
}
