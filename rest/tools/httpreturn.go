package tools

import (
	"fmt"
	"net/http"
)

// Return500 envia un respuesta 500
func Return500(w http.ResponseWriter, s string) {
	w.Header().Set("Content-Type", "pplication/json")
	w.WriteHeader(500)
	fmt.Fprintf(w, "{\"error\":\"fallo en el servidor\",\"message_error\":\"%s\"}", s)
}

// Return400 envia un respuesta 400
func Return400(w http.ResponseWriter, s string) {
	w.Header().Set("Content-Type", "pplication/json")
	w.WriteHeader(400)
	fmt.Fprintf(w, "{\"error\":\"fallo en la peticion del cliente\",\"message_error\":\"%s\"}", s)
}
