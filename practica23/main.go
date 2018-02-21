package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	defer func() {
		serr := recover()
		if serr != nil {
			fmt.Printf("Fallo {%s} [%T]\r\n", serr, serr)
		}
	}()

	fmt.Println("Probando !!")

	lanzaMierda()
}

func lanzaMierda() {
	fmt.Println("ini lanzaMierda()")
	fmt.Println(os.Stdin.Name())
	if os.Stdin.Name() != "Hola" {
		panic(errors.New("Mierda en lanzaMierda()"))
	}
	fmt.Println("fin lanzaMierda()")
}
