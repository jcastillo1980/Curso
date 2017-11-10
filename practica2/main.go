package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Mia Esto es una estructura de prueba.
type Mia struct {
	Nombre   string `json:"nombre"`
	Edad     int    `json:"edad,omitempty"`
	IsHombre bool   `json:"ishombre"`
}

func must(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	fmt.Println("Probando Estructuras")
	mm := Mia{
		Nombre:   "Javier Castillo Calvo",
		IsHombre: true,
	}

	buf, err := json.Marshal(mm)
	must(err)

	fmt.Println("<", string(buf), ">")
	fmt.Printf("%#v\r\n", mm)
}
