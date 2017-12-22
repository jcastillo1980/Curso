package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"
)

// FragmentaAgumentos Crea un slice de string de una cadena con argumentos
func FragmentaAgumentos(s string) []string {
	sec := strings.Split(s, " ")
	retorno := []string{}
	for _, valor := range sec {
		if len(valor) != 0 {
			retorno = append(retorno, valor)
		}
	}

	return retorno
}

func main() {
	port := 0
	flag.IntVar(&port, "port", 8585, "Puerto del servidor.")
	flag.Parse()
	/*if len(os.Args) > 1 {
		cmd := exec.Command(os.Args[1], os.Args[2:]...)
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.SysProcAttr = &syscall.SysProcAttr{
			Setpgid: true,
			Setsid:  true,
		}
		err := cmd.Start()
		if err != nil {
			fmt.Printf("0")
		} else {
			fmt.Printf("%d", cmd.Process.Pid)
		}

	} else {
		fmt.Printf("0")
	}*/
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%d", time.Now().Unix())
	})
	http.HandleFunc("/exe", func(w http.ResponseWriter, r *http.Request) {
		bf, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			fmt.Fprintf(w, "0")
		} else {
			args := FragmentaAgumentos(string(bf))
			if len(args) > 0 {
				cmd := exec.Command(args[0], args[1:]...)
				cmd.Stderr = os.Stderr
				cmd.Stdin = os.Stdin
				cmd.Stdout = os.Stdout
				err := cmd.Start()
				if err != nil {
					fmt.Fprintf(w, "0")
				} else {
					go func() {
						cmd.Wait()
					}()
					fmt.Fprintf(w, "%d", cmd.Process.Pid)
				}

			}
		}
	})
	log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))

}
