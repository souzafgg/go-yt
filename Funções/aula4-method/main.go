package main

import "fmt"

type pessoa struct {
	Nome  string
	Idade int
}

func (p pessoa) bomdia() {
	fmt.Println(p.Nome, "diz bom dia")
}

func main() {
	m := pessoa{"mauricio", 30}
	m.bomdia()

}
