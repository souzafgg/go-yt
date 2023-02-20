package main

import (
	"fmt"
	"math"
)

type quadrado struct {
	base, altura float64
}
type circulo struct {
	raio float64
}

func (q quadrado) area() float64 {
	area := q.base * q.altura
	return area
}
func (c circulo) area() float64 {
	area := math.Pi * (c.raio * c.raio)
	return area
}

type calculo interface {
	area() float64
}

func retorno(c calculo) {
	fmt.Println(c.area())
}
func main() {

	q := quadrado{base: 5, altura: 2}
	b := circulo{raio: 4}

	retorno(q)
	retorno(b)

	// fmt.Println("Área quadrado:", q.area())
	// fmt.Println("Área do círculo", b.area())

}

// métodos > interfaces > função q toma interface como parâmetro
