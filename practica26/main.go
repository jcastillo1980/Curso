package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/jcastillo1980/Curso/practica26/nocache"
	"github.com/jcastillo1980/Curso/practica26/token"
	"github.com/jcastillo1980/Curso/practica26/tools"
)

var (
	directorio string
	puerto     int
)

// Respuesta ??????
type Respuesta struct {
	Name      string `json:"name"`
	Edad      int    `json:"edad"`
	Direccion string `json:"address"`
}

// Respuestas ?????
var Respuestas []Respuesta

func funElementos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	json.NewEncoder(w).Encode(Respuestas)
}

func funHTML(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	valores := vars["valores"]

	r.ParseForm()

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(200)

	fmt.Fprintf(w, "<%s> el valor es %s [%#v]", r.FormValue("valor1"), valores, r.Form)

	//json.NewEncoder(w).Encode(Respuestas)
}

func funPost(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "application/json")
	//w.WriteHeader(200)

	r.ParseForm()

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(200)

	fmt.Fprintf(w, "los valores pasado por post [%#v]", r.Form)

	//json.NewEncoder(w).Encode(Respuestas)
}

type PostJ struct {
	Pv1 string `json:"pv1"`
	Pv2 string `json:"pv2"`
	Pv3 int    `json:"pv3"`
	Pv4 string `json:"pv4"`
}

func funPostJ(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	var rr PostJ
	json.NewDecoder(r.Body).Decode(&rr)

	/*ss := map[string]string{
		"casa":   "nada",
		"nombre": "javier castillo calvo",
	}*/
	r.ParseForm()
	rr.Pv4 = fmt.Sprintf("%#v", r.Form)
	json.NewEncoder(w).Encode(rr)

	//json.NewEncoder(w).Encode(Respuestas)
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func funWS(w http.ResponseWriter, r *http.Request) {

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	go func(conn *websocket.Conn) {
		defer conn.Close()

		for {
			messageType, p, err := conn.ReadMessage()
			if err != nil {
				log.Println(err)
				return
			}
			if err := conn.WriteMessage(messageType, p); err != nil {
				log.Println(err)
				return
			}
		}
	}(conn)
}

func init() {
	flag.StringVar(&directorio, "dir", "./public", "Directorio de los fichero estático del servidor web")
	flag.IntVar(&puerto, "port", 8080, "Puerto de escucha del servidor web")

	flag.Parse()

	Respuestas = append(Respuestas, Respuesta{
		Name:      "javier Castillo Calvo",
		Edad:      38,
		Direccion: "Alguna parte",
	})

	Respuestas = append(Respuestas, Respuesta{
		Name:      "Pepe el tuerto",
		Edad:      44,
		Direccion: "Cualquiere dirección",
	})
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
	log.Println(token.GenToken("javier", "1"))

	router := mux.NewRouter().StrictSlash(true)

	// servidor de ficheros estaticos
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, directorio+"/index.html")
	}).Methods("GET")
	router.PathPrefix("/static/").Handler(nocache.NoCache(http.StripPrefix("/", http.FileServer(http.Dir(directorio)))))

	// api
	router.HandleFunc("/api/v2/elementos", funElementos).Methods("GET")
	router.HandleFunc("/api/v2/html/{valores}", funHTML).Methods("GET")
	router.HandleFunc("/api/v2/post", funPost).Methods("POST")
	router.HandleFunc("/api/v2/postJ", funPostJ).Methods("POST")
	router.HandleFunc("/api/v2/ws", funWS)

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
