package main

import "fmt"

func main() {
	x := 5
	func() {
		for x < 10 {
			x++
			fmt.Println(x)
		}
	}()
}
