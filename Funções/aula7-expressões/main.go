package main

import "fmt"

func main() {
	x := 50

	y := func(x int) {
		fmt.Println(x * 500)
	}
	y(x)
}
