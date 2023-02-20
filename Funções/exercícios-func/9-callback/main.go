package main

import "fmt"

func main() {
	receber(callBack)
}

func callBack() {
	fmt.Println("callback")
}

func receber(x func()) {
	fmt.Println("Recebendo")
	x()
}
