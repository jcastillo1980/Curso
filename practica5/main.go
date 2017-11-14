package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hola mundo")
	fmt.Println(`
Esto es asi
Esto no es asi...
		`)

	arr := [10]int{0, 1, 2, 3, 5, 6, 7, 8, 9}
	sl := arr[:1]
	fmt.Printf("%T\r\n", arr)
	fmt.Printf("%T  ... %#v\r\n", sl, sl)
	sl = append(sl, 11, 22, 33, 44)
	fmt.Printf("%#v  cap=%d\r\n", sl, cap(sl))

	elementos := make(map[string]string)
	elementos["cap"] = "olimex"
	elementos["zop"] = "javier"

	for k, v := range elementos {
		fmt.Printf("elementos[%s]=%s\r\n", k, v)
	}

	if valor, ok := elementos["xx"]; ok == true {
		fmt.Println("Existe: zop ->", valor)
	} else {
		fmt.Println("No Existe")
	}
}
