package main

import (
	"fmt"
	"html/template"
	"os"

	"github.com/jcastillo1980/Curso/practica17/modulo"
)

const (
	//NameFile1 Nombre del fichero para template
	NameFile1 = "./file1.html"
)

type tema struct {
	Nombre  string
	Funcion string
	Anno    int
}

func main() {
	estructura := tema{
		Nombre:  "Javier Castillo Calvo",
		Funcion: "Administrador",
		Anno:    0,
	}
	fmt.Printf("Probando Template de html ->>>[")
	t := template.Must(template.ParseFiles(NameFile1))
	t.ExecuteTemplate(os.Stdout, "TT", &estructura)
	fmt.Printf("]\r\n")

	_, v := modulo.GetValor(44)
	fmt.Println("-------------", v)
}
