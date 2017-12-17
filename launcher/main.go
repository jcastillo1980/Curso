package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/jcastillo1980/Curso/launcher/ElementExe"
	"github.com/jcastillo1980/Curso/launcher/tools"
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

// Args ??
var Args = ""

func main() {

	log.Printf("Inicio Programa Launcher [pid:%d]\r\n", os.Getpid())

	flag.StringVar(&NameExe, "exe", "", "Nombre del ejecutable.")
	flag.StringVar(&NamePath, "path", "", "Ruta de ejecución.")
	flag.StringVar(&Arg1, "arg1", "", "Argumento 1.")
	flag.StringVar(&Arg2, "arg2", "", "Argumento 2.")
	flag.StringVar(&Arg3, "arg3", "", "Argumento 3.")
	flag.StringVar(&Args, "args", "", "Argumentos para ejecución planificada")
	flag.IntVar(&TR, "tr", 0, "Tiempo (segundos) restart aplicación cuando se cierra sola.")
	flag.IntVar(&TI, "ti", 0, "Tiempo (segundos) retardo inicio.")
	flag.IntVar(&TM, "tm", 0, "Tiempo (segundos) máximo de ejecución para cerrado automatico. ")

	vvv := 0
	flag.IntVar(&vvv, "vvv", 0, "que cosasssss")

	flag.Parse()

	if vvv == 1 {
		//tools.ProbarConexion()
		//fmt.Println("holaaa", time.Now().Unix())
		//tools.SetDbExeRun(13, 5555)
		//tools.SetDbExeStop(13)
		//tools.SetDbTaskRun(13)
		tools.SetDbTaskStop(13)
		fmt.Println(tools.GetListaGroups(true))
		return
	}

	EE := ElementExe.ElementExe{
		Modo:     ElementExe.ModoHora,
		IDG:      1,
		NH:       1,
		HM:       22*60 + 23,
		NameExe:  NameExe,
		NamePath: NamePath,
		TR:       TR,
		TI:       TI,
		TM:       TM,
		Arg1:     Arg1,
		Arg2:     Arg2,
		Arg3:     Arg3,
		Args:     Args,
		NHminut:  true,
	}

	EE.Ejecuta()

	retorno := ""
	fmt.Scanln(&retorno)
	fmt.Println("OK STOP", retorno)
	EE.Stop()
	fmt.Scanln(&retorno)
	fmt.Println("FIN TOTAL", retorno)
}
