package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Probando NewRequest")
	req, err := http.NewRequest("GET", "http://www.xuitec.com/", nil)
	if err != nil {
		log.Fatalln("Error: ", err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalln("Error: ", err)
	}
	fmt.Printf("%#v\r\n", res.Header)

	b, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Println(string(b))
}
