package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jcastillo1980/curso/rest/tools"
)

// ServerPort ???
var ServerPort int

// PathPublicDir ???
var PathPublicDir string

func init() {
	flag.IntVar(&ServerPort, "port", 8080, "Puerto del Servidor")
	flag.StringVar(&PathPublicDir, "path", "./public/", "Directorio de los ficheros estáticos")
}

// FuncGetAll envia todos los elementos
func FuncGetAll(w http.ResponseWriter, r *http.Request) {
	tools.Return500(w, "FuncGetAll() no implementado")
}

// FuncGetByID envia todos los elementos que coincide con el ID
func FuncGetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	tools.Return500(w, fmt.Sprintf("FuncGetByID(%s) no implementado", id))
}

// FuncCreate crea un nuevo elemnto
func FuncCreate(w http.ResponseWriter, r *http.Request) {
	tools.Return500(w, "FuncCreate() no implementado")
}

// FuncUpdate actualiza elementos
func FuncUpdate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	tools.Return500(w, fmt.Sprintf("FuncUpdate(%s) no implementado", id))
}

// FuncDelete borra elementos
func FuncDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	tools.Return500(w, fmt.Sprintf("FuncDelete(%s) no implementado", id))
}

// FuncDefault funcion por defecto
func FuncDefault(w http.ResponseWriter, r *http.Request) {
	tools.Return500(w, fmt.Sprintf("No implementado URL:%v  Method:%v", r.URL, r.Method))
}

func main() {
	flag.Parse()

	router := mux.NewRouter().StrictSlash(true)
	router.Handle("/", tools.NoCache(http.StripPrefix("/", http.FileServer(http.Dir(PathPublicDir)))))

	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", ServerPort),
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	// definir el api
	router.HandleFunc("/api/element", FuncGetAll).Methods("GET")
	router.HandleFunc("/api/element/{id}", FuncGetByID).Methods("GET")
	router.HandleFunc("/api/element", FuncCreate).Methods("POST")
	router.HandleFunc("/api/element/{id}", FuncUpdate).Methods("PUT")
	router.HandleFunc("/api/element/{id}", FuncDelete).Methods("DELETE")

	// funcion por defecto
	router.NotFoundHandler = http.HandlerFunc(FuncDefault)

	log.Printf("Servidor en marcha http://127.0.0.1:%d/ \r\n", ServerPort)
	log.Fatalln(server.ListenAndServe())
}
