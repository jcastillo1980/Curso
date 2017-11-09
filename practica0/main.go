package main

import (
	"fmt"
)

// 0 1 2 3 4 5
// 2 3 4 5
// 1 2 3 4 5
func remove(buff []int, indice int) []int {
	copy(buff[indice:], buff[indice+1:])
	return buff[:len(buff)-1]
}

func main() {
	fmt.Println("Esto es una prueba")
	pp := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(pp)
	fmt.Println(remove(pp, 1))
}
