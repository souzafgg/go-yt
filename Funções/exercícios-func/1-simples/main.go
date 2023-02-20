package main

import "fmt"

func main() {

	fmt.Println(rInt())
	fmt.Println(rIntString())
}

func rInt() int {
	return 10
}

func rIntString() (int, string) {
	return 10, "string"
}
