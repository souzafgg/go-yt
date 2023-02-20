package main

import "fmt"

func main() {

	x := 10
	y := &x

	fmt.Println(*y) //dereferenciação -- mostrar o q tá no endereço

}
