package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var lista = make(map[string]net.Conn)
var bloqueo = &sync.Mutex{}

func addLista(ipp string, nc net.Conn) {
	bloqueo.Lock()
	lista[ipp] = nc
	bloqueo.Unlock()
}

func delLista(ipp string) {
	bloqueo.Lock()
	delete(lista, ipp)
	bloqueo.Unlock()
}

func getLista(ipp string) (net.Conn, bool) {
	bloqueo.Lock()
	a, b := lista[ipp]
	bloqueo.Unlock()

	return a, b
}

func mainServer() {
	arguments := os.Args
	if len(arguments) == 1 {
		log.Println("Please provide a port number!")
		return
	}

	PORT := ":" + arguments[1]
	l, err := net.Listen("tcp4", PORT)
	if err != nil {
		log.Println(err)
		return
	}
	defer l.Close()
	rand.Seed(time.Now().Unix())

	for {
		c, err := l.Accept()
		if err != nil {
			log.Println(err)
			return
		}
		go handleConnection(c)
	}
}

func listar() {
	for {
		bloqueo.Lock()
		for k := range lista {
			fmt.Println(k)
		}
		bloqueo.Unlock()
		time.Sleep(10 * time.Second)
	}
}

func handleConnection(c net.Conn) {

	ipp := c.RemoteAddr().String()
	log.Printf("Serving %s\n", ipp)

	addLista(ipp, c)

	defer func(sipp string) {
		delLista(sipp)
	}(ipp)

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

func main() {
	log.Println("Inicio gClient")

	// db.Setup("oficina.xuitec.com", 3306, "externo", "externo", "blx2")
	// db.ListaNombres()

	// go listar()

	// mainServer()

	bf := make([]byte, 30)
	for i := 0; i < len(bf); i++ {
		bf[i] = 0
	}
	bf[0] = byte('H')
	bf[1] = byte('O')
	bf[2] = byte('L')
	bf[3] = byte('A')
	bf[12] = byte('A')

	log.Printf("[%s] %d\r\n", string(bf), len(string(bf)))

	resultado := ""
	for i := 0; i < len(bf); i++ {
		if bf[i] == 0 {
			resultado = string(bf[0:i])
			break
		}
	}

	log.Printf("[%s] %d\r\n", resultado, len(resultado))

	//rd := strings.NewReader("Esto un pruba y nada mas")

	/*for {
		bb, rr := os.Stdin.Read
		if rr != nil {
			if rr == io.EOF {
				log.Fatal("Final")
			} else {
				log.Fatal(rr)
			}
		}

		log.Println(bb)
	}*/

}
