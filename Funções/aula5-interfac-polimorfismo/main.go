package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

type dentista struct {
	pessoa
	dentesarrancados int
	salario          float64
}

type arquiteto struct {
	pessoa
	tipodeconstrução string
	tamanhodalocura  string
}

func (x dentista) oibomdia() {
	fmt.Println("Meu nome é:", x.nome, "e já arranquei", x.dentesarrancados, "dentes")
}

func (x arquiteto) oibomdia() {
	fmt.Println("Meu nome é:", x.nome, "e faço o tipo de construção de", x.tipodeconstrução)
}

type gente interface {
	oibomdia()
}

func serhumano(g gente) {
	g.oibomdia()
}

func main() {
	mrdente := dentista{
		pessoa: pessoa{
			nome:      "Mr dente",
			sobrenome: "Da silva",
			idade:     50,
		},
		dentesarrancados: 500,
		salario:          5000,
	}
	mrpredio := arquiteto{
		pessoa: pessoa{
			nome:      "Mr prédio",
			sobrenome: "De souza",
			idade:     50,
		},
		tipodeconstrução: "casas",
		tamanhodalocura:  "grande",
	}
	serhumano(mrdente)
	serhumano(mrpredio)
}
