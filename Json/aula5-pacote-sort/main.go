package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []string{"nome", "sobrenome", "smith", "profissão", "ferreiro"}
	sort.Strings(s)
	fmt.Println(s)

	i := []int{6, 4, 2, 76, 8, 2, 1, 3, 6, 7, 8, 5, 2}
	sort.Ints(i)
	fmt.Println(i)
}
