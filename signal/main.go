package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	fmt.Println("INICIO", os.Getpid())
	signalChan := make(chan os.Signal, 1)
	cleanupDone := make(chan bool)
	signal.Notify(signalChan, os.Interrupt, os.Kill, syscall.SIGTERM)
	go func() {
		for s := range signalChan {
			fmt.Println("Received an interrupt, stopping services...", s)
			cleanupDone <- true
		}
	}()
	<-cleanupDone
	fmt.Println("FIN")
}
