package main

import (
	"fmt"
)

func main() {
	fmt.Println("Esto Es una prueba")
	var i int
	for i = 1; i < 10; i++ {
		fmt.Printf("%d....\r\n", i)
	}
}
