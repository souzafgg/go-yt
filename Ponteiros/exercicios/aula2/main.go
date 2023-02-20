package main

import "fmt"

type pessoa struct {
	Nome string
}

func mudeMe(p *pessoa) {
	(*p).Nome = "Luiz"
}

func main() {
	x := pessoa{"hana"}
	fmt.Println(x)
	mudeMe(&x)
	fmt.Println(x)

}
