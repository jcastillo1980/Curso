package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Pars ?????
type Pars struct {
	Parametros []Ps `json:"parametros"`
}

// Ps ??????
type Ps struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("Ejemplo Json!!")
	bf, err := ioutil.ReadFile("./datos.json")
	must(err)

	var xxxx Pars
	must(json.Unmarshal(bf, &xxxx))
	fmt.Printf("%#v\r\n", xxxx)
	for i, v := range xxxx.Parametros {
		fmt.Printf("%d -> [%s]:[%s]\r\n", i, v.Name, v.Value)
	}
}
