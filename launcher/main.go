package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/jcastillo1980/Curso/launcher/ElementExe"
)

// NameExe ??
var NameExe = ""

// NamePath ??
var NamePath = ""

// TR ??
var TR = 0

// TI ??
var TI = 0

// TM ??
var TM = 0

// Arg1 ??
var Arg1 = ""

// Arg2 ??
var Arg2 = ""

// Arg3 ??
var Arg3 = ""

func main() {
	var argumentos []string
	var temporizador *time.Timer
	var timeout bool

	log.Printf("Inicio Programa Launcher [pid:%d]\r\n", os.Getpid())

	flag.StringVar(&NameExe, "exe", "", "Nombre del ejecutable.")
	flag.StringVar(&NamePath, "path", "", "Ruta de ejecución.")
	flag.StringVar(&Arg1, "arg1", "", "Argumento 1.")
	flag.StringVar(&Arg2, "arg2", "", "Argumento 2.")
	flag.StringVar(&Arg3, "arg3", "", "Argumento 3.")
	flag.IntVar(&TR, "tr", 0, "Tiempo (segundos) restart aplicación cuando se cierra sola.")
	flag.IntVar(&TI, "ti", 0, "Tiempo (segundos) retardo inicio.")
	flag.IntVar(&TM, "tm", 0, "Tiempo (segundos) máximo de ejecución para cerrado automatico. ")
	flag.Parse()

	if len(NameExe) == 0 {
		log.Println("Sin Programa a ejecutar !!")
		os.Exit(0)
	}

	if len(Arg1) != 0 {
		argumentos = append(argumentos, Arg1)
	}
	if len(Arg2) != 0 {
		argumentos = append(argumentos, Arg2)
	}
	if len(Arg3) != 0 {
		argumentos = append(argumentos, Arg3)
	}

	cmd := exec.Command(NameExe, argumentos...)
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	if len(NamePath) != 0 {
		cmd.Path = NamePath
	}

	timeout = false
	if TM > 0 {
		temporizador = time.AfterFunc(time.Second*time.Duration(TM), func() {
			cmd.Process.Kill()
			timeout = true
		})
	}
	err := cmd.Start()

	canal := make(chan int)
	go func(c chan int) {
		err = cmd.Wait()
		c <- 1
	}(canal)

	<-canal

	temporizador.Stop()

	if TM > 0 && timeout == true {
		log.Printf("TimeOut!! [pid:%d]\r\n", cmd.Process.Pid)
	}

	log.Printf("Final Proceso [pid:%d]: %v\r\n", cmd.Process.Pid, err)
	fmt.Println(ElementExe.ElementExe{})
}
