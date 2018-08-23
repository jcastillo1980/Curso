package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jcastillo1980/Curso/practica26/nocache"
	"github.com/jcastillo1980/Curso/practica26/tools"
)

var (
	directorio string
	puerto     int
)

func init() {
	flag.StringVar(&directorio, "dir", "./public", "Directorio de los fichero estático del servidor web")
	flag.IntVar(&puerto, "port", 8080, "Puerto de escucha del servidor web")

	flag.Parse()
}

func funcGetAllElement(w http.ResponseWriter, r *http.Request) {
	tools.ReturnError(w, 200, "que cosas  !!!")
	//fmt.Fprintf(w, "caca")
}

func funcDefault(w http.ResponseWriter, r *http.Request) {
	tools.Return500(w, fmt.Sprintf("No implementado URL:%v  Method:%v", r.URL, r.Method))
}

func main() {
	log.Println("Esto es un servidor de prueba en GO   s!!")

	router := mux.NewRouter().StrictSlash(true)

	// servidor de ficheros estaticos
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, directorio+"/index.html")
	}).Methods("GET")
	router.PathPrefix("/static/").Handler(nocache.NoCache(http.StripPrefix("/", http.FileServer(http.Dir(directorio)))))

	// api
	router.HandleFunc("/api/v2/elementos", funcGetAllElement).Methods("GET")

	// si hay algo que no existe !!
	router.NotFoundHandler = http.HandlerFunc(funcDefault)

	server := &http.Server{
		Handler:        router,
		Addr:           fmt.Sprintf(":%d", puerto),
		WriteTimeout:   15 * time.Second,
		ReadTimeout:    15 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Println(fmt.Sprintf("http://127.0.0.1:%d/", puerto))
	log.Fatal(server.ListenAndServe())

}
