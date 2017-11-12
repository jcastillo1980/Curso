package main

import (
	"flag"
	"fmt"
)

func main() {

	ValorEntero := flag.Int("valor", 0, "Valor Entero de linea de comandos de ejemplo")
	ValorCadena := flag.String("cadena", "?", "Valor Cadena de linea de comando de ejemplo")

	flag.Parse()

	fmt.Println("Valor Entero: ", *ValorEntero)
	fmt.Println("Valor Cadena: ", *ValorCadena)
	fmt.Printf("%#v\r\n", flag.Args())
}
