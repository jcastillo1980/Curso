package modulo

import (
	"fmt"
)

func init() {
	fmt.Println("Inicio package {modulo}")
}

// GetValor ???
func GetValor(valorEntrada int) (err error, resultado int) {
	if valorEntrada > 100 {
		resultado = 0
		err = fmt.Errorf("Valor Entrada > 100")
	} else {
		resultado = valorEntrada + 1
		err = nil
	}

	return
}
