package token

import (
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

const (
	// NAME_HEADER_TOKEN nombre de la cabecera que contiene el token
	NAME_HEADER_TOKEN = "XS-TK"
	// SECOND_LIVE_TOKEN tiempo de vida del token
	SECOND_LIVE_TOKEN = 60 * 60
	secret            = "passwordCaracol1234"
)

// ContenidoToken ????
type ContenidoToken struct {
	ID   string `json:"id"`
	Type string `json:"type"`
	MTS  int64  `json:"mts"`
	jwt.StandardClaims
}

// GenToken ?????
func GenToken(id string, tipo string) (string, error) {
	s := ContenidoToken{
		ID:   id,
		Type: tipo,
		MTS:  time.Now().Unix() + SECOND_LIVE_TOKEN,
	}
	tt := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), &s)
	tss, err := tt.SignedString([]byte(secret))
	if err != nil {
		log.Println("Error en GenToken() --> SignedString ")
		return "", err
	}

	return tss, nil
}

// GetToken ???
func GetToken(t string) *ContenidoToken {
	retorno := &ContenidoToken{}
	tk, err := jwt.ParseWithClaims(t, retorno, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return nil
	}

	if tk.Valid == true {
		return retorno
	}

	return nil
}

// CheckTockenHTTP ????
func CheckTockenHTTP(w http.ResponseWriter, r *http.Request) (id string, tipo string, ret bool) {
	stk := r.Header.Get(NAME_HEADER_TOKEN)
	tk := GetToken(stk)
	if tk == nil {
		return "", "", false
	}
	if tk.MTS >= time.Now().Unix() {
		return "", "", false
	}

	return tk.ID, tk.Type, true
}
