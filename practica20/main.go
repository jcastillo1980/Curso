package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("Numero de procesadores: ", runtime.NumCPU())
	fmt.Println("Esto espera 10s ...")
	c := time.After(time.Second * 10)
	s := make(chan int)
	go func(canal chan int) {
		fmt.Println("Numero GoRuninass .. ", runtime.NumGoroutine())
		valor := 0
		fmt.Scanln(&valor)
		canal <- valor
	}(s)
	select {
	case <-c:
		fmt.Println("Fin!! (TimeOut)")
	case v := <-s:
		fmt.Println("Fin!!", v)
	}
}
