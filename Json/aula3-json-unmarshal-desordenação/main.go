package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type informacoes struct {
	Nome      string `json:"Nome"`
	Sobrenome string `json:"Sobrenome"`
	Profissao string `json:"Profissao"`
}

func main() {
	s := []byte(`{"Nome":"james","Sobrenome":"smith","Profissao":"ferreiro"}`)

	var j informacoes
	erro := json.Unmarshal(s, &j)
	if erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(j)
}
