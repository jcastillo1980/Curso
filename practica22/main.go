package main

import (
	"io/ioutil"
	"log"
	"os"
)

func main() {
	log.Println("Probando escritura fichero !!")
	ioutil.WriteFile("./texto.txt", []byte("Esto es una prueba!!\r\n"), 0777)
	fs, err := os.Open("./texto.txt")
	if err != nil {
		log.Panicln(err)
	}
	defer fs.Close()

	valor, _ := ioutil.ReadAll(fs)
	log.Printf("%#v\r\n", valor)

	ds, err := ioutil.ReadDir("./")
	if err != nil {
		log.Panicln(err)
	}
	for _, dd := range ds {
		log.Printf("---> %#v\r\n\r\n\r\n", dd)
	}
}
