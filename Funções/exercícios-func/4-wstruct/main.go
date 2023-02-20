package main

import "fmt"

type pessoa struct {
	Nome      string
	Sobrenome string
	Idade     int
}

func (p pessoa) retorno() {
	fmt.Println(p.Nome, p.Sobrenome, p.Idade)
}

func main() {

	n := pessoa{
		Nome:      "Adalberto",
		Sobrenome: "Alves",
		Idade:     20,
	}

	n.retorno()
}
