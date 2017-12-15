package tools

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// AppendFile ??
type AppendFile struct {
	NameFile string
}

//Add ??
func (a AppendFile) Add(text string) error {

	if len(a.NameFile) == 0 {
		return nil
	}

	tt := time.Now()
	text = tt.Format("02/01/2006 15:04:05") + " : [" + text + "]\r\n"

	f, err := os.OpenFile(a.NameFile, os.O_APPEND|os.O_CREATE, 0777)
	if err != nil {
		return err
	}

	defer f.Close()

	if _, err = f.WriteString(text); err != nil {
		return err
	}

	return nil
}

// String ??
func (a AppendFile) String() string {
	return fmt.Sprintf("AppendFile{ NameFile: \"%s\"}", a.NameFile)
}

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
