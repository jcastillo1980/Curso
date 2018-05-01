package main

import (
	"encoding/json"
	"log"
)

// MiEstrctura ???
type MiEstrctura struct {
	Name string `json:"name"`
	Tlf  string `json:"telefono"`
	Edad int64  `json:"edad"`
}

func main() {
	mivar := MiEstrctura{
		Name: "Javier Castillo Calvo",
		Tlf:  "949271341",
		Edad: 33,
	}
	log.Println("Inicio de aplicación")
	bf, _ := json.Marshal(mivar)
	log.Println(string(bf))
}
