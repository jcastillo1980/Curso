package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	fmt.Println("Iniciando Practica 3")
	for _, v := range os.Environ() {
		//fmt.Println("[", k, "]->[", v, "]")
		variable := strings.Split(v, "=")
		if len(variable) == 2 {
			//fmt.Println("[", strings.TrimSpace(variable[0]), "]->[", strings.TrimSpace(variable[1]), "]")
			fmt.Printf("[%s]--->[%s]\r\n", variable[0], variable[1])
		}
	}

	cmd := exec.Command("ls", "-alh")
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	err := cmd.Start()
	if err != nil {
		log.Fatalln(err)
	}
	err = cmd.Wait()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("----------OK----------")
}
