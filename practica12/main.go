package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	fmt.Println("http://127.0.01:5050/")

	http.HandleFunc("/datos", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Time: %#v", time.Now())
	})
	log.Fatalln(http.ListenAndServe(":5050", nil))
}
