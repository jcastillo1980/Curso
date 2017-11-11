package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Iniciando Practica 3")
	for _, v := range os.Environ() {
		//fmt.Println("[", k, "]->[", v, "]")
		variable := strings.Split(v, "=")
		if len(variable) == 2 {
			//fmt.Println("[", strings.TrimSpace(variable[0]), "]->[", strings.TrimSpace(variable[1]), "]")
			fmt.Printf("[%s]--->[%s]\r\n", variable[0], variable[1])
		}
	}
}
