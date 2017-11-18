package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/jcastillo1980/PruebaGo/nocache"
)

// Page ?????
type Page struct {
	Title   string
	Content string
}

func paginaPrincipal(w http.ResponseWriter, r *http.Request) {
	p := &Page{
		Title:   "Página de Prueba",
		Content: "Probando los template",
	}

	t := template.Must(template.ParseFiles("./template/template.html"))
	t.Execute(w, p)
}

func main() {
	fmt.Println("Server http://127.0.0.1:5050")

	http.Handle("/", nocache.NoCache(http.StripPrefix("/", http.FileServer(http.Dir("./public")))))
	http.HandleFunc("/template", paginaPrincipal)
	log.Fatalln(http.ListenAndServe(":5050", nil))

}
