package ElementExe

import (
	"log"
	"os"
	"os/exec"
	"time"
)

//ElementExe ??????
type ElementExe struct {
	IDG      int
	NameExe  string
	NamePath string
	TR       int
	TI       int
	TM       int
	Arg1     string
	Arg2     string
	Arg3     string
	Finaliza chan int
	run      bool
}

// Stop ???
func (e *ElementExe) Stop() {
	if e.run == true {
		e.Finaliza <- 1
	}
	e.run = false
}

//Ejecuta ejecuta la tarea en un nuevo hilo
func (e *ElementExe) Ejecuta() {
	go e.ejecutaBloqueo()
}

//ejecutaBloqueo ejecuta una tarea de forma indefinida hasta que se lanza "Finaliza<-1". Esta funcion se queda blouqeda
func (e *ElementExe) ejecutaBloqueo() {

	e.Finaliza = make(chan int, 2)
	e.run = true
	defer func() {
		e.run = false
		close(e.Finaliza)
		log.Printf("[%d] Tarea Finalizada\r\n", e.IDG)
	}()

	var temporizador *time.Timer
	var timeout bool

	if len(e.NameExe) == 0 {
		log.Printf("[%d] Sin Programa a ejecutar !!\r\n", e.IDG)
		return
	}

	argumentos := []string{}
	if len(e.Arg1) != 0 {
		argumentos = append(argumentos, e.Arg1)
	}
	if len(e.Arg2) != 0 {
		argumentos = append(argumentos, e.Arg2)
	}
	if len(e.Arg3) != 0 {
		argumentos = append(argumentos, e.Arg3)
	}

	if e.TI > 0 {
		log.Printf("[%d] Esperando Inicio .. (%d s)\r\n", e.IDG, e.TI)
		canalesperainicial := make(chan int)
		temporizador = time.AfterFunc(time.Second*time.Duration(e.TI), func() {
			canalesperainicial <- 1
		})
		select {
		case <-canalesperainicial:
			temporizador.Stop()
			close(canalesperainicial)
		case <-e.Finaliza:
			temporizador.Stop()
			close(canalesperainicial)
			return
		}

	}
	for {

		cmd := exec.Command(e.NameExe, argumentos...)
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		if len(e.NamePath) != 0 {
			cmd.Path = e.NamePath
		}

		timeout = false
		if e.TM > 0 {
			temporizador = time.AfterFunc(time.Second*time.Duration(e.TM), func() {
				cmd.Process.Kill()
				timeout = true
			})
		}

		err := cmd.Start()

		log.Printf("[%d] Inicio de Ejecutable [pid:%d]\r\n", e.IDG, cmd.Process.Pid)

		canal := make(chan int)
		go func(c chan int) {
			err = cmd.Wait()
			c <- 1
		}(canal)

		hefinalizado := false
		pasa := false
		for pasa == false {
			select {
			case <-canal:
				pasa = true
			case <-e.Finaliza:
				cmd.Process.Kill()
				hefinalizado = true
			}
		}

		if e.TM > 0 {
			temporizador.Stop()
		}

		if e.TM > 0 && timeout == true {
			log.Printf("[%d] TimeOut!! [pid:%d]\r\n", e.IDG, cmd.Process.Pid)
		}

		if hefinalizado == true {
			log.Printf("[%d] Final Proceso y Final TAREA [pid:%d]: %v\r\n", e.IDG, cmd.Process.Pid, err)
			return
		} else {
			log.Printf("[%d] Final Proceso [pid:%d]: %v\r\n", e.IDG, cmd.Process.Pid, err)
		}

		if e.TR > 0 {
			log.Printf("[%d] Esperando ReInicio .. (%d s)\r\n", e.IDG, e.TR)
			canalesperareinicio := make(chan int)
			temporizador = time.AfterFunc(time.Second*time.Duration(e.TR), func() {
				canalesperareinicio <- 1
			})
			select {
			case <-canalesperareinicio:
				close(canalesperareinicio)
				temporizador.Stop()
			case <-e.Finaliza:
				close(canalesperareinicio)
				temporizador.Stop()
				return
			}

		}

	}
}
