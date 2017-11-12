package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {

	ValorEntero := flag.Int("valor", 0, "Valor Entero de linea de comandos de ejemplo")
	ValorCadena := flag.String("cadena", "?", "Valor Cadena de linea de comando de ejemplo")
	ValorEspera := flag.Duration("delay", time.Second*1, "Tiempo de espera")
	flag.Parse()

	fmt.Println("Valor Entero: ", *ValorEntero)
	fmt.Println("Valor Cadena: ", *ValorCadena)
	fmt.Println("Tiempo espera: ", *ValorEspera)
	fmt.Printf("%#v\r\nwait....", flag.Args())
	time.Sleep(*ValorEspera)
	fmt.Println("")
}
