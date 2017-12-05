package main

import (
	"fmt"
	"log"
	"net/http"
)

func funcPractica(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Fprintf(w, "FORM:%#v  URL:%#v\r\n", r.Form, r.URL)
}

func main() {
	fmt.Println("http://127.0.0.1:5050")
	http.HandleFunc("/practica", funcPractica)
	log.Fatalln(http.ListenAndServe(":5050", nil))
}
