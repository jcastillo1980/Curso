package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

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
func ReadDirectorio(filtro string) {
	files, err := filepath.Glob(filtro)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(files)
}

func main() {

	ReadDirectorio("*")
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
