package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("******** Lectura de teclado durante 60 segundos *********")
	canal := make(chan []byte)
	timeout := time.After(60 * time.Second)
	go ReadStdin(canal)
	for {
		select {
		case buf := <-canal:
			os.Stdout.Write(buf)
		case <-timeout:
			os.Exit(0)
		}
	}
}

// ReadStdin funcion que lee del teclado y lo envia por el canal en formato slice []byte
func ReadStdin(out chan []byte) {
	for {
		data := make([]byte, 1024)
		l, _ := os.Stdin.Read(data)
		if l > 0 {
			out <- data
		}
	}
}
