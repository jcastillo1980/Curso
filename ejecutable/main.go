package main

import (
	"fmt"
	"os"
	"time"
)

var retardo = 10

func main() {

	if len(os.Args) > 1 {
		fmt.Printf("Inicio Ejecutable [pid:%d] :\r\n", os.Getpid())

		fmt.Sscan(os.Args[1], &retardo)
		if retardo <= 0 {
			fmt.Println("Contador a cero !!")
			os.Exit(0)
		}

		for i := 0; i < retardo; i++ {
			if len(os.Args) == 3 {
				fmt.Printf("Valor: %d , [%s]\r\n", i+1, os.Args[2])
			} else if len(os.Args) == 4 {
				fmt.Printf("Valor: %d , [%s] [%s]\r\n", i+1, os.Args[2], os.Args[3])
			} else {
				fmt.Printf("Valor: %d\r\n", i+1)
			}
			time.Sleep(time.Second * 1)
		}

		os.Exit(0)

	} else {
		fmt.Println("Error en argumentos")
		os.Exit(-1)
	}
}
