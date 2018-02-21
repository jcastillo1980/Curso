package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/jcastillo1980/Curso/email/sendmail"
)

func main() {
	log.Println("En marcha")

	ss := sendmail.SendMailObj{
		Host:     "smtp.serviciodecorreo.es",
		Port:     "587",
		From:     "webkonery@xuitec.com",
		NameFrom: "Web de Gestion",
		Pass:     "cst2018C",
		To:       []string{"jcastillo@xuitec.com"},
	}
	//ss.GetParamSendMail()
	//ss.PrintObj()
	//log.Println(ss.SendTextMail("tema", "Probando ... "))

	myHandler := http.NewServeMux()
	myHandler.Handle("/EM.php", ss)
	myHandler.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Add("Content-Type", "text/html")
		fmt.Fprintf(w, "%v", time.Now())
	})
	s := &http.Server{
		Addr:           ":8080",
		Handler:        myHandler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}
