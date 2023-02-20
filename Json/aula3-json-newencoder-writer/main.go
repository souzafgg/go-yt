package main

import (
	"encoding/json"
	"fmt"
	"os"
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
	fmt.Println(james, jaeger)
	encoder := json.NewEncoder(os.Stdout)
	encoder.Encode(james)
}
