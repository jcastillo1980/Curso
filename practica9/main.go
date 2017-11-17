package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	// CategoryBooks ??
	CategoryBooks = iota
	// CategoryHealth ??
	CategoryHealth
	// CategoryClothing ??
	CategoryClothing
)

const (
	// ValorPrincipal ??
	ValorPrincipal = iota
	//ValorSecundario ??
	ValorSecundario
	//ValorTercero ??
	ValorTercero
)

func main() {
	var cadena string
	fmt.Printf("valor: %#v,%#v,%#v,%#v\r\n", CategoryBooks, CategoryHealth, CategoryClothing, 3.33)
	scn := bufio.NewScanner(os.Stdin)
	for scn.Scan() {
		cadena = scn.Text()
		if cadena == "fin" {
			break
		}
		fmt.Println("--->", cadena, "<----")
	}
	fmt.Printf("%#v\r\n", os.FileMode(0777))
}
