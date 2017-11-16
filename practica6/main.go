package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"sync"
)

var mutex = sync.Mutex{}

func main() {
	fmt.Println("Comprimiento ...")
	for _, v := range os.Args[1:] {
		go fmt.Println(comprime(v))
	}
}

func comprime(name string) error {
	mutex.Lock()
	fmt.Println("comprimiendo: " + name + " -> " + name + ".gz")
	dsrc, err := os.Open(name)
	if err != nil {
		return err
	}
	defer dsrc.Close()

	ddst, err := os.Create(name + ".gz")
	if err != nil {
		return err
	}
	defer ddst.Close()

	dgz := gzip.NewWriter(ddst)
	_, err = io.Copy(dgz, dsrc)

	dgz.Close()

	mutex.Unlock()

	return err
}
