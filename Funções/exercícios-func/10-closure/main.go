package main

import "fmt"

func main() {

	primeira()()

}

func primeira() func() {
	p := "palavra"
	fmt.Println(p)
	return func() {
		c := "segunda"
		fmt.Println(c)
	}
}
