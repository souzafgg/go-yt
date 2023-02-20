package main

import "fmt"

func main() {
	c := retornaVar("Chamando função na função")
	c()
}

func retornaVar(x string) func() {
	return func() {
		fmt.Println(x)
	}
}
