package main

import (
	"bufio"
	"crypto/tls"
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type Cosas struct {
	valor int `json:"value"`
}

func (c Cosas) String() string {
	return fmt.Sprintf("<%d>", c.valor)
}

func init() {
	log.Println("Hola Mundo")
}

func handleConnection(c net.Conn) {
	fmt.Printf("Serving %s\n", c.RemoteAddr().String())
	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		temp := strings.TrimSpace(string(netData))
		if temp == "STOP" {
			break
		}

		result := strconv.Itoa(rand.Int()) + "\n"
		c.Write([]byte(string(result)))
	}
	c.Close()
}

func mainServer() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a port number!")
		return
	}

	PORT := ":" + arguments[1]
	l, err := net.Listen("tcp4", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()
	rand.Seed(time.Now().Unix())

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleConnection(c)
	}
}

// ReadDirectorio ?????
func ReadDirectorio(filtro string) []string {
	files, err := filepath.Glob(filtro)
	if err != nil {
		log.Fatal(err)
	}
	return files
}

// LeerCSV ????
func LeerCSV(file string) map[string]string {

	resp, err := os.Open(file)
	if err != nil {
		return nil
	}
	defer resp.Close()

	r := csv.NewReader(resp)
	r.Comma = ';'
	r.FieldsPerRecord = 0

	mapa := make(map[string]string)
	nombres := []string{}

	record, err := r.Read()
	if err == io.EOF {
		return nil
	}
	if err != nil {
		log.Println(err)
		return nil
	}

	for _, v := range record {
		nombres = append(nombres, v)
	}

	record, err = r.Read()
	if err == io.EOF {
		return nil
	}
	if err != nil {
		log.Println(err)
		return nil
	}

	for i, v := range record {
		mapa[nombres[i]] = v
	}

	return mapa
}

// GetWeb ????
func GetWeb() {
	link := "https://eslaremotecontroller.tk"

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	response, err := client.Get(link)
	if err != nil {
		log.Println(err)
	}
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	s := strings.TrimSpace(string(content))

	log.Println(s)
}

// PostWeb ???
func PostWeb() {
	link := "https://eslaremotecontroller.tk/pruebaPost.php"

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	response, err := client.PostForm(link, url.Values{"key": {"Value"}, "id": {"123"}})
	if err != nil {
		log.Println(err)
	}
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	s := strings.TrimSpace(string(content))

	log.Println(s)
}

func main() {

	// obj := Cosas{2222}
	// fmt.Println(obj)

	// db.UpdateName(37, "zzzzzzzzzo")
	// db.ListaNombres()
	// db.ListaNombresMS()

	PostWeb()

	os.Exit(-1)

	l := ReadDirectorio("/Users/javiercastillocalvo/Downloads/*.csv")
	for _, v := range l {
		fmt.Println("-------------------------------------------------------")
		fmt.Println("FILA:", v)
		fmt.Println("-------------------------------------------------------")
		fmt.Printf("%#v\r\n", LeerCSV(v))
		fmt.Println("*******************************************************")
	}

	os.Exit(-1)

	mm := map[string]string{}
	ss := []int{11, 22, 33, 44, 55, 66, 77, 88, 99}

	mm["uno"] = "texto uno"
	mm["dos"] = "texto dos"
	mm["tres"] = "texto tres"

	for k, v := range mm {
		log.Println(k, "->", v)
	}

	for index, value := range ss {
		log.Printf("ss[%d] = %d\r\n", index, value)
	}

	canal := make(chan bool)
	canal2 := make(chan bool)

	go func(c chan bool) {
		for i := 0; i < 10000; i++ {
			log.Println("1", i)
			time.Sleep(time.Millisecond * 500)
		}

		c <- true
	}(canal)

	go func(c chan bool) {
		for i := 0; i < 1500000000; i++ {
			log.Println("2", i)
			time.Sleep(time.Millisecond * 500)
		}

		c <- true
	}(canal2)

	go mainServer()

	for {
		select {
		case <-canal:
			log.Println("ha terminado uno")
		case <-canal2:
			log.Println("ha terminado dos")
			os.Exit(3)
		}
	}

}
