package main

import (
	"encoding/json"
	"fmt"
)

const (
	//MyNombre ??
	MyNombre = "Javier Castillo Calvo"
	//MyValor ??
	MyValor = 44
)

// Estructura ??
type Estructura struct {
	Nombre    string `json:"nombre"`
	Edad      int16  `json:"edad"`
	Direccion string `json:"direccion"`
}

// ToJSON ??
func (e Estructura) ToJSON() string {
	bb, err := json.Marshal(e)
	if err != nil {
		return ""
	}

	return string(bb)
}

func tipoVariable(v interface{}) {
	//fmt.Printf("%T\r\n", v)
	switch v.(type) {
	case string:
		fmt.Println("string")
	case int:
		fmt.Println("int")
	case uint16:
		fmt.Println("uint16")
	}
}

func main() {
	valor := Estructura{
		Nombre:    "Javier Castillo Calvo",
		Edad:      int16(44),
		Direccion: "España",
	}
	fmt.Println("Probando : ", MyNombre)
	tipoVariable(uint16(MyValor))
	fmt.Println(valor.ToJSON())
}
