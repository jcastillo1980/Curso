package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

// Usuarios ??
type Usuarios struct {
	Nombre string `json:"nombre"`
	Pwd    string `json:"contraseña,omitempty"`
	Admin  bool   `json:"isAdmin"`
}

// ListadoUsuarios ??
var ListadoUsuarios []Usuarios

// IndexGet ??
func IndexGet(w http.ResponseWriter, r *http.Request) {
	tt := time.Now()
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "%s", tt.Format("02/01/2006 15:04:05"))
}

// ListaGet ??
func ListaGet(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ListadoUsuarios)
}

// VariablesEntorno ??
func VariablesEntorno(w http.ResponseWriter, r *http.Request) {
	variables := os.Environ()
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(variables)
}

func main() {
	ListadoUsuarios = append(ListadoUsuarios, Usuarios{
		Nombre: "Javier Castillo Calvo",
		Pwd:    "cst2014C",
		Admin:  true,
	})
	ListadoUsuarios = append(ListadoUsuarios, Usuarios{
		Nombre: "Pepe el Tuerto",
		Pwd:    "ss22333",
		Admin:  false,
	})
	ListadoUsuarios = append(ListadoUsuarios, Usuarios{
		Nombre: "Antonio Gomez",
		Pwd:    "",
		Admin:  false,
	})

	fmt.Println("Servidor Web http://127.0.0.1:4040/")
	route := mux.NewRouter()

	route.HandleFunc("/", IndexGet).Methods("GET")
	route.HandleFunc("/lista", ListaGet).Methods("GET")
	route.HandleFunc("/env", VariablesEntorno).Methods("GET")

	log.Fatalln(http.ListenAndServe(":4040", route))
}
