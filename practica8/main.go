package main

import (
	"fmt"
	"os"
)

func main() {
	defer (func() {
		captura := recover()
		fmt.Println("capture en defer main() ->", captura)
	})()

	fmt.Println("Probando Panic!!")
	if len(os.Args) > 1 {
		panic(fmt.Errorf("%s", os.Args[0]))
	}
	fmt.Println("Por aqui no entramos ...")
}
