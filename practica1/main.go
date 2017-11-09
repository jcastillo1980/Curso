package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	for _, v := range os.Args[1:] {
		resp, err := http.Get(v)
		if err != nil {
			log.Fatalln("Error Apertura: ", v, " ---> ", err)
		}
		cont, _ := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		fmt.Println(string(cont))
	}
}
