package main

import "fmt"

func main() {

	canal := make(chan int)

	go func() {
		canal <- 10
	}()

	fmt.Println(<-canal)
}
