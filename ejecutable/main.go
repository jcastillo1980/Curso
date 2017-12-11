package main

import (
	"fmt"
	"os"
	"time"
)

var retardo = 10

func main() {
	if len(os.Args) > 1 {
		fmt.Printf("Inicio Ejecutable [%d] :\r\n", os.Getpid())

		fmt.Sscan(os.Args[1], &retardo)
		if retardo <= 0 {
			fmt.Println("Contador a cero !!")
			os.Exit(0)
		}

		for i := 0; i < retardo; i++ {
			fmt.Println("Valor:", i+1)
			time.Sleep(time.Second * 1)
		}

		os.Exit(0)

	} else {
		fmt.Println("Error en argumentos")
		os.Exit(-1)
	}
}
