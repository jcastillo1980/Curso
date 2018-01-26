package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

// Parametros ???
type Parametros struct {
	XMLName    xml.Name `xml:"parametros"`
	Parametros []Param  `xml:"var"`
	Otros      []Param  `xml:"otros>var"`
}

// Param ???
type Param struct {
	Name      string `xml:"name,attr"`
	Contenido string `xml:",chardata"`
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("Parseando un fichero xml")
	bf, err := ioutil.ReadFile("./datos.xml")
	must(err)

	var xmlEstructura Parametros
	err = xml.Unmarshal(bf, &xmlEstructura)
	must(err)
	for i, v := range xmlEstructura.Parametros {
		fmt.Printf("%d -> Name:[%s] Contenido:[%s]\r\n", i, v.Name, v.Contenido)
	}
	for i, v := range xmlEstructura.Otros {
		fmt.Printf("---> %d -> Name:[%s] Contenido:[%s]\r\n", i, v.Name, v.Contenido)
	}
	//fmt.Println(string(bf))

}
