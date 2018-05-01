package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// PortServidor ?????????????????
var PortServidor = 8080

func functionAjax(w http.ResponseWriter, r *http.Request) {
	mapa := map[string]string{
		"key1": "valor key 1",
		"key2": "valor key 2",
		"key3": "valor key 3",
		"time": fmt.Sprintf("%v", time.Now().Unix()),
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(mapa)
}

func main() {
	log.Printf("Inicio del servidor http://127.0.0.1:%d/\r\n", PortServidor)
	router := http.NewServeMux()
	router.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./public/"))))
	router.HandleFunc("/ajax", functionAjax)
	server := http.Server{
		Addr:           fmt.Sprintf("0.0.0.0:%d", PortServidor),
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatalln(server.ListenAndServe())
}
