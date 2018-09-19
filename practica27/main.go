package main

import (
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	jwt.StandardClaims
}

type XXX struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Tiempo int64  `json:"ts"`
	jwt.StandardClaims
}

func main() {

	tt := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), &XXX{
		Name:   "javier castillo calvo",
		Age:    38,
		Tiempo: time.Now().Unix(),
	})
	tss, err := tt.SignedString([]byte("CACAMALA1234"))
	if err != nil {
		log.Panicln("ERROR")
	}
	//tss = tss + "EEA3F66"
	log.Println(".....", tss)

	re := XXX{}
	tk, err := jwt.ParseWithClaims(tss, &re, func(token *jwt.Token) (interface{}, error) {
		return []byte("CACAMALA1234"), nil
	})
	if tk.Valid == true {
		log.Println(re)
	} else {
		log.Println("ERROR")
	}

}
