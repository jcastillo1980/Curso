package main

import (
	"flag"
	"fmt"
	"html"
	"log"
	"net/http"
)

var (
	// PuertoServidor ??????
	PuertoServidor int
	// TextoRespuesta ??????
	TextoRespuesta string
	// Debug ????
	Debug bool
)

func init() {
	flag.IntVar(&PuertoServidor, "port", 8080, "Puerto del servidor Web")
	flag.StringVar(&TextoRespuesta, "texto", "1", "Texto respuesta en cualquier petición")
	flag.BoolVar(&Debug, "debug", true, "Mensaje de depuración")
}

// Controlador ?????
func Controlador(w http.ResponseWriter, r *http.Request) {
	if Debug == true {
		log.Println(html.EscapeString(r.URL.Path))
	}

	if html.EscapeString(r.URL.Path) == "/favicon.ico" {
		if Debug == true {
			log.Println("favicon close")
		}
		return
	}

	fmt.Fprintf(w, "%s", TextoRespuesta)

}

func main() {
	flag.Parse()

	http.HandleFunc("/", Controlador)

	if Debug == true {
		log.Println("Inicio del Servidor Puerto:", PuertoServidor)
	}

	log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%d", PuertoServidor), nil))
}
