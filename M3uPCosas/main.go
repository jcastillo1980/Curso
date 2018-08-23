package main

import (
	"log"
)

const (
	mINOMBRE = "Javier Castillo Calvo"
)

func printB(bf []byte) {
	if bf == nil {
		log.Println("es nil !!")
	} else if len(bf) != 22 {
		log.Println("Error")
	} else {
		log.Println(string(bf))
	}

}

func main() {
	log.Println("Hola:", mINOMBRE)
	buf := []byte{55, 56, 57, 58, 59}
	ff := []byte("ñoñ")

	log.Printf("%s \r\n %#v", string(buf), ff)

	printB(nil)
}
