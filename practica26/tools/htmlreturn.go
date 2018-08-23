package tools

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type respuestaError struct {
	NumeroError  int    `json:"nError`
	MensajeError string `json:"messageError"`
}

// Return500 envia un respuesta 500
func Return500(w http.ResponseWriter, s string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)
	fmt.Fprintf(w, "{\"error\":\"fallo en el servidor\",\"message_error\":\"%s\"}", s)
}

// Return400 envia un respuesta 400
func Return400(w http.ResponseWriter, s string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)
	fmt.Fprintf(w, "{\"error\":\"fallo en la peticion del cliente\",\"message_error\":\"%s\"}", s)
}

// ReturnError ??????
func ReturnError(w http.ResponseWriter, code int, mensaje string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	msg := respuestaError{
		NumeroError:  code,
		MensajeError: mensaje,
	}
	json.NewEncoder(w).Encode(msg)
}
