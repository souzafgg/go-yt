package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type pessoa struct {
	Nome      string
	Sobrenome string
	Profissao string
}

func main() {

	james := pessoa{"james", "smith", "ferreiro"}

	jaeger := pessoa{
		Nome:      "jager",
		Sobrenome: "jr",
		Profissao: "guerreiro",
	}

	j, err := json.Marshal(james)
	if err != nil {
		log.Fatal(err)
	}

	k, err := json.Marshal(jaeger)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(j))
	fmt.Println(string(k))
}
