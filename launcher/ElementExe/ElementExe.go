package ElementExe

import (
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/jcastillo1980/Curso/launcher/tools"
)

const (
	// ModoFix ???
	ModoFix = 0
	//ModoHora ???
	ModoHora = 1
)

//ElementExe ??????
type ElementExe struct {
	Modo     int
	IDG      int
	HM       int
	NH       int
	NameExe  string
	NamePath string
	TR       int
	TI       int
	TM       int
	Arg1     string
	Arg2     string
	Arg3     string
	// Args son los argumentos para tareas planificdas
	Args string
	// NHminut hace que cada NH no sean horas si no minutos
	NHminut  bool
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
	if e.Modo == ModoFix {
		go e.ejecutaBloqueoFix()
	} else if e.Modo == ModoHora {
		if e.NH == 0 {
			go e.ejecutaBloqueoHM()
		} else {
			go e.ejecutaBloqueoNH()
		}
	}

}

// ejecutaBloqueoHM ejecuta una tarea cuando llega una hora en concreto. .Stop() la para. Esta funcion se queda bloqueada
func (e *ElementExe) ejecutaBloqueoNH() {

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

	argumentos := tools.FragmentaAgumentos(e.Args)

	tsAntes := time.Now().Unix()
	tsDespues := tsAntes
	diferencia := 0

	for {
		if e.NHminut == true {
			diferencia = e.NH*60 - int(tsDespues-tsAntes)
		} else {
			diferencia = e.NH*60*60 - int(tsDespues-tsAntes)
		}

		if diferencia > 0 {
			log.Printf("[%d] Esperando  .. %d sec [Cada %d horas]\r\n", e.IDG, diferencia, e.NH)
			canalesperainicial := make(chan int)
			temporizador = time.AfterFunc(time.Second*time.Duration(diferencia), func() {
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
		} else {
			log.Printf("[%d] Esperando  .. %d sec [Cada %d horas]\r\n", e.IDG, 0, e.NH)
		}

		tsAntes = time.Now().Unix()

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
		}

		log.Printf("[%d] Final Proceso [pid:%d]: %v\r\n", e.IDG, cmd.Process.Pid, err)

		tsDespues = time.Now().Unix()

		// aqui ya se ha terminado el ejecutable, y volvemos

	}
}

// ejecutaBloqueoHM ejecuta una tarea cuando cada numero de horas .Stop() la para. Esta funcion se queda bloqueada
func (e *ElementExe) ejecutaBloqueoHM() {

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

	argumentos := tools.FragmentaAgumentos(e.Args)

	for {

		hmahora := time.Now().Minute() + time.Now().Hour()*60
		hmdif := e.HM - hmahora
		if hmdif < 0 {
			hmdif = 1440 + hmdif
		}

		if hmdif > 0 {
			log.Printf("[%d] Esperando  .. %d [%d:%d]\r\n", e.IDG, hmdif, e.HM/60, e.HM%60)
			canalesperainicial := make(chan int)
			temporizador = time.AfterFunc(time.Minute*time.Duration(hmdif), func() {
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
		}

		log.Printf("[%d] Final Proceso [pid:%d]: %v\r\n", e.IDG, cmd.Process.Pid, err)

		log.Printf("[%d] Esperando 1 Minuto\r\n", e.IDG)
		canalesperareinicio := make(chan int)
		temporizador = time.AfterFunc(time.Minute*1, func() {
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

//ejecutaBloqueo ejecuta una tarea de forma indefinida hasta que se lanza ".Stop()". Esta funcion se queda bloqueada
func (e *ElementExe) ejecutaBloqueoFix() {

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
